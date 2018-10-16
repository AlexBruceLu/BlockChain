package main

import (
	"fmt"
	"github.com/bolt"
	"log"
)

//
func main() {
	//1.打开数据库
	db, err := bolt.Open("test.db", 0600, nil) //参数2：打开文件的权限0600读写权限
	if err != nil {
		log.Panic("打开数据库失败")
	}
	defer db.Close()
	//2.操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屉如果没有就创建
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("数据库bucket创建失败")
			}
		}
		bucket.Put([]byte("11111"), []byte("hello"))
		bucket.Put([]byte("22222"), []byte("world"))

		return nil
	})
	//3.读取数据库数据
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket不应该为空，请核对")
		}
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))
		fmt.Printf("v1	:%s\n", v1)
		fmt.Printf("v2	:%s\n", v2)
		return nil
	})
}
