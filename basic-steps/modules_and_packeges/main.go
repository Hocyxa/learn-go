package main

import (
	"fmt"

	lib "github.com/Hocyxa/learn-go/basic-steps/modules_and_packeges/package_libary"
	"github.com/k0kubun/pp"
)

func main() {
	u, err := lib.NewUser("Pavel", 23)

	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	pp.Println(u)
	fmt.Println(lib.Some_const)
	lib.SayHello()
}
