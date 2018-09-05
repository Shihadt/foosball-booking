package postgres

import "testing"

func TestCreateTableIfNotExits(t *testing.T) {

	err := DefualtDriver.createTableIfNotExits()
	if err != nil {
		t.Error("Error while executing create table")
	}
}
