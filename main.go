package main

import (
	"fmt"

	"github.com/amamov/kyle-coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("2nd Block")
	chain.AddBlock("3th Block")
	for _, block := range chain.AllBlocks() {
		fmt.Println("-----------------------")
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
