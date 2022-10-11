package utils

import (
	"crypto/rand"
	"math/big"
)

//80%概率
func EightyProbablity() bool {

	//p := rand.Intn(100)
	p := Random(100)
	if p < 80 {
		return true
	} else {
		return false
	}
}

//真随机数
func Random(max int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(r.Int64())
}
