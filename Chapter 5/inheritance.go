type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Println("Animal is eating.")
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	dog := Dog{
		Animal: Animal{
			Name: "Bobby",
		},
		Breed: "Labrador",
	}
	fmt.Println(dog.Name)
	dog.Eat()
}
