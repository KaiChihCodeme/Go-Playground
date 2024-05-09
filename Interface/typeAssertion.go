package main

import (
	"fmt"
)

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name + "said bark!")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name + " said meow~")
}

func main() {
	var cat Animal
	cat = Cat{
		Name: "kitty",
	}

	cat.Speak()

	kitty_cat := cat.(Cat)
	kitty_cat.Speak()

	// will not panic, due to if we do not handle it will cause panic
	if nocat, ok := cat.(Dog); ok {
		nocat.Speak()
	}

	dog := Dog {
		Name: "Lucky",
	}

	PrintSpeak(dog)
}

func PrintSpeak(animal Animal) {
	// type to go through case
	switch animal.(type){
	case Cat:
		animal.Speak()
	case Dog:
		animal.Speak()
	}
}