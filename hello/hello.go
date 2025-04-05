package main

import (
	"fmt"
	"log"
    "strings"
	"example.com/greetings"
)

func main() {
	log.SetPrefix(("greetings123: "))
	log.SetFlags(0)
	names := []string{"Vlad", "Kesha", "Yana"}

	message, err := greetings.Hellos(names)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println("asdsad")
	fmt.Println(strings.Fields("asd asd 12"))
	for _,value:= range strings.Fields("asd asd 12") {
		fmt.Println(value)
	}
}
