package main

import (
        "fmt"
        "github.com/alejandrosanchezortiz/algblockchain/blockchain"
)

func main(){
    var chain *blockchain.BlockChain
    chain = blockchain.InitBlockChain("Primero")

    AddBlock(chain,"Segundo")
    AddBlock(chain,"Tercero")

    for _, block := range chain.blocks { // index,element we don't need index, so we use '_'
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Printf("**************************************\n")
	}
}