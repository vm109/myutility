package main

import (
	"fmt"
	"myutility.com/m/v2/pkg/server"
)

func main() {
	// Entry point to timezone_converter
	fmt.Println("Starting time zone converter")
	server.NewTimeZoneServer(8080, "timezone").Start()
}
