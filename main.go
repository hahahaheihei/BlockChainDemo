package main

import (
	"blockDemo/model"
	"fmt"
	"strconv"
)

func main(){
	bc:=model.NewBlockchain()
	bc.AddBlock("send 1 btc to cong")
	bc.AddBlock("send 2 more btc to cong")
	for _,block:=range bc.Blocks{
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow:=model.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
