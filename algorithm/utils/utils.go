package utils

import (
	"crypto/rand"
	"math/big"
)

//真随机数
func Random(max int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(r.Int64())
}

//判断id是否在ids里
func IsIN(id int, ids []int) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
