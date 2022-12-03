package timespan

import (
	"fmt"
)

const (
	TicksPerMillisecond int64   = 10000
	MillisecondsPerTick float64 = 0.0001
	TicksPerSecond      int64   = 10000000
	SecondsPerTick              = 1e-07
	TicksPerMinute              = 600000000
	MinutesPerTick              = 1.6666666666666667e-09
	TicksPerHour                = 36000000000
	HoursPerTick                = 2.7777777777777777e-11
	TicksPerDay         int64   = 864000000000
	DaysPerTick                 = 1.1574074074074074e-12
	MillisPerSecond             = 1000
	MillisPerMinute             = 60000
	MillisPerHour               = 3600000
	MillisPerDay                = 86400000
	MaxSeconds                  = 922337203685
	MinSeconds                  = -922337203685
	MaxMilliSeconds             = 922337203685477
	MinMilliSeconds             = -922337203685477
	TicksPerTenthSecond         = 1000000
)

var DaysToMonth365 = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}
var DaysToMonth366 = []int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335, 366}

type TimeSpan struct {
	ticks int64
}

func NewTimeSpanFromTicks(ticks int64) *TimeSpan {
	return &TimeSpan{ticks: ticks}
}

func (t *TimeSpan) Days() int {

}
func (t *TimeSpan) Hours() int {

}
func (t *TimeSpan) Milliseconds() {

}
func (t *TimeSpan) Minutes() int {

}
func (t *TimeSpan) Seconds() int {

}
func (t *TimeSpan) Ticks() {

}
func (t *TimeSpan) TotalDays() {

}
func (t *TimeSpan) TotalHours() int {

}
func (t *TimeSpan) TotalMiniseconds() int {

}
func (t *TimeSpan) TotalMinutes() int {

}
func (t *TimeSpan) TotalSeconds() int {

}
func (t *TimeSpan) Add(ts TimeSpan) *TimeSpan {

}
func (t *TimeSpan) Duration() *TimeSpan {

}
func (t *TimeSpan) Equals() {

}
func (t *TimeSpan) Subtract() *TimeSpan {

}
func (t *TimeSpan) ToString() string {

}

type TimeSpanBuilder struct {
	day         int64
	hour        int64
	minute      int64
	second      int64
	millisecond int64
}

func (t *TimeSpanBuilder) SetDay(day int) *TimeSpanBuilder {
	t.day = int64(day)
	return t
}
func (t *TimeSpanBuilder) SetHour(hour int) *TimeSpanBuilder {
	t.hour = int64(hour)
	return t
}
func (t *TimeSpanBuilder) SetMinute(minute int) *TimeSpanBuilder {
	t.minute = int64(minute)
	return t
}
func (t *TimeSpanBuilder) SetSecond(second int) *TimeSpanBuilder {
	t.second = int64(second)
	return t
}
func (t *TimeSpanBuilder) SetMilliSecond(millisecond int) *TimeSpanBuilder {
	t.millisecond = int64(millisecond)
	return t
}
func (t *TimeSpanBuilder) Build() (*TimeSpan, error) {
	totalMilliSeconds := (t.day*3600*24+t.hour*3600+t.minute*60+t.second)*1000 + t.millisecond
	if totalMilliSeconds > MaxMilliSeconds || totalMilliSeconds < MinMilliSeconds {
		return nil, fmt.Errorf("totoalmiliseconds : %d is out of range", totalMilliSeconds)
	}
	ticks := totalMilliSeconds * TicksPerMillisecond
	return &TimeSpan{ticks: ticks}, nil
}

// static method
func Compare(t1, t2 *TimeSpan) int {

}
func FromDays(value float64) *TimeSpan {

}
func FromHours(value float64) *TimeSpan {

}
func FromMilliseconds(value float64) *TimeSpan {

}
func FromMinutes(value float64) *TimeSpan {

}
func FromSeconds(value float64) *TimeSpan {

}
func FromTicks(value float64) *TimeSpan {

}
func Parse(s string) (*TimeSpan, error) {

}
func ParseExact(s string, cslayout string) (*TimeSpan, error) {

}
