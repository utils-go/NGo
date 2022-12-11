package datetime

import (
	"github.com/lishuangquan1987/ngo/timespan"
	"github.com/lishuangquan1987/ngo/utils"
	"time"
)

const (
	// Number of days in a non-leap year
	daysPerYear = 365
	// Number of days in 4 years
	daysPer4Years = daysPerYear*4 + 1 // 1461
	// Number of days in 100 years
	daysPer100Years = daysPer4Years*25 - 1 // 36524
	// Number of days in 400 years
	daysPer400Years = daysPer100Years*4 + 1 // 146097

	// Number of days from 1/1/0001 to 12/31/1600
	daysTo1601 = daysPer400Years * 4 // 584388
	// Number of days from 1/1/0001 to 12/30/1899
	daysTo1899 = daysPer400Years*4 + daysPer100Years*3 - 367
	// Number of days from 1/1/0001 to 12/31/1969
	daysTo1970 = daysPer400Years*4 + daysPer100Years*3 + daysPer4Years*17 + daysPerYear // 719,162
	// Number of days from 1/1/0001 to 12/31/9999
	daysTo10000 = daysPer400Years*25 - 366 // 3652059

	maxTicks = daysTo10000*timespan.TicksPerDay - 1
	minTicks = 0
)

var MaxValue = DateTime{t: time.Unix(0, maxTicks)}
var MinValue = DateTime{t: time.Unix(0, minTicks)}

type DateTime struct {
	t time.Time
}

func Now() *DateTime {
	return &DateTime{t: time.Now()}
}
func NewDateTime(t time.Time) *DateTime {
	return &DateTime{t: t}
}

func (t *DateTime) ToString(cslayout string) string {
	cslayout = utils.ConvertLayout(cslayout)
	return t.t.Format(cslayout)
}

func (t *DateTime) ToTime() *time.Time {
	return &t.t
}
func (t *DateTime) Add(dateTime DateTime) {

}
func (t *DateTime) String() string {
	return t.t.String()
}
func (t *DateTime) Date() *DateTime {
	//internalTick := t.t.UnixNano() & 0x3FFFFFFFFFFFFFFF
	//internalKind := t.t.UnixNano() & 0xC000000000000000
	//
	//newTick := (internalTick - internalTick%timespan.TicksPerDay) | internalKind
	//return &DateTime{t: time.Unix(0, newTick)}
	year, month, day := t.t.Date()
	return &DateTime{t: time.Date(year, month, day, 0, 0, 0, 0, time.Local)}
}

//func (t *DateTime) getDatePart(part int) int {
//	internalTick := t.t.UnixNano() & 0x3FFFFFFFFFFFFFFF
//	ticks := internalTick
//	// n = number of days since 1/1/0001
//	n := int(ticks / timespan.TicksPerDay)
//	// y400 = number of whole 400-year periods since 1/1/0001
//	y400 := n / daysPer400Years
//	// n = day number within 400-year period
//	n -= y400 * daysPer400Years
//	// y100 = number of whole 100-year periods within 400-year period
//	y100 := n / daysPer100Years
//	// Last 100-year period has an extra day, so decrement result if 4
//	if y100 == 4 {
//		y100 = 3
//	}
//	// n = day number within 100-year period
//	n -= y100 * daysPer100Years
//	// y4 = number of whole 4-year periods within 100-year period
//	y4 := n / daysPer4Years
//	// n = day number within 4-year period
//	n -= y4 * daysPer4Years
//	// y1 = number of whole years within 4-year period
//	y1 := n / daysPerYear
//	// Last year has an extra day, so decrement result if 4
//	if y1 == 4 {
//		y1 = 3
//	}
//	// If year was requested, compute and return it
//	if part == 0 {
//		return y400*400 + y100*100 + y4*4 + y1 + 1
//	}
//	// n = day number within year
//	n -= y1 * daysPerYear
//	// If day-of-year was requested, return it
//	if part == 1 {
//		return n + 1
//	}
//	// Leap year calculation looks different from IsLeapYear since y1, y4,
//	// and y100 are relative to year 1, not year 0
//	leapYear := y1 == 3 && (y4 != 24 || y100 == 3)
//	days := make([]int, 0)
//	if leapYear {
//		days = timespan.DaysToMonth366
//	} else {
//		days = timespan.DaysToMonth365
//	}
//	// All months have less than 32 days, so n >> 5 is a good conservative
//	// estimate for the month
//	m := n>>5 + 1
//	// m = 1-based month number
//	for {
//		if n >= days[m] {
//			m++
//		} else {
//			break
//		}
//	}
//	// If month was requested, return it
//	if part == 2 {
//		return m
//	}
//	// Return 1-based day-of-month
//	return n - days[m-1] + 1
//}

func (t *DateTime) Day() int {
	return t.t.Day()
}
func (t *DateTime) Month() int {
	return int(t.t.Month())
}
func (t *DateTime) Year() int {
	return t.t.Year()
}
func (t *DateTime) DayOfWeek() time.Weekday {
	return t.t.Weekday()
}
func (t *DateTime) DayOfYear() int {
	return t.t.YearDay()
}
func (t *DateTime) Hour() int {
	return t.t.Hour()
}
