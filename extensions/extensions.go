package extensions

import "time"

type NgoTime struct {
	t *time.Time
}

type layoutInfo struct {
	index    int
	cslayout string
	golayout string
}

var layouts []layoutInfo

func init() {
	layouts = []layoutInfo{
		//year
		{index: 0, cslayout: "y", golayout: "2"},
		{index: 1, cslayout: "y", golayout: "0"},
		{index: 2, cslayout: "y", golayout: "0"},
		{index: 3, cslayout: "y", golayout: "6"},
		//month
		{index: 0, cslayout: "M", golayout: "01"},
		{index: 1, cslayout: "M", golayout: "0"},
		//day
		{index: 0, cslayout: "d", golayout: "2"},
		{index: 1, cslayout: "d", golayout: "0"},
		//hour
		{index: 0, cslayout: "H", golayout: "2"},
		{index: 1, cslayout: "H", golayout: "0"},
		//Minute
		{index: 0, cslayout: "m", golayout: "2"},
		{index: 1, cslayout: "m", golayout: "0"},
		//second
		{index: 0, cslayout: "s", golayout: "2"},
		{index: 1, cslayout: "s", golayout: "0"},
	}
}
func NewNgoTime(t time.Time) *NgoTime {
	return &NgoTime{t: &t}
}

func (t *NgoTime) ToString(cslayout string) {

}

// convert cslayout to golayout
func formatConvert(csLayout string) string {

}
