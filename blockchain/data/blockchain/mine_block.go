package chain

import (
	"time"
	"blockchain/data/block"
)

func (chain Blockchain) mineBlock(b *block.Block) {
	hash := b.CalculateHash()
	currentPoW := chain.PoWs[len(chain.PoWs)-1]

	for !currentPoW.Proof(hash) {
		b.Nonce++
		b.Timestamp = time.Now().Unix()
		hash = b.CalculateHash()
	}
}