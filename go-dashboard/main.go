package main

import (
	"go-dashboard/config"
	v1 "go-dashboard/v1"
)

// "username": "mvadly",
//         "password": "pmZY39!!",

func main() {
	var db = config.DB()
	v1.EchoRoute(db)
}
