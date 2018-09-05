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
	DBPassword: "admin",
	DBName:     "foosball",
}

func init() {
	DefualtDriver.Connect()
}

//Connect Database and return DB instance
func (d *Driver) Connect() error {
	qry := "user=" + d.DBUser + " password=" + d.DBPassword
	if d.DBName != "" {
		qry += " dbname=" + d.DBName
	}
	qry += " sslmode=disable"

	db, _ := sql.Open("postgres", qry)
	d.driver = db
	return d.driver.Ping()
}

//Close the driver
func (d *Driver) Close() {
	d.driver.Close()
}

//Initialize driver
func (d *Driver) Initialize() error {
	if err := d.Connect(); err != nil {
		return err
	}
	if err := d.createTableIfNotExits(); err != nil {
		return err
	}
	return nil
}
