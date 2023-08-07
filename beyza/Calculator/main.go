package main

import (
	"fmt"
)
func main(){
	var no1, no2 float64
	var operator string

	fmt.Println("Enter your first number: ")
	fmt.Scan(&no1)

	fmt.Println("Enter your second number: ")
	fmt.Scan(&no2)

	fmt.Println("Select your operation symbols: ")
	fmt.Scan(&operator)

	result := calc(no1, no2, operator)
	fmt.Println("Your result: " ,result)
}
 func calc(no1 float64, no2 float64, operator string) float64{
	var result float64
	switch operator{
	case "+":
		result = no1 + no2
	case "-":
		result = no1 - no2
	case "/":
		result = no1 / no2
	default:
		fmt.Println("Something wrong")
	}
	return result
 }
