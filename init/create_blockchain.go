package init

import (
	"blockchain/data"
	"time"
)

func CreateBlockchain() data.Blockchain {
	firstBlock := data.Block{
		Hash: "0",
		Timestamp: time.Now(),
	}

	return data.Blockchain {
		FirstBlock: firstBlock,
		Blocks: []data.Block{firstBlock},
		Difficulty: 3,
	}
}