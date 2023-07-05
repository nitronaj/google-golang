package main

import "fmt"
import "os"
import "bufio"
import "strings"

type fullName struct{
	fname, lname string
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the name of the text file.")
	scanner.Scan()
	fileName := scanner.Text()
	var f *os.File
	nameArr := make([]fullName, 0, 1)
	for{
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			fmt.Println("Please Enter valid file name:")
			scanner.Scan()
			fileName = scanner.Text()
		}else{
			f = fi
			break
		}
	}
	fScanner := bufio.NewScanner(f)
	var arr []string
	for fScanner.Scan() {
		arr = strings.Split(fScanner.Text(), " ")
		nameArr = append(nameArr, fullName{arr[0], arr[1]})
	}
	f.Close()
	for _, name := range nameArr {
		fmt.Printf("%v - %v\n", name.fname, name.lname)
	}
}
