package main

import (
	b "blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := b.InitBlockChain()

	chain.AddBlock("Second Block after First")
	chain.AddBlock("Third Block after First")
	chain.AddBlock("Forth Block after First")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := b.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
