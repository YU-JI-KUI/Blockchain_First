package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block   // 为什么加 * ？
}


/*
	创建区块链
*/
func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}


/*
	将数据直接添加到区块链中
*/
func (bc *BlockChain)SendData(date string){
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateNewBlock(*preBlock,date)
	bc.AppendBlock(&newBlock)
}


/*
	添加区块到区块链中
*/
func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks,newBlock)
		return
	}
	// 对新区块进行校验
	if isValid(*newBlock,*bc.Blocks[len(bc.Blocks) - 1]) {
		// 将 Block 添加到 BlockChain 上
		bc.Blocks = append(bc.Blocks,newBlock)
	} else {
		log.Fatal("invalid block !")
	}
}


/*
	校验区块是否合法
*/
func isValid(newBlock Block,oldBlock Block) bool  {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}


/*
	打印区块链
*/
func (bc *BlockChain) Print(){
	for _ , block := range bc.Blocks {
		fmt.Printf("Index : %d \n" , block.Index)
		fmt.Printf("Prev.Hash : %s \n" , block.PrevBlockHash)
		fmt.Printf("Curr.Hash : %s \n" , block.Hash)
		fmt.Printf("Data : %s \n" , block.Data)
		fmt.Printf("Timestamp : %d \n" , block.Timestamp)
		fmt.Println()
	}
}