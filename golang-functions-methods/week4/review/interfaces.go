package main

import (
	"fmt"
)

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a animal) Eat() {
	fmt.Println(a.food)
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
}

func (a animal) Speak() {
	fmt.Println(a.noise)
}

func animalCreatedMessage() {
	fmt.Println("Created it!")
}

func main() {
	fmt.Println("Enter a request in the form of 'newanimal <name> <type>' or 'query <name> <info>'")
	fmt.Println("Where <name> is the name of the animal and <type> is the type of animal. <info> could be eat,move,speak. Press x to exit.")

	var request string
	var name string
	var info string

	var animals = make(map[string]animalInterface)

	for {
		if len(animals) > 0 {
			fmt.Print("Current animals: (")
			for key := range animals {
				fmt.Print(key + ", ")
			}
			fmt.Println(")")
		}

		fmt.Print("> ")
		fmt.Scan(&request)
		if request == "x" {
			break
		}

		fmt.Scan(&name, &info)

		if request == "newanimal" {
			if info == "cow" {
				animals[name] = animal{"grass", "walk", "moo"}
				animalCreatedMessage()
			} else if info == "bird" {
				animals[name] = animal{"worms", "fly", "peep"}
				animalCreatedMessage()
			} else if info == "snake" {
				animals[name] = animal{"mice", "slither", "hsss"}
				animalCreatedMessage()
			} else {
				fmt.Println("Invalid animal type")
			}
		} else if request == "query" {
			if info == "eat" {
				animals[name].Eat()
			} else if info == "move" {
				animals[name].Move()
			} else if info == "speak" {
				animals[name].Speak()
			} else {
				fmt.Println("Invalid request")
			}
		} else {
			fmt.Println("Invalid request")
		}
	}

}
