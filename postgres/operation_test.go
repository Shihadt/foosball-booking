package postgres

import (
	"testing"
)

func TestExec(t *testing.T) {
	qry := "INSERT INTO employee(employeename) VALUES('John')"
	res, _ := DefualtDriver.Exec(qry)
	i, _ := res.RowsAffected()
	if i != 1 {
		t.Error("Cant execute the query")
	}
}

func TestQuery(t *testing.T) {
	qry := "SELECT employeename FROM employee where employeename='John'"
	res, _ := DefualtDriver.Query(qry)
	var str string
	res.Next()
	res.Scan(&str)
	if str != "John" {
		t.Error("Expected value is John but recieved ", str)
	}
}
