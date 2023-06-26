package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

/****** COW ******/
type Cow string

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

/****** Bird ******/
type Bird string

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

/****** Snake ******/
type Snake string

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func getCommand() {

}

func addNewAnimal(animalName string, animalType string) *Animal {
	var newAnimal Animal
	switch animalType {
	case "cow":
		cow := Cow(animalName)
		newAnimal = cow

	case "bird":
		newAnimal = Bird(animalName)

	case "snake":
		newAnimal = Snake(animalName)

	}
	fmt.Println(newAnimal)
	return &newAnimal
}

func playAction(animal Animal, action string) {
	switch action {
	case "food":
		animal.Eat()
	case "Locomtion":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func queryAnimal(animals map[string]*Animal, name string, action string) {
	for animalName := range animals {
		if animalName == name {
			playAction(*animals[name], action)
		}
	}
}

func main() {
	var animals map[string]*Animal = make(map[string]*Animal)

	/*
		|	command		|		arg1		| 		arg2		|
		|	newanimal	|	animal name		|	animal type		|
		|	query		|	animal name		|	animal info		|
		e.g: newanimal Milky cow
			 query Milky food
	*/

	var command, arg1, arg2 string = "newanimal", "milky", "cow"
	for {
		fmt.Print(">")
		fmt.Scan(&command, &arg1, &arg2)

		switch command {
		case "newanimal":
			animals[arg1] = addNewAnimal(arg1, arg2)
			fmt.Println("Created it!")
		case "query":
			queryAnimal(animals, arg1, arg2)
		}
	}

}
