package main

import "fmt"

func for_example() {
	arr := []int{2, 3, 5, 7, 11, 13}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, a := range arr {
		fmt.Println(i, a)
	}
}
func infinity_cicle() {
	var i int = 0
	for {
		fmt.Println(i)
		i++
		if i == 25 {
			break
		}
	}
}

func main() {
	infinity_cicle()
	for_example()
}
