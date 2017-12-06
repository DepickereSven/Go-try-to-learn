package main

import (
	"fmt"
)

func main() {
	//For is Go's "while"
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("finally done:", sum)

}
