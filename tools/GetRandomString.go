package tools

import "math/rand"

func GetRandomString(strLen int) string {

	alph := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, strLen)
	for i := range s {
		s[i] = alph[rand.Intn(len(alph))]
	}

	return string(s)
}
