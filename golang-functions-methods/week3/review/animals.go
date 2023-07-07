package main

import "fmt"

type Animal struct {
	Food             string
	LocomotionMethod string
	Sound            string
}

func (a *Animal) Eat() {
	fmt.Println(a.Food)
}

func (a *Animal) Move() {
	fmt.Println(a.LocomotionMethod)
}

func (a *Animal) Speak() {
	fmt.Println(a.Sound)
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		fmt.Print(">")

		var animalName, action string
		_, _ = fmt.Scan(&animalName, &action)

		animal, found := animals[animalName]
		if !found {
			fmt.Println("valid animals: cow, bird, snake")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("valid actions: eat, move, speak")
		}
	}
}
