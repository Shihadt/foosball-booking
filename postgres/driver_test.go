package postgres

import "testing"

var d = Driver{
	DBUser:     "postgres",
	DBPassword: "admin",
	DBName:     "test",
}

func TestConnect(t *testing.T) {

	err := d.Connect()
	if err != nil {
		t.Error("Cant connect db")
	}
}

func TestInitilize(t *testing.T) {
	err := d.Initialize()
	if err != nil {
		t.Error("Cant initialize")
	}
}
