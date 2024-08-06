package services

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"test_be_practical/internal/models"
	"test_be_practical/pkg/utils"
)

const (
	bookingURL     = "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"
	consumptionURL = "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"
	roomCapacity   = 73 // Asumsi kapasitas tetap untuk semua ruangan Asumsi harga tetap per item
)

func GetBookings() ([]models.GetBookingsResponse, error) {
	bookings, err := fetchBookings()
	if err != nil {
		return nil, utils.ErrMsg("error fetching bookings: %w", err)
	}

	consumptionItems, err := fetchConsumptionData()
	if err != nil {
		return nil, utils.ErrMsg("error fetching consumption data: %w", err)
	}

	if len(bookings) == 0 {
		return nil, utils.ErrMsg("no bookings found", err)
	}

	priceMap := make(map[string]int)
	for _, item := range consumptionItems {
		priceMap[item.Name] = item.MaxPrice
	}

	officeDataMap := make(map[string]models.GetBookingsResponse)
	for _, booking := range bookings {
		result, exists := officeDataMap[booking.OfficeName]
		if !exists {
			result = models.GetBookingsResponse{
				OfficeName:  booking.OfficeName,
				DataBooking: make([]models.RoomData, 0),
			}
		}

		roomData := models.RoomData{
			ID:              booking.ID,
			RoomName:        booking.RoomName,
			BookingDate:     booking.BookingDate,
			ListConsumption: make([]models.Consumption, 0),
		}

		for _, consumption := range booking.ListConsumption {
			_, exists := priceMap[consumption.Name]
			if !exists {
				continue
			}

			found := false
			for i, item := range roomData.ListConsumption {
				if item.Name == consumption.Name {
					roomData.ListConsumption[i].Jml += booking.Participants
					found = true
					break
				}
			}
			if !found {
				roomData.ListConsumption = append(roomData.ListConsumption, models.Consumption{
					Name: consumption.Name,
					Jml:  booking.Participants,
				})
			}
		}

		var totalPrice float64
		for _, item := range roomData.ListConsumption {
			price, exists := priceMap[item.Name]
			if exists {
				totalPrice += float64(item.Jml * price)
			}
		}

		usagePercentage := (float64(booking.Participants) / float64(roomCapacity)) * 100
		roundedPercentage := math.Round(usagePercentage*100) / 100
		roomData.PersentasePemakaian = roundedPercentage
		roomData.TotalPriceConsumption = totalPrice

		result.DataBooking = append(result.DataBooking, roomData)
		officeDataMap[booking.OfficeName] = result
	}

	officeDataSlice := make([]models.GetBookingsResponse, 0, len(officeDataMap))
	for _, v := range officeDataMap {
		officeDataSlice = append(officeDataSlice, v)
	}

	return officeDataSlice, nil
}

func fetchBookings() ([]models.Booking, error) {
	resp, err := http.Get(bookingURL)
	if err != nil {
		return nil, utils.ErrMsg("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, utils.ErrMsg("error reading response body: %w", err)
	}

	var bookings []models.Booking
	if err := json.Unmarshal(body, &bookings); err != nil {
		return nil, utils.ErrMsg("error decoding JSON: %w", err)
	}

	return bookings, nil
}

func fetchConsumptionData() ([]models.ConsumptionItem, error) {
	resp, err := http.Get(consumptionURL)
	if err != nil {
		return nil, utils.ErrMsg("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, utils.ErrMsg("error reading response body: %w", err)
	}

	var consumptionItems []models.ConsumptionItem
	if err := json.Unmarshal(body, &consumptionItems); err != nil {
		return nil, utils.ErrMsg("error decoding consumption JSON", err)
	}

	return consumptionItems, nil
}
