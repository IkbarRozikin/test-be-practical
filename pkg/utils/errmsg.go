package utils

import (
	"fmt"
)

func ErrMsg(context string, err error) error {
	return fmt.Errorf("%s: %w", context, err)
}
