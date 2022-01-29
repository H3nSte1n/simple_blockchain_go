package data

import (
	"time"
	"strings"
)

type Blockchain struct {
	Blocks []Block
	Genesis Block
	Difficulty int
}

func (chain *Blockchain) SetDifficutly(difficulty int) {
	chain.Difficulty = difficulty
}

func (chain Blockchain) mineBlock(b *Block) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", chain.Difficulty)) {
		b.PoW++
		b.Hash = b.calculateHash()
	}
}

func (chain *Blockchain) AddBlock(data map[string]interface{}) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := Block{
		PrevHash: prevBlock.Hash,
		Timestamp: time.Now(),
		Data: data,
		PoW: 0,
	}

	chain.mineBlock(&newBlock)
	chain.Blocks = append(chain.Blocks, newBlock)
}