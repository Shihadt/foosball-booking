package date

import (
	"errors"
	"strconv"
	"time"
)

//Date type implements date data type
type Date struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

//FormatDate accept formating string as arguments.
// Arguments should be in the form of **op**op** where op can be - or /
// and ** can be dd or yy or mm or MM. MM will return month short form
func (d Date) FormatDate(format string) (string, error) {

	result := ""
	for i := 0; i < 8; i++ {
		if format[i:i+1] == "dd" {
			result += strconv.FormatInt(int64(d.Day), 10)
		}
		switch format[i : i+1] {
		case "dd":
			result += strconv.FormatInt(int64(d.Day), 10)
		case "mm":
			result += strconv.FormatInt(int64(d.Month), 10)
		case "yy":
			result += strconv.FormatInt(int64(d.Year), 10)
		default:
			return "", errors.New("Invalid format")
		}
		i = i + 2
		result += string(format[i])
	}
	return result, nil
}
