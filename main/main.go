package main

import (
	"fmt"

	"github.com/tensor-programming/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block aftrer Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Println()
		if len(block.PrevHash) < 1 {
			fmt.Printf("Previous Hash: null\n")
		} else {
			fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		}
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}
}
