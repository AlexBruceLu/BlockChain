package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const WalletFile = "wallet.dat"

//定于wallets结构，用于存储所有wallet以及他的地址
type Wallets struct {
	WalletsMap map[string]*Wallet
}

//创建方法
func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	//ws.SaveToFile()
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()
	ws.WalletsMap[address] = wallet
	ws.SaveToFile()
	return address
}

//保存方法，把创建好的Wallet保存到文件中去
func (ws *Wallets) SaveToFile() {
	var buffer bytes.Buffer
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic("保存wallet时，序列化失败")
	}
	ioutil.WriteFile(WalletFile, buffer.Bytes(), 0600)
}

//读取方法，把所有的wallet从文件中读出来
func (ws *Wallets) LoadFile() {
	_, err := os.Stat(WalletFile)
	if os.IsExist(err) {
		return
	}
	content, err := ioutil.ReadFile(WalletFile)
	if err != nil {
		log.Panic("读取钱包文件失败")
	}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	var walletsTmp Wallets
	err = decoder.Decode(&walletsTmp)
	if err != nil {
		log.Panic("读取钱包文件，反序列化失败")
	}
	ws.WalletsMap = walletsTmp.WalletsMap
}

//显示所有地址
func (ws *Wallets) ListAllAddresses() []string {
	var addresses []string
	for address := range ws.WalletsMap {
		addresses = append(addresses, address)
	}
	fmt.Println(addresses)
	return addresses
}
