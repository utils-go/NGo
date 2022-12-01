package utils

import "fmt"

var maxSeconds int64 = 922337203685
var minSeconds int64 = -922337203685
var ticksPerSecond int64 = 10000000

func IsLeapYear(year int) bool {
	if year < 1 || year > 9999 {
		return false
	}
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func TimeToTicks(hour, minute, second int) (int64, error) {
	totalSeconds := int64(hour)*3600 + int64(minute)*60 + int64(second)
	if totalSeconds > maxSeconds || totalSeconds < minSeconds {
		return 0, fmt.Errorf("totalseconds:%d is out of range", totalSeconds)
	}
	return totalSeconds * ticksPerSecond, nil
}
