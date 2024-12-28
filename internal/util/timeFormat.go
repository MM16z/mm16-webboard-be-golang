package util

import (
	"fmt"
	"time"
)

func FormatTime(t time.Time) string {
	return t.Format("January 02 2006 - 15:04:05")
}

func FormatTimeThaiEng(t time.Time) string {
	thaiMonth := map[string]string{
		"January":   "มกราคม",
		"February":  "กุมภาพันธ์",
		"March":     "มีนาคม",
		"April":     "เมษายน",
		"May":       "พฤษภาคม",
		"June":      "มิถุนายน",
		"July":      "กรกฎาคม",
		"August":    "สิงหาคม",
		"September": "กันยายน",
		"October":   "ตุลาคม",
		"November":  "พฤศจิกายน",
		"December":  "ธันวาคม",
	}

	engFormat := t.Format("January 02 2006 - 15:04:05")
	engMonth := t.Format("January")
	thaiFormat := engFormat
	if thaiName, ok := thaiMonth[engMonth]; ok {
		thaiYear := t.Year() + 543
		thaiFormat = fmt.Sprintf("%02d %s %d - %s น.",
			t.Day(),
			thaiName,
			thaiYear,
			t.Format("15:04:05"),
		)
	}

	return thaiFormat + " / " + engFormat
}
