package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type ResJSON struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSON(e echo.Context, statusCode int, res ResJSON) error {
	return e.JSON(statusCode, res)
}

func JsonEncode(v interface{}) string {
	d, _ := json.Marshal(v)
	return string(d)
}

func JsonDecode(v string) map[string]interface{} {
	var maps map[string]interface{}
	json.Unmarshal([]byte(v), &maps)
	return maps
}

func RandStr(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func TimeNow() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc)
}

func ViewRoutes(route []*echo.Route) {
	for _, v := range route {
		fmt.Println(v.Method, "---", os.Getenv("APP_URL")+v.Path, "---", v.Name)
	}
}
