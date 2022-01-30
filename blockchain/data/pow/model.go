package pow

import "math/big"

type PoW struct {
	Nr         int      `json:"nr" bson:"_id,omitempty"`
	Target     *big.Int `json:"target"`
	Difficulty *int     `json:"difficulty"`
}
