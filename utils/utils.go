package utils

import (
	"fmt"
)

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15:04",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyy":                "2006",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"mm:ss":               "04:05",
	}
}

func ConvertLayout(cslayout string) (string, error) {
	////yyyy-MM-dd HH:mm:ss
	////2006-01-02 15:04:05
	////year
	//if strings.Contains(cslayout, "yyyy") {
	//	cslayout = strings.ReplaceAll(cslayout, "yyyy", "2006")
	//}
	//if strings.Contains(cslayout, "yy") {
	//	cslayout = strings.ReplaceAll(cslayout, "yy", "06")
	//}
	//if strings.Contains(cslayout, "y") {
	//	cslayout = strings.ReplaceAll(cslayout, "y", "6")
	//}
	////Moth
	//if strings.Contains(cslayout, "MM") {
	//	cslayout = strings.ReplaceAll(cslayout, "MM", "01")
	//}
	//if strings.Contains(cslayout, "M") {
	//	cslayout = strings.ReplaceAll(cslayout, "M", "1")
	//}
	////DAY
	//if strings.Contains(cslayout, "dd") {
	//	cslayout = strings.ReplaceAll(cslayout, "dd", "02")
	//}
	//if strings.Contains(cslayout, "d") {
	//	cslayout = strings.ReplaceAll(cslayout, "d", "2")
	//}
	////Hour
	//if strings.Contains(cslayout, "HH") {
	//	cslayout = strings.ReplaceAll(cslayout, "HH", "15")
	//}
	//if strings.Contains(cslayout, "H") {
	//	cslayout = strings.ReplaceAll(cslayout, "H", "5")
	//}
	////Minute
	//if strings.Contains(cslayout, "mm") {
	//	cslayout = strings.ReplaceAll(cslayout, "mm", "04")
	//}
	//if strings.Contains(cslayout, "m") {
	//	cslayout = strings.ReplaceAll(cslayout, "m", "4")
	//}
	////Second
	//if strings.Contains(cslayout, "ss") {
	//	cslayout = strings.ReplaceAll(cslayout, "ss", "05")
	//}
	//if strings.Contains(cslayout, "s") {
	//	cslayout = strings.ReplaceAll(cslayout, "d", "5")
	//}
	////millisecond
	//if strings.Contains(cslayout, "f") {
	//	cslayout = strings.ReplaceAll(cslayout, "f", "0")
	//}
	//return cslayout
	str, ok := timeFormat[cslayout]
	if ok {
		return str, nil
	} else {
		return "", fmt.Errorf("%s is not the right format", cslayout)
	}
}
