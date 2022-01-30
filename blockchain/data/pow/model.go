package pow

import "math/big"

type PoW struct {
	Target				 *big.Int  	`json:"target"`
	Difficulty		 *int				`json:"difficulty"`
}