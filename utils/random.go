package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(l int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandomIntString(l int) string {
	var intString string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		intString += fmt.Sprintf("%v", RandInt(0, 9))
	}
	return intString
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
