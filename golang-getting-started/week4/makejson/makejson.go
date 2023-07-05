// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	var user map[string]string = make(map[string]string)

	fmt.Println("Enter a name:")
	fmt.Scan(&name)
	fmt.Println("Enter an address:")
	fmt.Scan(&address)

	user["name"] = name
	user["address"] = address

	userJson, _ := json.MarshalIndent(user, "", " ")

	fmt.Println(string(userJson))
}
