package main

import "fmt"

func variables() {
	//int
	int_index := 0
	var num int
	//double(float)
	double_number := 0.5
	var numf float64
	//string
	string_game_over := "hello"
	var s string
	//char
	c := 'a'
	var literal rune
	//bool
	flag := true

	var subscribe bool
	fmt.Printf("%d %f %s %c ", int_index, double_number, string_game_over, c)
	fmt.Println(flag)
	fmt.Println(num, numf, s, literal, subscribe)
}
func array() {
	arr := []int{1, 2, 3, 45}
	fmt.Println(arr)
}

func main() {
	array()
	variables()
	const t = "qwewe"
	fmt.Println(t)
}
