package main

import (
	"fmt"
)

func main() {
	fmt.Println("kokowaa")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d - %b %x \n", i, i, i)
	}

}
