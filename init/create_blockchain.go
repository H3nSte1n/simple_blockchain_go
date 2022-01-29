package init

import (
	"blockchain/data"
	"time"
)

func CreateBlockchain() data.Blockchain {
	genesis := data.Block{
		Hash: "0",
		Timestamp: time.Now(),
	}

	return data.Blockchain {
		Genesis: genesis,
		Blocks: []data.Block{genesis},
		Difficulty: 3,
	}
}