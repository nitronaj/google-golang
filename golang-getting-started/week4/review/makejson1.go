package main
import "fmt"
import "encoding/json"
func main() {
	var n string
	var add string
	fmt.Printf("Enter your name : ")
	fmt.Scan(&n)
	fmt.Printf("Enter your address : ")
	fmt.Scan(&add)
	m:=make(map[string]string)
	m["name"]=n
	m["address"]=add
	btar,_:=json.Marshal(m)
	fmt.Println(string(btar))
}