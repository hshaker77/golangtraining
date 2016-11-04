package main

import (
	"fmt"

	"github.com/hshaker77/golangtraining/test/lib"
)

func main() {
	fmt.Println("kokowaa")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d - %b %x \n", i, i, i)
	}
	lib.Hany()

}
