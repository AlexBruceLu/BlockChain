package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block *Block
	//目标值非常大的数
	target *big.Int
}

//提供创建pow的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//指定一个难度值，但类型需要从string转为big.Int
	//sha256计算出哈希值32字节，要与之匹配(十六进制两个数一字节)
	tmpTarget := "0000f00000000000000000000000000000000000000000000000000000000000"
	//引入辅助变量将上面难度值转化为big.Int
	tmpInt := big.Int{}
	//将难度值赋值给辅助变量，使用16进制的格式
	tmpInt.SetString(tmpTarget, 16)
	//fmt.Printf("%x\n",tmpInt)
	pow.target = &tmpInt
	//fmt.Printf("%x\n",tmpInt)
	//fmt.Println(pow.target)
	return &pow
}

//提供不断计算哈希的函数
//1. 拼装数据（区块的数据，还有不断变化的随机数）
//2. 做哈希运算
//3. 与pow中的target比较
//a. 找到了返回退出
//b. 没找到nonce加一，继续寻找
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64 =0
	var hash [32]byte
	block := pow.block
	//fmt.Printf("难度：%x/n",*pow.target)
	for {
		tmp := [][]byte{
			block.Hash,
			Uint64ToByte(block.Version),
			Uint64ToByte(nonce),//用定义的nonce值
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			block.MerKelRoot,
			block.Data,
			block.PrevHash,
		}
		//fmt.Println(nonce)
		//将区块各个字段连接起来准备做哈希运算
		blockInfo := bytes.Join(tmp, []byte{})
		//中间变量，方便将哈希值转成big.Int
		hash = sha256.Sum256(blockInfo)
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		//fmt.Println(tmpInt)
		//fmt.Println(tmpInt.Cmp(pow.target))
		//比较当前的哈希与目标哈希值，如果当前的哈希值小于目标的哈希值，就说明找到了，否则继续找
		//   -1 if x <  y  比pow.target哈希值小即为找到
		//    0 if x == y
		//   +1 if x >  y
		//func (x *Int) Cmp(y *Int) (r int) 
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功 hash：%x,nonce：%d\n",hash,nonce)
			return hash[:],nonce
		} else {
			nonce++
		}
		//return hash[:],nonce
	}
}
