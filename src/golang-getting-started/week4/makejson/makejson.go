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
	var user2 map[string]string = make(map[string]string)

	fmt.Scan(&name, &address)

	user["name"] = name
	user["address"] = address

	userJson, _ := json.Marshal(user)
	fmt.Println(userJson)

	// json.Unmarshal(userJson, &user2)
	// fmt.Println(user2)

}
