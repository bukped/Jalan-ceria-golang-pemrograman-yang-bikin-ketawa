package main

import "fmt"

type user struct {
	name string
}

func (u user) getName() string {
	return u.name
}
func newUser(name string) user {
	return user{name: name}
}
func main() {
	user := newUser("Rakasona")
	fmt.Println(user.getName()) // Rakasona
}
