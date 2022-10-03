package util

import (
	"math/rand"

	"github.com/labstack/echo"
)

type ResJSON struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(e echo.Context, statusCode int, res ResJSON) {
	e.JSON(statusCode, res)
	return
}

func RandStr(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
