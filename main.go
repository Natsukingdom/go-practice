package main

import (
	"./fizzbuzz"
	"fmt"
)

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func main() {
	results := fizzbuzz.FizzBuzzResult(30)
	for _, result := range results {
		fmt.Println(result)
	}
}
