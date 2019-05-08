package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	var stringbuffer bytes.Buffer

	for _, arg := range args {
		stringbuffer.WriteString(arg)
	}

	bytesData, err := hex.DecodeString(stringbuffer.String())

	if err != nil {
		log.Fatalln(err)
	}

	length := len(bytesData)
	fmt.Print("{ ")
	for idx, bytedata := range bytesData {
		fmt.Printf("0x%02X", bytedata)
		if length-idx > 1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(" }")
}
