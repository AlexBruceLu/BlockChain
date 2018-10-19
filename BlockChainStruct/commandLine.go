package main

import (
	"fmt"
	"time"
)

//func (c *CLI) AddBlock(txs []*Transaction) {
//	c.BC.AddBlocks(txs)
//	fmt.Println("区块添加成功！")
//}

func (c *CLI) PrintChain() {
	//正向打印区块链
	c.BC.PrintBlockChain()
	fmt.Println("区块正向打印完毕！")
}

func (c *CLI) PrintChainR() {
	//反向打印区块
	it := c.BC.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("==========================\n\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerKelRoot)
		tmpTime := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳: %d\n", tmpTime)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Transaction[0].TxInputs[0].Sig)
		if len(block.PrevHash) == 0 {
			fmt.Println("区块反向打印完毕！")
			break
		}
	}
}

//获取余额
func (c *CLI) GetBalance(address string) {
	utxos := c.BC.FindUTXOs(address)
	total := 0.0
	if utxos == nil {
		fmt.Println("地址错误，请查证！")
	}
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("\"%s\"的余额为：%f\n", address, total)
}

//转账
func (c *CLI) Send(from, to string, amount float64, miner, data string) {
	//创建挖矿交易
	coinBase := NewCoinBaseTx(miner, data)
	//创建普通交易
	tx := NewTransaction(from, to, amount, c.BC)
	if tx == nil {
		return
	}
	//添加区块
	c.BC.AddBlocks([]*Transaction{coinBase, tx})
	fmt.Println("转账成功")
}
