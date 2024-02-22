package main

import (
	"fmt"
	"math/rand"
	"time"
)

func use(_ any) interface{} {
	return nil
}

func print_it(text ...any) interface{} {
	fmt.Println(text...)

	return nil
}

func main() {
	var min, max int
	min = 1
	max = 20
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(20) + 1
	var nb int; use(nb);
	for {
		if nb != randomNumber {
			fmt.Print("Enter a number between ", min, " and ", max, " : ")
			fmt.Scanln(&nb)
			if nb == randomNumber {
				print_it("Exact")
			} else if nb > randomNumber {
				print_it("Too big")
			} else {
				print_it("Too little")
			}
		} else {
			break
		}
	}
}
