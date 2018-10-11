/*
	提供浏览器 URL 访问方式，返回数据格式为 JSON
*/
package main

import (
	"demochain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.BlockChain

func run()  {
	// 定义对 http 访问的处理
	http.HandleFunc("/blockchain/get",blockchainGetHandler)
	http.HandleFunc("/blockchain/write",blockchainWriteHandler)

	// 启动server ，监听端口
	http.ListenAndServe("localhost:8888",nil)
}

func blockchainGetHandler(w http.ResponseWriter, r *http.Request)  {
	// 转化为 json 格式
	bytes , error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter, r *http.Request)  {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w, r)
}

func main() {
	blockchain = core.NewBlockChain()
	run()
}