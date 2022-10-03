package main

import (
	"go-dashboard/config"
	v1 "go-dashboard/v1"
)

// "username": "mvadly",
//         "password": "jqQY22._"

func main() {
	var db = config.DB()
	v1.EchoRoute(db)
}
