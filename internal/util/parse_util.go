package util

import (
	"strconv"

	"AvitoTest/internal/constants"
)

func ParseUserId(userStr string) (uint, error) {
	result, err := strconv.ParseUint(userStr, constants.Base10, constants.BitSize)
	return uint(result), err
}
