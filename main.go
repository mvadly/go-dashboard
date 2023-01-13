package main

import (
	"fmt"
	"go-dashboard/config"
	v1 "go-dashboard/v1"
)

// "username": "mvadly",
//         "password": "pmZY39!!",

func main() {
	fmt.Println("APP STARTED")
	var db = config.DB()
	v1.EchoRoute(db)
}
