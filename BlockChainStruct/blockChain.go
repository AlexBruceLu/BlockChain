package main

import (
	"bytes"
	"fmt"
	"github.com/bolt"
	//"github.com/btcsuite/btcutil"
	"log"
)

//定义一个区块链数组
type BlockChain struct {
	//Blocks []*Block
	//存储哈希值，哈希池
	Db *bolt.DB
	//用于存储最后一笔交易的哈希值
	Tail []byte
}

//定义常量，哈希池
const BlockChainDB = "blockChain.db"

//定义常量，bucket
const BlockBucket = "BlockBucket"

//定义最后一个哈希值的，db key
const LastHashKey = "LastHashKey"

//定义一个区块链
func NewBlockChain(address string) *BlockChain {
	//创建创世块，添加到区块链中
	//genesisBlock := GenesisBlock()
	//return &BlockChain{
	//	Blocks: []*Block{genesisBlock},
	//}
	//定义最后一个区块的哈希值，从数据库中读出来
	var lastBlockHash []byte
	//1. 打开数据库
	db, err := bolt.Open(BlockChainDB, 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败")
	}
	//2. 操作数据库，添加创世块信息
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket == nil {
			//如果bucket为空就要创建一个bucket，并写入创世块
			bucket, err = tx.CreateBucket([]byte(BlockBucket))
			if err != nil {
				log.Panic("bucket创建失败")
			}
			//创世块写入bucket
			genesisBlock := GenesisBlock(address)
			// 存储创世块的哈希值，实现序列化方法
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte(LastHashKey), genesisBlock.Hash)
			lastBlockHash = genesisBlock.Hash
		} else {
			lastBlockHash = bucket.Get([]byte(LastHashKey))
		}
		return nil
	})
	//3. 返回当前数据库
	return &BlockChain{db, lastBlockHash}
}

//定义一个创世块
func GenesisBlock(address string) *Block {
	//return NewBlock([]byte{}, "这是第一个创世块")
	tx := NewCoinBaseTx(address, "这是第一个创世块")
	return NewBlock([]byte{}, []*Transaction{tx})
}

//添加区块
func (bc *BlockChain) AddBlocks(txs []*Transaction) {
	//获取最后一个区块，得到前区块的哈希值
	//lastBlock := bc.Blocks[len(bc.Blocks)-1]
	//prevHash := lastBlock.Hash
	//block := NewBlock(prevHash, data)
	////添加新区块到区块链上
	//bc.Blocks = append(bc.Blocks, block)
	//--------------------------------------------------

	//获得区块数据库，和最后一个区块哈希值
	db := bc.Db
	lastHash := bc.Tail
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket == nil {
			log.Panic("BlockBucket 添加时不能为空请检查")
		}
		block := NewBlock(lastHash, txs)

		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(LastHashKey), block.Hash)
		bc.Tail = block.Hash
		return nil
	})
}

//打印区块
func (bc *BlockChain) PrintBlockChain() {
	blockHeight := 0
	bc.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		//从第一个key-value开始遍历，知道最后一个退出
		bucket.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte(LastHashKey)) {
				return nil
			}
			block := Deserialize(v)
			fmt.Printf("==============当前区块高度：%d================\n", blockHeight)
			fmt.Printf("版本号：%d", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerKelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.Hash)
			fmt.Printf("区块数据 :%s\n", block.Transaction)
			blockHeight++
			return nil
		})
		return nil
	})
}

//找出指定地址的所有UTXO
func (bc *BlockChain) FindUTXOs(address string) []TxOutput {
	var UTXO []TxOutput
	txs := bc.FindUTXOTransaction(address)
	for _, tx := range txs {
		for _, output := range tx.TxOutPuts {
			if address == output.PubKeyHash {
				UTXO = append(UTXO, output)
			}
		}
	}
	return UTXO
}

//根据需求找出合理的UTXO
func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	utxos := make(map[string][]uint64)
	var calc float64
	txs := bc.FindUTXOTransaction(from)
	for _, tx := range txs {
		for i, output := range tx.TxOutPuts {
			if from == output.PubKeyHash {
				if calc < amount {
					utxos[string(tx.TxId)] = append(utxos[string(tx.TxId)], uint64(i))
					calc += output.Value
					if calc >= amount {
						fmt.Printf("找到了满足条件的金额：%f\n", calc)
						return utxos, calc
					}
				}
			} else {
				fmt.Printf("金额不足，转账金额：%f,账户余额：%f\n", amount, utxos)
			}
		}
	}
	return utxos, calc
}

//找出所有的UTXO交易
func (bc *BlockChain) FindUTXOTransaction(address string) []*Transaction {
	var txs []*Transaction //存储所有包含UTXO的交易信息
	//定义一个map来保存未消费过的output，key为output的交易ID，value为交易引索数组
	spendOutputs := make(map[string][]int64)
	it := bc.NewIterator()
	for {
		//1. 遍历区块
		block := it.Next()
		//2. 遍历交易
		for _, tx := range block.Transaction {
			//遍历output找出，和自己相关的UTXO，添加之前检查是否已经消耗
		OUTPUT:
			for i, output := range tx.TxOutPuts {
				//剔除交易里面已经消耗的UTXO
				if spendOutputs[string(tx.TxId)] != nil {
					for _, j := range spendOutputs[string(tx.TxId)] {
						if int64(i) == j {
							continue OUTPUT
						}
					}
				}
				//经检查目标地址是否与output相同，如果满足条件再追加UTXO
				if output.PubKeyHash == address {
					txs = append(txs, tx)
				}
			}
			//如果是挖矿交易就直接跳过，不做遍历
			if !tx.IsCoinBase() {
				//遍历input变自己花费过的UTXO表示出来
				for _, input := range tx.TxInputs {
					//判断目标值与当前input是否一致
					if input.Sig == address {
						spendOutputs[string(input.TxId)] = append(spendOutputs[string(input.TxId)], input.Index)
					}
				}
			}
		}
		if len(block.PrevHash) == 0 {
			fmt.Println("区块遍历结束")
			break
		}
	}
	return txs
}
