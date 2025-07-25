package package_libary

import "errors"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) (User, error) {
	if name == "" {
		return User{}, errors.New("required name")
	}
	if age <= 0 || age >= 150 {
		return User{}, errors.New("not valid age")
	}
	return User{name: name, age: age}, nil
}
