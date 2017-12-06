package main

import (
	"fmt"
)

func longWayForIterating()  {
	sum := 0
	for i := 0; i < 10; i++{
		sum += 1
		fmt.Println(sum)
	}
	fmt.Println("finally done:", sum)
}

func shortWayForIterating()  {
	sum := 1
	for ; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("finally done:", sum)
}

func shortWayForIteratingByOne()  {
	sum := 1
	for ; sum < 1000; {
		sum ++
		fmt.Println(sum)
	}
	fmt.Println("finally done:", sum)
}

func main() {
	//simple for loop
	longWayForIterating()

	//short way for Iterating
	shortWayForIterating()
	shortWayForIteratingByOne()
}
