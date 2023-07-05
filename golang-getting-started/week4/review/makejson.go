package main
import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

type user struct {
	Name	string	`json:"name"`
	Addr	string	`json:"address"`
}

func main() {
	var m map[int]user
	m = make(map[int]user)
	counter := 0

	fmt.Printf("Hello, User!\nThis is a program for make JSON Objects from: Omar Ramirez\n\n")

	for {
		fmt.Printf("\nMenu: \nDo you want add someone? (yes/no)\nEnter your answer: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text()
		if answer == "no" || answer == "NO" {
			break
		}
		if answer == "yes" || answer == "YES" {
			fmt.Printf("Enter a name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Printf("Enter your address: ")
			scanner.Scan()
			address := scanner.Text()

			var user1 user
			user1.Name = name
			user1.Addr = address


			m[counter] = user1
			counter = counter+1

			barr, err := json.Marshal(m)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Print("\n")
			fmt.Println(string(barr))

		}
	}
}