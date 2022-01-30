package block

import (
	"blockchain/data/pow"
)

func (b Block) Validate(pow pow.PoW) bool {
	return pow.Proof(b.CalculateHash())
}
