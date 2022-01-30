package data

import (
	"math/big"
)

type PoW struct {
	Target				 *big.Int
	Difficulty		 *int
}

func (pow PoW) proof(hash big.Int) bool {
	return hash.Cmp(pow.Target) == -1
}