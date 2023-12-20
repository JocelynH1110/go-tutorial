package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	str := "1988-11-10"
	parsed, err := time.Parse(time.DateOnly, str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", parsed)
}
