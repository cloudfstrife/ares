package utils

import (
	"crypto/rand"
	"math/big"
)

//Judge 跟据概念返回事件发生与否
func Judge(probability int64) bool {
	if probability <= 0 {
		return false
	}
	if probability >= 100 {
		return true
	}
	v, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return false
	}
	return probability >= v.Int64()
}
