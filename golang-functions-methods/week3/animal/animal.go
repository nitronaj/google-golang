package animal

import (
	"fmt"
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
