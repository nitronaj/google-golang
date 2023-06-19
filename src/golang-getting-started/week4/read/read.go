// Write a program which reads information from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been read from the file,
// your program will have a slice containing one struct for each line in the file.
// After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	var names []Name = make([]Name, 0)

	// get the file name
	fmt.Scan(&fileName)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)
	// Iterate over the lines in the file.
	for scanner.Scan() {
		// Get the current line.
		line := scanner.Text()
		name := strings.Split(line, " ")

		names = append(names, Name{name[0], name[1]})
	}

	f.Close()

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}

}
