package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		bytes, err := hex.DecodeString(arg)
		if err != nil {
			log.Fatalln(err)
		}

		length := len(bytes)
		fmt.Print("{ ")
		for idx, bytedata := range bytes {
			fmt.Printf("0x%02X", bytedata)
			if length-idx > 1 {
				fmt.Print(", ")
			}
		}
		fmt.Println(" }")
	}
}
