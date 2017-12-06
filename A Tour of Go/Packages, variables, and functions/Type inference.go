package main

import (
	"fmt"
)

func main() {
	//Type inference
	v := 42
	fmt.Printf("v is of type %T\n", v)
	l := "hello"
	fmt.Printf("l is of type %T\n", l)
	m := 42.456487312
	fmt.Printf("m is of type %T\n", m)
}
