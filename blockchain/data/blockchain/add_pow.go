package chain

import (
	"blockchain/data/pow"
	"blockchain/db"
	"math/big"
)

func (chain *Blockchain) AddPoW(difficulty int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	newPoW := pow.PoW{
		Nr:         len(chain.PoWs) + 1,
		Target:     target,
		Difficulty: &difficulty,
	}

	db.InsertPoW(&newPoW)
	chain.PoWs = append(chain.PoWs, newPoW)
}
