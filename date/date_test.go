package date

import (
	"testing"
)

func TestFormatDate(t *testing.T) {
	d := Date{
		Year:  2018,
		Month: 9,
		Day:   28,
	}

	res, _ := d.FormatDate("dd/mm/yy")

	if res != "28/9/2018" {
		t.Error("Expected 28/9/2018 but got ", res)
	}

	res, _ = d.FormatDate("dd-MM-yy")
	if res != "28-sep-2018" {
		t.Error("Expected 28-sep-2018 but got ", res)
	}
}
