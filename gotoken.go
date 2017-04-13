package main

import (
	"math/rand"
	"time"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var rangeSize = len(chars)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Token interface {
	Generate() string
}

func GenerateToken(length int) string {
	bytes := make([]rune, length)
	for i := range(bytes) {
		bytes[i] = chars[rand.Intn(rangeSize)]
	}
	return string(bytes)
}