package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 1231211123123111111

	fmt.Printf("%T\n", i)
	fmt.Printf("Type %T size of %d byte\n", i, unsafe.Sizeof(i))
}
