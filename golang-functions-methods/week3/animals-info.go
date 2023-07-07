// Write a program which allows the user to get information about a predefined set of animals.
// Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
// The user can issue a request to find out one of three things about an animal:
// 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
// The following table contains the three animals and their associated data which should be hard-coded into your program.
/*
	+--------+------------+-------------------+--------------+
	| Animal | Food eaten | Locomotion method | Spoken sound |
	+--------+------------+-------------------+--------------+
	| cow    | grass      | walk              | moo          |
	| bird   | worms      | fly               | peep         |
	| snake  | mice       | slither           | hsss         |
	+--------+------------+-------------------+--------------+
*/
// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
// Your program should continue in this loop forever.
// Every request from the user must be a single line containing 2 strings.
// The first string is the name of an animal, either “cow”, “bird”, or “snake”.
// The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
// Your program should process each request by printing out the requested data.

// You will need a data structure to hold the information about each
// Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
// Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
// Your program should call the appropriate method when the user makes a request.
// Submit your Go program source code.

package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) InitAnimal(food string, locomotion string, noise string) *Animal {
	animal.food = food
	animal.locomotion = locomotion
	animal.noise = noise

	return animal
}

func (animal *Animal) ToString() {
	fmt.Println(animal.food, animal.locomotion, animal.noise)
}

func (animal *Animal) Print() {
	fmt.Println(animal)
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func initAnimals(animals *map[string]*Animal, animalNames []string) {

	for _, name := range animalNames {
		var newAnimal Animal
		switch name {
		case "cow":
			newAnimal.InitAnimal("grass", "walk", "moo")
		case "bird":
			newAnimal.InitAnimal("worms", "fly", "peep")
		case "snake":
			newAnimal.InitAnimal("mice", "slither", "hsss")

		}
		(*animals)[name] = &newAnimal
	}
}

func main() {
	var animalTypes = []string{"cow", "bird", "snake"}
	var animals map[string]*Animal = make(map[string]*Animal)
	initAnimals(&animals, animalTypes)
	var animalType, info string
	for {

		fmt.Print("Please enter the animal info >")
		fmt.Scan(&animalType, &info)

		index := slices.Index(animalTypes, animalType)

		if index == -1 {
			fmt.Printf("Please enter animal be one of %v\n\n", animalTypes)
			continue
		}

		switch info {
		case "eat":
			animals[animalType].Eat()
		case "move":
			animals[animalType].Move()
		case "speak":
			animals[animalType].Speak()

		}

	}

}
