package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

/****** COW ******/
type Cow struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (cow Cow) String() string {
	return "Cow"
}

/****** Bird ******/
type Bird struct{}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (bird Bird) String() string {
	return "Bird"
}

/****** Snake ******/
type Snake struct{}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func (snake Snake) String() string {
	return "Snake"
}

// this part is mainly for trying to work with ENUM
type CommandType int

const (
	NEW_ANIMAL CommandType = iota
	QUERY
	PRINT
)

var commandTypes = map[string]CommandType{
	"newanimal": NEW_ANIMAL,
	"query":     QUERY,
	"print":     PRINT,
}

type AnimalType int

const (
	COW AnimalType = iota
	BIRD
	SNAKE
)

var animalTypes = map[string]AnimalType{
	"cow":   COW,
	"bird":  BIRD,
	"snake": SNAKE,
}

type AnimalAction int

const (
	EAT AnimalAction = iota
	MOVE
	SPEAK
)

var animalActions = map[string]AnimalAction{
	"eat":   EAT,
	"move":  MOVE,
	"speak": SPEAK,
}

func verifyCommand(command string, arg1 string, arg2 string) bool {
	commandType, ok := commandTypes[strings.ToLower(command)]

	if !ok {
		fmt.Println("Command is not valid!!!")
		return false
	}

	if commandType == NEW_ANIMAL {
		_, ok = animalTypes[strings.ToLower(arg2)]
		if !ok {
			fmt.Println("The animal type is not valid!!!")
			return false
		}
	}

	if commandType == QUERY {
		_, ok = animalActions[strings.ToLower(arg2)]
		if !ok {
			fmt.Println("The animal action is not valid!!!")
			return false
		}
	}

	return true
}

func addNewAnimal(animalName string, animalType AnimalType) Animal {
	var newAnimal Animal
	switch animalType {
	case COW:
		newAnimal = Cow{}
	case BIRD:
		newAnimal = Bird{}
	case SNAKE:
		newAnimal = Snake{}
	}
	return newAnimal
}

func playAction(animal Animal, action AnimalAction) {
	switch action {
	case EAT:
		animal.Eat()
	case MOVE:
		animal.Move()
	case SPEAK:
		animal.Speak()
	}
}

func queryAnimal(animals *map[string]Animal, name string, action AnimalAction) {
	// for animalName := range *animals {
	// 	if animalName == name {
	// 		playAction((*animals)[name], action)
	// 	}
	// }

	_, ok := (*animals)[name]
	if !ok {
		fmt.Println("The animal does not exist")
	} else {
		playAction((*animals)[name], action)
	}

}

func displayAnimalList(animals *map[string]Animal) {
	for name, animal := range *animals {
		fmt.Printf("%s %v\n", name, animal)
	}

	fmt.Println()
}

func main() {
	var animals map[string]Animal = make(map[string]Animal)

	/*
		|	command		|		arg1		| 		arg2		|
		|	newanimal	|	animal name		|	animal type		|
		|	query		|	animal name		|	animal info		|
		e.g: newanimal Milky cow
			 query Milky eat
	*/

	for {
		var command, arg1, arg2 string
		fmt.Println("add new animal: newanimal Milky cow|bird|snake")
		fmt.Println("query an animal: query Milky eat|move|speak")

		fmt.Print("\n>")
		fmt.Scan(&command)

		if command == "print" {
			displayAnimalList(&animals)
			continue
		}

		fmt.Scan(&arg1, &arg2)

		if !verifyCommand(command, arg1, arg2) {
			continue
		}

		switch command {
		case "newanimal":
			animals[arg1] = addNewAnimal(arg1, animalTypes[strings.ToLower(arg2)])
			fmt.Println("Created it!")
		case "query":
			queryAnimal(&animals, arg1, animalActions[strings.ToLower(arg2)])
		}

		fmt.Println()

		// Clear the buffer.
		// os.Stdin.Read(make([]byte, 10))
	}

}
