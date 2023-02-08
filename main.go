package main

import (
	"fmt"
	"golang-cmgt-coin-miner/mod10"
)

func Hash(secret string) string {
	return secret + "123"
}

func main() {
	fmt.Println("Started program..")
	stripped := mod10.Strip("test9oke")
	convertedBytes := mod10.StringToAsciiBytes(stripped)
	fmt.Println(stripped)
	fmt.Println(convertedBytes)
	//todo: Get first block of 10 items
	//todo: Get second block of 10 items
	//if len(ascii) < 10 {
	//	count := 10 - len(ascii)
	//	fmt.Println(count)
	//	var a = make([]byte, count)
	//	for i := 0; i <= count-1; i++ {
	//		//fmt.Printf("appending: %b", i)
	//		a[i] = byte(i)
	//	}
	//	ascii2 := append(ascii, a...)
	//	//fmt.Println(a)
	//	fmt.Println(ascii2)
	//}
	////fmt.Println(len(ascii))
	//for i := 0; i < len(ascii); i++ {
	//}
	//fmt.Printf("%x\n", mod10.HashBlock("payload"))
	//fmt.Println()
}
