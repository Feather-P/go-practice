package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("misaka")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calling Hello(): " + message)

	names := []string{"murasame", "hoshiro", "misaka"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calling Hellos():", messages)
}
