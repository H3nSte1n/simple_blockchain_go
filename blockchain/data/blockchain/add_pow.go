package chain

import (
	"math/big"
	"blockchain/data/pow"
	"blockchain/db"
)

func (chain *Blockchain) AddPoW(difficulty int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	newPoW := pow.PoW{
		Target: target,
		Difficulty: &difficulty,
	}

	db.InsertPoW(&newPoW)
	chain.PoWs = append(chain.PoWs, newPoW)
}