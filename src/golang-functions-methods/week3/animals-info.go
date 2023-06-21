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

// You will need a data structure to hold the information about each animal.
// Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
// Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
// Your program should call the appropriate method when the user makes a request.
// Submit your Go program source code.

package main

import (
	"fmt"
	"google-golang/src/golang-functions-methods/week3/animal"
)

func initAnimals(animalNames []string) map[string]*animal.Animal {
	var animals map[string]*animal.Animal = make(map[string]*animal.Animal)
	for _, name := range animalNames {
		var newAnimal animal.Animal
		switch name {
		case "cow":
			newAnimal.InitAnimal("grass", "walk", "moo")
		case "bird":
			newAnimal.InitAnimal("worms", "fly", "peep")
		case "snake":
			newAnimal.InitAnimal("mice", "slither", "hsss")

		}
		animals[name] = &newAnimal
	}

	return animals

}

func initAnimals2(animalNames []string) map[string]*animal.Animal {
	var cow animal.Animal
	cow.InitAnimal("grass", "walk", "moo")

	var bird animal.Animal
	bird.InitAnimal("worms", "fly", "peep")

	var snake animal.Animal
	snake.InitAnimal("mice", "slither", "hsss")

	animals := map[string]*animal.Animal{
		"cow":   &cow,
		"snake": &bird,
		"bird":  &snake,
	}

	return animals
}

func main() {

	animals := initAnimals([]string{"cow", "bird", "snake"})
	var name, info string
	for {

		fmt.Print(">")
		fmt.Scan(&name, &info)

		switch info {
		case "eat":
			animals[name].Eat()
		case "move":
			animals[name].Move()
		case "speak":
			animals[name].Speak()
		}

	}

}
