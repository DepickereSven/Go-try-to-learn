package main

//if you don't use the import it will not work
import (
	"fmt"
)

//can't be changed
const (
	stringmessage       = "Second Hello world!"
	secondStringMessage = "Second message"
)

//just a simple var
var (
	messageForWorld string
	secondMessageForWorld string
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(stringmessage)
	fmt.Println(messageForWorld)
	secondMessageForWorld = "second message"
	fmt.Println(secondStringMessage)
	fmt.Println(secondMessageForWorld)
}


func init() {
	messageForWorld = "New hello World!"
}
