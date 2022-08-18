package main

import (
	"fmt"
	"strconv"
)

func task8(toChange *int64, whichBit int) {
	fmt.Printf("Было: \n%s\n", strconv.FormatUint(uint64(*toChange), 2))
	*toChange ^= 1 << (whichBit - 1)
	fmt.Printf("Стало: \n%s\n", strconv.FormatUint(uint64(*toChange), 2))
}
