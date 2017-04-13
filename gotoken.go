package gotoken

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