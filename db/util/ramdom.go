package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "qwertyuioasdfghjklzxcvbnm"

func init() {
	rand.Intn(int(time.Now().UnixNano()))
}

func RamdomInt(min, max int64) int64 {
	// return min + rand.Int63n(min-max+1)
	return rand.Int63n(max-min+1) + min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomView() int64 {
	return RamdomInt(1, 1000)
}
func RandomPrice() int64 {
	return RamdomInt(10000, 1000000)
}

// Hàm chuyển đổi số nguyên sang chuỗi để tránh lỗi
func RamdomDiscount() string {
	str := strconv.Itoa(int(RamdomInt(1, 100)))
	return str + "%"
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}
func RandomPhoneNumber() string {
	return strconv.Itoa(rand.Intn(10))
}
