package main

import "math/rand"

type AlphanumericToken struct {
	chars []rune
	rangeSize int
	length int
}

func NewAlphanumericTokenGenerator (length int) *AlphanumericToken {
	return &AlphanumericToken{
		chars: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"),
		rangeSize: len(chars),
		length: length,
	}
}

func (at *AlphanumericToken) Generate() string {
	bytes := make([]rune, at.length)
	for i := range(bytes) {
		bytes[i] = at.chars[rand.Intn(at.rangeSize)]
	}
	return string(bytes)
}


