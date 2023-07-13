package main
import "fmt"
type person struct (
name string
age int
}
type student struct {
person
age int
grade in:
}
func main“ {
var 51 = student{}
51.name = "wick"
51.1192 = 21 // age 01 student
$1.person.age = 22 // age of person
fmt.Print1n(51.name)
fmt.Println(sl.age)
fut.Print1n(51.person.age)
}
