package postgres

// CreateDatabaseIfNotExist is to create db if its not exists.
// This function should call at begining of the program execution.
// func (d *Driver) createDatabaseIfNotExist() error {
// 	qry := "CREATE DATABASE IF NOT EXISTS " + d.DBName + ";"
// 	_, err := d.driver.Exec(qry)
// 	return err
// }

// func (d *Driver) useDataBase() error {
// 	qry := "USE " + d.DBName
// 	_, err := d.driver.Exec(qry)
// 	return err
// }

// CreateTableIfNotExits will create all tables for the project if not exists.
// This function should call at begining of the program execution.
func (d *Driver) createTableIfNotExits() error {
	qry := ` CREATE TABLE IF NOT EXISTS employee (
		employeeID serial PRIMARY KEY,
		employeeName VARCHAR(50)
	);`
	if _, err := d.driver.Exec(qry); err != nil {
		return err
	}
	qry = `CREATE TABLE IF NOT EXISTS admin(
		adminID serial PRIMARY KEY,
		adminName VARCHAR(50)
	);`
	if _, err := d.driver.Exec(qry); err != nil {
		return err
	}
	qry = `CREATE TABLE IF NOT EXISTS request(
		requestID serial PRIMARY KEY,
		userID INTEGER NOT NULL,
		timeSlot INTEGER NOT NULL,
		isAllocated BOOLEAN
	);`
	if _, err := d.driver.Exec(qry); err != nil {
		return err
	}
	qry = `CREATE TABLE IF NOT EXISTS slot(
		slotID serial PRIMARY KEY,
		request INTEGER 
	);`
	if _, err := d.driver.Exec(qry); err != nil {
		return err
	}
	return nil
}
