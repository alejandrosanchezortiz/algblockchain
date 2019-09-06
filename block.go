package main

import (
        "fmt"
        "bytes"
        "crypto/sha256"
)

type Block struct {
    Hash []byte
    Data []byte
    PrevHash []byte
}

type BlockChain struct{
    blocks []*Block
}

func getHash(b *Block){
    info := bytes.Join([][]byte{b.Data,b.PrevHash},[]byte{}) // https://golang.org/pkg/bytes/#Join func Join(s [][]byte, sep []byte) []byte
    hash := sha256.Sum256(info)
    b.Hash = hash[:] // ':=' declares a new variable, '=' just assigns
}

func CreateBlock(data string,prevHash []byte) *Block{
    block := &Block{[]byte{}, []byte(data), prevHash} //pointer, constructor
    getHash(block)
    return block
}

func InitBlockChain(s string) *BlockChain {  
    fmt.Println("**InitBlockChain**")
    return &BlockChain{[]*Block{CreateBlock(s,[]byte{})}}
}

func main(){
    var chain *BlockChain
    chain = InitBlockChain("Primero")

    for _, block := range chain.blocks { // index,element we don't need index, so we use '_'
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}