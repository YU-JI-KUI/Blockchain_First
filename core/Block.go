package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	/* 一个区块包含区块头和区块体两部分，这里的区块体用 Data 来简化表示，实际的业务中，是由多条交易构成的，要复杂的多。*/

	Index int64 // 区块编号
	Timestamp int64 // 区块时间戳（创建时间）
	PrevBlockHash string // 上一个区块的 hash 值
	Hash string // 当前区块的 hash 值

	Data string // 区块数据
}


/*
	计算区块的 hash 值
*/
func calculateHash(b Block) string {
	blockData := b.Data + b.PrevBlockHash + string(b.Index) + string(b.Timestamp)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}


/*
	生成一个新的区块
*/
func GenerateNewBlock(preBlock Block,data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Data = data
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}


/*
	创建 "创世区块"
 */
func GenerateGenesisBlock() Block {
	newBlock := Block{}
	newBlock.Index = -1
	newBlock.Hash = ""
	return GenerateNewBlock(newBlock,"Genesis Block")
}
