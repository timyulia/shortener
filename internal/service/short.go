package service

import (
	"crypto/sha256"
	"strings"
)

const (
	forbidden = 63

	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func generateShortURL(long string) string {
	hash := sha256.Sum224([]byte(long))
	res := strings.Builder{}
	j := 0
	for i := 0; i < 10; i++ {
		curr := hash[i] & ((1 << 6) - 1)
		jOld := j
		for curr == forbidden {
			curr ^= hash[j+10] & ((1 << 6) - 1)
			j = (j + 1) % 18
			if j == jOld && curr == forbidden {
				curr ^= 51
			}
		}
		symbol := alphabet[curr]
		res.WriteByte(symbol)
	}
	return res.String()
}

//todo check if there are only zero values in the end
