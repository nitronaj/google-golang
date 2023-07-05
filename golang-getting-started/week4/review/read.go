package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	type name struct {
		fname string
		lname string
	}

	var FileName string

	fmt.Println("Please enter file name")
	fmt.Scan(&FileName)

	content, _ := ioutil.ReadFile(FileName)
	StringArr := strings.Fields(string(content))

	var StructArr []name

	for i := 0; i < len(StringArr); i += 2 {
		StructArr = append(StructArr, name{StringArr[i], StringArr[i+1]})

	}

	for _, v := range StructArr {
		fmt.Println(v)
	}

}
