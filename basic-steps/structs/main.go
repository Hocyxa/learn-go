package main

import (
	"fmt"
)

type User struct {
	// Возраст пользователя
	// Правило: [0 - 150]
	Age int

	// Имя пользователя
	// Правило: Не пустое
	Name string

	// Почта пользователя
	// Правило:
	Email string

	// Рэйинг пользователя
	// Правило:  [0.0 - 10.0]
	Ratting float32
	// Забанен ли пользоватлеь пользователя
	// Правило:  [0.0 - 10.0]
	Ban bool
}

func New(Age int,
	Name string,
	Email string,
	Ratting float32,
	Ban bool,
) User {
	if Name == "" || Age <= 0 || Age >= 150 {
		return User{}
	}
	return User{
		Name:    Name,
		Age:     Age,
		Email:   Email,
		Ratting: Ratting,
		Ban:     Ban,
	}
}
func (u *User) Greeting() {
	fmt.Println()
	fmt.Println("Age: ", u.Age)
	fmt.Println("Name: ", u.Name)
	fmt.Println("Email: ", u.Email)
	fmt.Println("Rating: ", u.Ratting)
	fmt.Println("Ban: ", u.Ban)
	fmt.Println()
}

func (u *User) RattingUP(ratting float32) {
	if u.Ratting+ratting <= 10 {
		u.Ratting += ratting
	}
}

func main() {
	user := New(23,
		"Peter",
		"string@mail.ru",
		1.2,
		false)
	user.Greeting()

	user.RattingUP(2)

}
