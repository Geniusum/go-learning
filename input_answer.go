package main

import "fmt"

func use(_ any) interface{} {
	return nil
}

func print_it(text ...any) interface{} {
	fmt.Println(text ...)

	return nil
}

func main() {
	fmt.Print("Veuillez saisir du texte : ")
	var input string; use(input)
	fmt.Scanln(&input)
    print_it("Vous avez saisie :", input)
}
