package main

import (
	"fmt"
	"log"
	"loyio.me/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	names := []string{"Loyio", "Hex", "Ayesup"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
