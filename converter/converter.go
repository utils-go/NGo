package converter

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertToIntFromString(s string) (int, error) {
	return strconv.Atoi(s)
}
func ConvertToStringFromInt(i int) string {
	return strconv.Itoa(i)
}
func ConvertStringToBool(s string) (bool, error) {
	if strings.ToLower(s) == "true" {
		return true, nil
	}
	if strings.ToLower(s) == "false" {
		return false, nil
	}
	if strings.ToLower(s) == "0" {
		return false, nil
	}
	if strings.ToLower(s) == "1" {
		return true, nil
	}
	return false, fmt.Errorf("%s is a invalid bool string", s)
}

func ConvertBoolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}
