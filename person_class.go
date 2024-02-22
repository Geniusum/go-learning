package main

import (
	"fmt"
)

func use(_ any) interface{} {
	return nil
}

func print_it(text ...any) interface{} {
	fmt.Println(text...)

	return nil
}

type Person struct {
	name string
	age int
}

func (self Person) say(text string) {
	fmt.Println(self.name + " > " + text)
}

func main() {
	person := Person{name: "John", age: 30}
	person.say("Hello !")
}
