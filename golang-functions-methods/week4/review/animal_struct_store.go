package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}
type Bird struct {
	food       string
	locomotion string
	noise      string
}
type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Println("It eats: ", a.food)
}
func (a Cow) Move() {
	fmt.Println("It moves like: ", a.locomotion)
}
func (a Cow) Speak() {
	fmt.Println("It sounds like : ", a.noise)
}

func (a Bird) Eat() {
	fmt.Println("It eats: ", a.food)
}
func (a Bird) Move() {
	fmt.Println("It moves like: ", a.locomotion)
}
func (a Bird) Speak() {
	fmt.Println("It sounds like : ", a.noise)
}

func (a Snake) Eat() {
	fmt.Println("It eats: ", a.food)
}
func (a Snake) Move() {
	fmt.Println("It moves like: ", a.locomotion)
}
func (a Snake) Speak() {
	fmt.Println("It sounds like : ", a.noise)
}

func main() {

	fmt.Println("Welcome to animal kingdom!")
	m := make(map[string]Animal)
	c1 := Cow{"grass", "walk", "moo"}
	b1 := Bird{"worms", "fly", "peep"}
	s1 := Snake{"mice", "slither", "hsss"}
	m["cow"] = c1
	m["bird"] = b1
	m["snake"] = s1

	for {
		fmt.Println("--Please Enter newanimal/query-- ")
		fmt.Println("1. newanimal <any name> <cow/bird/snake>")
		fmt.Println("2. query <persent name> <eat/move/speak>")
		in := bufio.NewReader(os.Stdin)
		anichosen, _ := in.ReadString('\n')
		process := strings.Split(anichosen, " ")

		animalType := strings.Trim(process[2], " \n")
		if process[0] == "newanimal" {

			switch animalType {
			case "cow":
				m[process[1]] = c1

			case "bird":
				m[process[1]] = b1

			case "snake":
				m[process[1]] = s1
			default:
				fmt.Println("Incorrect command!")
			}

		} else if process[0] == "query" {
			ani, ok := m[process[1]]
			if !ok {
				fmt.Println("Please enter proper command")
			} else {

				switch process[2] {
				case "eat":
					{
						ani.Eat()
					}
				case "move":
					{
						ani.Move()
					}
				case "speak":
					{
						ani.Speak()
					}

				}

			}
		} else {
			fmt.Println("Incorrect command!!")
		}

	}
}
