package main

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
)

// var declaration of booleans
var c, python, java bool

//var declaration of int's
var ik, jk int = 1, 2

//declaration of return value
func add(x int, y int) int {
	return x + y
}

//shorter declaration of setup arguments
func addShort(x, y int) int {
	return x + y;
}

func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	//format newline \n
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	//math.Pi needs capital letter because it's an export from math
	fmt.Println(math.Pi)
	//call func add
	fmt.Println(add(42, 13))
	//call func addshort
	fmt.Println(addShort(42, 13))
	// get 2 String values back
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	//"naked" return
	fmt.Println(split(17))
	//var
	var i int
	fmt.Println(i, c, python, java)
	// var declaration of values
	var c, python, java = true, false, "no!"
	fmt.Println(ik, jk, c, python, java)
	//shorter declaration of vars
	var id, jd int = 1, 2
	k := 3
	cd, pythond, javad := true, false, "no!"
	fmt.Println(id, jd, k, cd, pythond, javad)
	//big numbers
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	//Zero values
	var im int
	var fm float64
	var bm bool
	var sm string
	fmt.Printf("%v %v %v %q\n", im, fm, bm, sm)
}
