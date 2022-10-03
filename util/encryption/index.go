package encryption

import (
	"go-dashboard/util"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func DefaultPassword(username string) string {
	var defaultPass = username + "P@ssw0rd123"
	pass, _ := HashPassword(defaultPass)
	return pass
}

func GeneratePassword() (str string) {
	var lower = []rune("abcdefghijklmnopqrstuvwxyz")
	var upper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var number = []rune("1234567890")
	var symbol = []rune(`@.!#_?*`)
	rand.Seed(time.Now().UnixNano())
	lowers := util.RandStr(2, lower)
	uppers := util.RandStr(2, upper)
	numbers := util.RandStr(2, number)
	symbols := util.RandStr(2, symbol)
	strs := lowers + uppers + numbers + symbols
	return strs
}
