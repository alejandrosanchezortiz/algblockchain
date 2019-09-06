package blockchain

import (
	"fmt"
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

func AddBlock( bc *BlockChain, s string){
    prevBlock := bc.blocks[len(bc.blocks)-1]
    new := CreateBlock (s,prevBlock.Hash)
    bc.blocks = append (bc.blocks, new)
}