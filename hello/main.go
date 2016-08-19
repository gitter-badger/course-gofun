package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
)

func main() {
	fmt.Print(c.Clear)
	println(c.O + "hello debug world!" + c.X)
	print(c.O + "hello again debug world!\n" + c.X)
	fmt.Println(c.Multi("Hello world!"))
	fmt.Print(c.B + "Hello again world!\n" + c.X)
}
