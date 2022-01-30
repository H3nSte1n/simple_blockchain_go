package initialisation

import (
	"blockchain/data/block"
	chain "blockchain/data/blockchain"
	"blockchain/data/pow"
	"blockchain/db"
	"time"
)

func ResetCollection() {
	err_block_c := db.ResetBlockCollection()
	err_pow_c := db.ResetPoWCollection()

	if err_block_c != nil || err_pow_c != nil {
		panic("Error resetting collections")
	}
}

func CreateBlockchain() chain.Blockchain {
	ResetCollection()

	genesis := block.Block{
		Nr:        0,
		Hash:      "0",
		Timestamp: time.Now().Unix(),
		Nonce:     0,
	}

	chain := chain.Blockchain{
		Genesis: genesis,
		Blocks:  []block.Block{genesis},
		PoWs:    []pow.PoW{},
	}

	difficulty := 24
	chain.AddPoW(difficulty)
	return chain
}
