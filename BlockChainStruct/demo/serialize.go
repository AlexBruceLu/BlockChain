package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	data := []byte("这是一个要解码的数据")
	data1 := []byte{}
	buf := bytes.Buffer{}
	//buf1 := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(data)
	a := buf.Bytes()

	de := gob.NewDecoder(bytes.NewReader(a))

	de.Decode(&data1)
	fmt.Printf("要解码的数据：%s\n", data1)
}
