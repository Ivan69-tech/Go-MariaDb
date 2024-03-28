package main

import (
	"Grafana/database"
	"Grafana/modbus"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	go modbus.StartServer()
	database.FillDbModbus()

}
