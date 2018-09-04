package postgres

import (
	"database/sql"
)

//Exec execute given query
func (d *Driver) Exec(qry string) (sql.Result, error) {
	return d.driver.Exec(qry)
}

// Query will retrieve resulted values
func (d *Driver) Query(qry string) (*sql.Rows, error) {
	return d.driver.Query(qry)
}

//Prepare generally used for update informations
func (d *Driver) Prepare(qry string) (*sql.Stmt, error) {
	return d.driver.Prepare(qry)
}
