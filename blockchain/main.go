package main

import (
	"blockchain/initialisation"
	"fmt"
)

func main() {
	fmt.Println("Init Blockchain")
	chain := initialisation.CreateBlockchain()

	chain.AddBlock(map[string]interface{}{"from": "82173817238781273817283", "to": "126376712637126312763", "amount": 4})
	// go chain.AddBlock(map[string]interface{}{"from": "71263721673217361273672", "to": "726371627361273672163", "amount": 10})
	fmt.Println(chain)
}
