package date

import (
	"errors"
	"strconv"
)

var month = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

//Date type implements date data type
type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

//FormatDate accept formating string as arguments.
// Arguments should be in the form of **op**op** where op can be - or /
// and ** can be dd or yy or mm or MM. MM will return month short form
func (d Date) FormatDate(format string) (string, error) {
	result := ""
	for i := 0; i < 8; i++ {

		switch format[i : i+2] {
		case "dd":
			result += strconv.FormatInt(int64(d.Day), 10)
		case "mm":
			result += strconv.FormatInt(int64(d.Month), 10)
		case "yy":
			result += strconv.FormatInt(int64(d.Year), 10)
		case "MM":
			result += month[d.Month-1]
		default:
			return "", errors.New("Invalid format")
		}
		i = i + 2
		if i < 6 {
			result += string(format[i])
		}
	}
	return result, nil
}
