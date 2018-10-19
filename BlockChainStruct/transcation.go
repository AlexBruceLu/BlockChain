package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const Reward = 50

//定义交易结构体
type Transaction struct {
	TxId      []byte     //交易Id
	TxInputs  []TxInput  //交易输入数组
	TxOutPuts []TxOutput //交易输出数组
}

//定义结构交易输入结构
type TxInput struct {
	TxId  []byte //引用的交易ID
	Index int64  //引用的output的索引值
	Sig   string //解锁脚本，我们用地址来模拟
}

//定义交易结构输出结构
type TxOutput struct {
	Value      float64 //转账金额
	PubKeyHash string  //锁定脚本,我们用地址模拟
}

//实现交易哈希方法
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic("交易序列化失败")
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TxId = hash[:]
}

//实现挖矿交易
func NewCoinBaseTx(address, data string) *Transaction {
	//挖矿交易的特点：
	//1. 只有一个input
	//2. 无需引用交易id
	//3. 无需引用index
	//矿工由于挖矿时无需指定签名，所以这个sig字段可以由矿工自由填写数据，一般是填写矿池的名字
	input := TxInput{[]byte{}, -1, data}
	output := TxOutput{Reward, address}
	tx := Transaction{[]byte{}, []TxInput{input}, []TxOutput{output}}
	tx.SetHash()
	return &tx
}

//实现普通转账交易
//创建普通交易，引入输出
//如果有找零，需要找零
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	//找到最合理的UTXO集合
	utxos, resValue := bc.FindNeedUTXOs(from, amount)
	if resValue < amount {
		fmt.Println("余额不足，交易失败")
		return nil
	}
	var inputs []TxInput
	var outputs []TxOutput
	//创建交易输入，将utxo转化成inputs
	for id, indexArr := range utxos {
		for _, i := range indexArr {
			input := TxInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}
	//创建交易输出
	output := TxOutput{amount, to}
	outputs = append(outputs, output)
	//找零
	if resValue > amount {
		outputs = append(outputs, TxOutput{resValue - amount, from})
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}

//实现判断是否为挖矿交易的方法
func (tx *Transaction) IsCoinBase() bool {
	//是否为挖矿交易满足以下条件
	//1. input只有一个
	//2. 交易ID为空
	//3. 交易的Index为空
	if len(tx.TxInputs) == 1 && len(tx.TxInputs[0].TxId) == 0 && tx.TxInputs[0].Index == -1 {
		return true
	}
	return false
}
