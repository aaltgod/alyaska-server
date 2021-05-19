package tools

import (
	"math/rand"
	"time"
)

func GetRandomString(strLen int) string {

	rand.Seed(time.Now().UnixNano())

	alph := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	result := make([]rune, strLen)
	for i := range result {
		result[i] = alph[rand.Intn(len(alph))]
	}

	for i := range result {
		result[i] = alph[rand.Intn(len(result))]
	}

	return string(result)
}
