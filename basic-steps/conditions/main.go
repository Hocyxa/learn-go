package main

import (
	"fmt"
	"time"
)

func ifelse() {
	score := 12
	if score > 15 {
		fmt.Println("Good boy")
	} else if score > 10 {
		fmt.Println("Normal boy")
	} else {
		fmt.Println("Try one more time")
	}
}
func logicalOperator() {
	score := 10
	if score >= 5 && score < 20 {
		fmt.Println("Good boy")
	} else {
		fmt.Println("Try one more time")
	}
}
func switchCase() {
	day := time.Now().Weekday()
	switch day {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Go to work")
	case time.Saturday, time.Sunday:
		fmt.Println("Chill")
	default:
		fmt.Println("WTF")
	}
}
func main() {
	ifelse()
	switchCase()
}
