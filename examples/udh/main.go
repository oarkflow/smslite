package main

import (
	"fmt"

	"github.com/oarkflow/smslite"
)

func main() {
	u8 := smslite.BuildConcatUDH8(0xAB, 2, 1)
	u16 := smslite.BuildConcatUDH16(0xBEEF, 3, 2)
	fmt.Println("8-bit:", u8.Hex)
	fmt.Println("16-bit:", u16.Hex)
	parsed, err := smslite.ParseUDH([]byte{0x05, 0x00, 0x03, 0xAB, 0x02, 0x01})
	fmt.Printf("parsed=%+v err=%v\n", parsed, err)
}
