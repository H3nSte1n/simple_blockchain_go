package main

import (
	"fmt"
	"blockchain/initialisation"
)

func main() {
	fmt.Println("Init Blockchain")
	chain := initialisation.CreateBlockchain()

	chain.AddBlock(map[string]interface{}{"from": "82173817238781273817283", "to": "126376712637126312763", "amount": 4})
	fmt.Println(chain)
}