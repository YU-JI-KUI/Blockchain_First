/*
	命令行直接测试
 */
package main

import "demochain/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send ! EOS to Jack")
	bc.Print()
}
