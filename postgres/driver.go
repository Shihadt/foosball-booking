package postgres

import (
	"database/sql"

	// For initializing postgres
	_ "github.com/lib/pq"
)

//CustomDriver implementation give extendibility for using your custom driver
type CustomDriver interface {
	Connect() *sql.DB
}

//Driver configuration
type Driver struct {
	DBUser     string
	DBPassword string
	DBName     string
	driver     *sql.DB
}

//DefualtDriver for postgres
var DefualtDriver = Driver{
	DBUser:     "postgres",
	DBPassword: "postgres",
}

func init() {
	DefualtDriver.connect()
}

//Connect Database and return DB instance
func (d *Driver) connect() error {
	qry := "user=" + d.DBUser + " password=" + d.DBPassword
	if d.DBName != "" {
		qry += " dbname=" + d.DBName
	}
	qry += " sslmode=disable"

	db, err := sql.Open("postgres", qry)
	d.driver = db

	return err
}

//Close the driver
func (d *Driver) Close() {
	d.driver.Close()
}

//Initialize driver
func (d *Driver) Initialize() error {
	if err := d.connect(); err != nil {
		return err
	}
	// if err := d.createDatabaseIfNotExist(); err != nil {
	// 	return err
	// }
	// if err := d.useDataBase(); err != nil {
	// 	return err
	// }
	if err := d.createTableIfNotExits(); err != nil {
		return err
	}
	return nil
}
