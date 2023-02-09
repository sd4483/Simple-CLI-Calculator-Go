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
		fmt.Println("Choose what operation you would like to perform: \n 1 -> Add \n 2 -> Subtract \n 3 -> Multiply \n 4 -> Divide")
		fmt.Scanln(&operation)

		switch operation {
		case "1":
			result = num_1 + num_2
		case "2":
			result = num_1 - num_2
		case "3":
			result = num_1 * num_2
		case "4":
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
