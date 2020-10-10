package main

import (
	"fmt"
	"github.com/Hilux-engine/golang-sdk/lib"
)

func main() {

	var foo string = "foo"
	var seed uint32 = 0x12345678

	var key = []byte(foo)

	hashNum := lib.MurmurHash3_x64_128(key, uint64(seed))

	fmt.Printf("%d\n", hashNum)

}
