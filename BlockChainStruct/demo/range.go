package main

import "fmt"

func main() {
	data := []byte("hello world!")
	for i,v := range data {
		fmt.Println(i,v)
	}
	//range 只用一个是返回index
}
