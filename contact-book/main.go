package main

import "fmt"

func main() {
	var (
		firstName, lastName, email string
		number, birth              int
	)
	fmt.Print("First name: ")
	fmt.Scanf("%s", &firstName)

	fmt.Print("Last name: ")
	fmt.Scanf("%s", &lastName)

	fmt.Print("Phone no: ")
	fmt.Scanf("%d", &number)

	fmt.Print("Email: ")
	fmt.Scanf("%s", &email)

	fmt.Print("Date of birth: ")
	fmt.Scanf("%s", &birth)
}
