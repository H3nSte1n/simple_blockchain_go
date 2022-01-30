package data

import (
	"time"
	"math/big"
	"fmt"
)

type Blockchain struct {
	Blocks []Block
	Genesis Block
	PoWs 		[]PoW
}

func (chain *Blockchain) AddBlock(data map[string]interface{}) {
	currentPoW := chain.PoWs[len(chain.PoWs)-1]
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := Block{
		Nr: prevBlock.Nr + 1,
		PrevHash: prevBlock.Hash,
		Timestamp: time.Now().Unix(),
		Data: data,
		Nonce: 0,
	}

	chain.mineBlock(&newBlock)
	if !newBlock.validate(currentPoW) {
		return
	}
	chain.Blocks = append(chain.Blocks, newBlock)

	if len(chain.Blocks) % 2016 == 0 {
		difficulty := *currentPoW.Difficulty
		chain.AddPoW(difficulty)
	}
}

func (chain Blockchain) mineBlock(b *Block) {
	hash := b.calculateHash()
	currentPoW := chain.PoWs[len(chain.PoWs)-1]

	for !currentPoW.proof(hash) {
		fmt.Println("mining")
		b.Nonce++
		b.Timestamp = time.Now().Unix()
		hash = b.calculateHash()
	}
}

func (chain *Blockchain) AddPoW(difficulty int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	newPoW := PoW{
		Target: target,
		Difficulty: &difficulty,
	}

	chain.PoWs = append(chain.PoWs, newPoW)
}