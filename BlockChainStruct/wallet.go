package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"io/ioutil"
	"log"
	"os"
)

//定义钱包结构
type Wallet struct {
	//私钥
	Private   *ecdsa.PrivateKey
	PublicKey []byte
	//公钥，并不是原始公钥，而是r，s拼接而成的
}

func NewWallet() *Wallet {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic("创建密钥对失败")
	}
	publicKeyTmp := privateKey.PublicKey
	publicKey := append(publicKeyTmp.X.Bytes(), publicKeyTmp.Y.Bytes()...)
	return &Wallet{Private: privateKey, PublicKey: publicKey}
}

//生成地址
func (w *Wallet) NewAddress() string {
	pubkey := w.PublicKey
	rp160Hash := HashPublicKey(pubkey)
	version := byte(00)
	payload := append([]byte{version}, rp160Hash...)
	checkcode := CheckSum(payload)
	payload = append(payload, checkcode...)
	address := base58.Encode(payload)
	return address
}

//创建钱包 判断
func (ws *Wallets) LoadWalletFromFile() {
	_, err := os.Stat(WalletFile)
	if os.IsNotExist(err) {
		return
	}
	content, err := ioutil.ReadFile(WalletFile)
	ErrorFunc(err, "ioutil.ReadFile")
	var wsLocal Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	err = decoder.Decode(&wsLocal)
	ErrorFunc(err, "decoder.Decode")
	ws.WalletsMap = wsLocal.WalletsMap
}

//ErrorFunc
func ErrorFunc(err error, data string) {
	if err != nil {
		log.Panic(data)
	}
}

//哈希公钥
func HashPublicKey(data []byte) []byte {
	hash := sha256.Sum256(data)
	rp160Hasher := ripemd160.New()
	_, err := rp160Hasher.Write(hash[:])
	if err != nil {
		log.Panic("ripemd160 err")
	}
	rp160HashValue := rp160Hasher.Sum(nil)
	return rp160HashValue
}

//识别检测位
func CheckSum(data []byte) []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	//前面4字节校验码
	checkCode := hash2[:4]
	return checkCode
}
