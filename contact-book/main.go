package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var (
		firstName, lastName, email, number, birth string
	)
	fmt.Print("First name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Phone no: ")
	fmt.Scanln(&number)

	fmt.Print("Email: ")
	fmt.Scanln(&email)

	fmt.Print("Date of birth: ")
	fmt.Scanln(&birth)
	str := birth
	_, err := time.Parse(time.DateOnly, str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SAVED!")
}
