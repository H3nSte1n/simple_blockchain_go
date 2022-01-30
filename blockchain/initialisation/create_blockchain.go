package initialisation

import (
	"blockchain/data/block"
	"blockchain/data/pow"
	"blockchain/data/blockchain"
	"time"
)

func CreateBlockchain() data.Blockchain {
	genesis := data.Block{
		Nr: 0,
		Hash: "0",
		Timestamp: time.Now().Unix(),
		Nonce: 0,
	}

	chain := data.Blockchain{
		Genesis: genesis,
		Blocks: []data.Block{genesis},
		PoWs: []data.PoW{},
	}

	difficulty := 3
	chain.AddPoW(difficulty)
	return chain
}