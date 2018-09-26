package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name, Role, Email string
	Age               int
}

func (u User) Salary() (int, error) {
	switch u.Role {
	case "Developer":
		return 100, nil
	case "Architect":
		return 200, nil
	}
	return 0, errors.New("I'm not able to handle this role.")
}

func main() {
	users := User{Name: "Mark", Role: "Ninja", Age: 36}
	fmt.Println(users.Salary())
}
