package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	message := greetings.Hello("Justin")

	fmt.Println(quote.Go())

	fmt.Println(message)
}
