package main

import (
	"fmt"

	"github.com/user/work/foosball-booking/postgres"
)

func main() {
	driver := postgres.Driver{
		DBUser:     "postgres",
		DBPassword: "admin",
		DBName:     "foosball",
	}
	fmt.Println(driver.Initialize())
}
