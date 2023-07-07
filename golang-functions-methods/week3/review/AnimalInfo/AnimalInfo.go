package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) Eat () {
	fmt.Println(a.food)
}

func (a *Animal) Move () {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak () {
	fmt.Println(a.noise)
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animal_name, action string

	fmt.Println("Enter animal name(cow/bird/snake) and action(eat/move/speak)")

	for {
		fmt.Print("\n>")
		fmt.Scanf("%s", &animal_name);
		fmt.Scanf("%s", &action);

		switch (animal_name){
			case "cow" : 
				if action == "eat" {
					cow.Eat()
				} else if action == "move" {
					cow.Move()
				} else if action == "speak" {
					cow.Speak()
				} else {
					fmt.Println("Invalid Input")
				}
			
			case "bird" :
				if action == "eat" {
					bird.Eat()
				} else if action == "move" {
					bird.Move()
				} else if action == "speak" {
					bird.Speak()
				} else {
					fmt.Println("Invalid Input")
				}
			
			case "snake" :
				if action == "eat" {
					snake.Eat()
				} else if action == "move" {
					snake.Move()
				} else if action == "speak" {
					snake.Speak()
				} else {
					fmt.Println("Invalid Input")
				}
		}
	}

}

/*
SAMPLE RUN

Enter animal name(cow/bird/snake) and action(eat/move/speak)

>bird eat
worms

>snake speak
hsss

>cow move
walk

>
*/