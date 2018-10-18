package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

func main() {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic()
	}
	publicKey := privateKey.PublicKey
	data := []byte("这是一个加密数据")
	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic("签名失败")
	}
	fmt.Printf("publicKey：%v\n", publicKey)
	fmt.Printf("r ：%s,len(r)：%d\n", r, len(r.Bytes()))
	fmt.Printf("s ：%s,len(s)：%d\n", s, len(s.Bytes()))
	signature := append(r.Bytes(), s.Bytes()...)

	var r1, s1 big.Int
	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])

	res := ecdsa.Verify(&publicKey, hash[:], &r1, &s1)
	fmt.Printf("签名验证结果：%t\n", res)

}
