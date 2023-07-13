package main
import "fmt"
type person struct {
name string
age int
}
type student struct {
grade int
person
}
func main() {
var $1 = student{}
51.name = "wick"
$1.age = 21
sl.grade = 2
fmt.Print1n("name z", 51.name)
fmt.Print1n("age :", 51.age)
fmt.Print1n("age z", 51.person.age)
fmt.Println("grade z", 51.grade)
}