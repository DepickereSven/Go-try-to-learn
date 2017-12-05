package main

//if you don't use the import it will not work
import (
	"fmt"
)

//can't be changed
const (
	Stringmessage = "Second Hello world!"
)

//just a simple var
var (
	messageForWorld string
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(Stringmessage)
	fmt.Println(messageForWorld)
}


func init() {
	messageForWorld = "New hello World!"
}
