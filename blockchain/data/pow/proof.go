package pow

import "math/big"

func (pow *PoW) Proof(hash big.Int) bool {
	return hash.Cmp(pow.Target) == -1
}
