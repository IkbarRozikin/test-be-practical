package models

type ConsumptionItem struct {
	Name     string `json:"name"`
	MaxPrice int    `json:"maxPrice"`
}

type Consumption struct {
	Name string `json:"name"`
	Jml  int    `json:"jml"`
}

type Booking struct {
	BookingDate     string `json:"bookingDate"`
	OfficeName      string `json:"officeName"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	ListConsumption []Consumption
	Participants    int    `json:"participants"`
	RoomName        string `json:"roomName"`
	ID              string `json:"id"`
}

type RoomData struct {
	ID                    string        `json:"id"`
	RoomName              string        `json:"room_name"`
	BookingDate           string        `json:"booking_date"`
	PersentasePemakaian   float64       `json:"persentase_pemakaian"`
	TotalPriceConsumption float64       `json:"totalPrice_consumption"`
	ListConsumption       []Consumption `json:"list_consumption"`
}

type GetBookingsResponse struct {
	OfficeName  string     `json:"office_name"`
	DataBooking []RoomData `json:"data_booking"`
}
