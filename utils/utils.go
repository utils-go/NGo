package utils

import (
	"strings"
)

var timeFormat map[string]string

func init() {
	//timeFormat = map[string]string{
	//	"yyyy-MM-dd HH:mm:ss": "2006-01-02 15:04:05",
	//	"yyyy-MM-dd HH:mm":    "2006-01-02 15:04",
	//	"yyyy-MM-dd HH":       "2006-01-02 15:04",
	//	"yyyy-MM-dd":          "2006-01-02",
	//	"yyyy-MM":             "2006-01",
	//	"mm-dd":               "01-02",
	//	"dd-MM-yy HH:mm:ss":   "02-01-06 15:04:05",
	//	"yyyy/MM/dd HH:mm:ss": "2006/01/02 15:04:05",
	//	"yyyy/MM/dd HH:mm":    "2006/01/02 15:04",
	//	"yyyy/MM/dd HH":       "2006/01/02 15",
	//	"yyyy/MM/dd":          "2006/01/02",
	//	"yyyy/MM":             "2006/01",
	//	"mm/dd":               "01/02",
	//	"dd/MM/yy HH:mm:ss":   "02/01/06 15:04:05",
	//	"yyyy":                "2006",
	//	"MM":                  "01",
	//	"HH:mm:ss":            "15:04:05",
	//	"mm:ss":               "04:05",
	//
	//	"yyyyMMddHHmmss": "20060102150405",
	//	"yyyyMMdd":       "20060102",
	//	"YYYYMM":         "200601",
	//}
	timeFormat = map[string]string{
		"yyyy": "2006",
		"MM":   "01",
		"dd":   "02",
		"HH":   "15",
		"mm":   "04",
		"ss":   "05",
		"fff":  "000",
	}
}

func ConvertLayout(cslayout string) string {
	cslayoutNew := strings.ReplaceAll(cslayout, "h", "H")
	for k, v := range timeFormat {
		if strings.Contains(cslayoutNew, k) {
			cslayoutNew = strings.ReplaceAll(cslayout, k, v)
		}
	}
	return cslayoutNew
}
