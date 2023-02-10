package main

import "fmt"

func CLI_Calc() string {

	var num_1 float64
	var num_2 float64
	var operation string

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num_1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&num_2)

	var result float64 = 0

	for result == 0 {
		fmt.Println("Choose what operation you would like to perform:" +
			"\n + -> Add \n - -> Subtract \n * -> Multiply \n / -> Divide")
		fmt.Scanln(&operation)

		switch operation {
		case "+":
			result = num_1 + num_2
		case "-":
			result = num_1 - num_2
		case "*":
			result = num_1 * num_2
		case "/":
			result = num_1 / num_2
		default:
			fmt.Println("--- Wrong input. Try again or click ctrl+c to exit ---")
			result = 0
		}
		if result != 0 {
			break
		}
	}

	return fmt.Sprintf("The result is: %v", result)
}
