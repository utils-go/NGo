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
	return int(t.ticks / TicksPerDay)
}
func (t *TimeSpan) Hours() int {
	return int(t.ticks / TicksPerHour)
}
func (t *TimeSpan) Milliseconds() int {
	return int(t.ticks / TicksPerMillisecond)
}
func (t *TimeSpan) Minutes() int {
	return int(t.ticks / TicksPerMinute)
}
func (t *TimeSpan) Seconds() int {
	return int(t.ticks / TicksPerSecond)
}
func (t *TimeSpan) Ticks() int {
	return int(t.ticks)
}
func (t *TimeSpan) TotalDays() float64 {
	return float64(t.ticks) * DaysPerTick
}
func (t *TimeSpan) TotalHours() float64 {
	return float64(t.ticks) * HoursPerTick
}
func (t *TimeSpan) TotalMilliseconds() float64 {
	return float64(t.ticks) * MillisecondsPerTick
}
func (t *TimeSpan) TotalMinutes() float64 {
	return float64(t.ticks) * MillisecondsPerTick
}
func (t *TimeSpan) TotalSeconds() float64 {
	return float64(t.ticks) * SecondsPerTick
}
func (t *TimeSpan) Add(ts TimeSpan) *TimeSpan {
	totalTicks := t.ticks + ts.ticks
	return &TimeSpan{ticks: totalTicks}
}
func (t *TimeSpan) Duration() *TimeSpan {
	if t.ticks > 0 {
		return &TimeSpan{ticks: t.ticks}
	} else {
		return &TimeSpan{ticks: -t.ticks}
	}

}
func (t *TimeSpan) Equals(ts TimeSpan) bool {
	return t.ticks == ts.ticks
}
func (t *TimeSpan) Subtract(ts TimeSpan) *TimeSpan {
	totalTicks := t.ticks - ts.ticks
	return &TimeSpan{ticks: totalTicks}
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
	if t1.ticks > t2.ticks {
		return 1
	}
	if t1.ticks < t2.ticks {
		return -1
	}
	return 0
}
func FromDays(value float64) *TimeSpan {
	return interval(value, MillisPerDay)
}
func FromHours(value float64) *TimeSpan {
	return interval(value, MillisPerHour)
}
func FromMilliseconds(value float64) *TimeSpan {
	return interval(value, 1)
}
func FromMinutes(value float64) *TimeSpan {
	return interval(value, MillisPerMinute)
}
func FromSeconds(value float64) *TimeSpan {
	return interval(value, MillisPerSecond)
}
func FromTicks(value int64) *TimeSpan {
	return &TimeSpan{ticks: value}
}
func Parse(s string) (*TimeSpan, error) {
	//TODO
	return nil, nil
}
func ParseExact(s string, cslayout string) (*TimeSpan, error) {
	//TODO
	return nil, nil
}
func interval(value float64, scale int) *TimeSpan {
	tmp := value * float64(scale)
	var millis float64
	if value > 0 {
		millis = tmp + 0.5
	} else {
		millis = tmp - 0.5
	}

	return &TimeSpan{ticks: int64(millis) * TicksPerMillisecond}
}
