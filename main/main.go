package main

import (
	"fmt"
	"strconv"

	"github.com/tensor-programming/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block aftrer Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		if len(block.PrevHash) < 1 {
			fmt.Printf("Previous Hash: None\n")
		} else {
			fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		}
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
