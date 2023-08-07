package main

import (
	"fmt"
)
func main(){
	var number1, number2 float64
	var operator string

	fmt.Println("Enter your first number: ")
	fmt.Scan(&number1)

	fmt.Println("Enter your second number: ")
	fmt.Scan(&number2)
	
	fmt.Println("Select your operator symbols: ")
	fmt.Scan(&operator)

	result := calculator(number1, number2, operator)
	fmt.Println("Your result: ", result)
}

func calculator(number1 float64, number2 float64, operator string) float64{
	var result float64
	switch operator{
	case "+":
		result= number1 + number2
	case "-":
		result= number1 - number2
	case "*":
		result= number1 * number2
	case "/":
		result= number1 / number2
	default:
		fmt.Println("Something went wrong!")
	}
	return result
}

