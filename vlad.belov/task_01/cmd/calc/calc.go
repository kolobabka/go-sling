package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	SUCCESS = iota
	WRONG_NUM_INPUT
	WRONG_OP_INPUT
	ZERO_DIV
	UNKNOWN_ERROR
)

const (
	SUM string = "+"
	SUB string = "-"
	MUL string = "*"
	DIV string = "/"
)

func checkOperation(operation string) {
	switch operation {
	default:
		fmt.Fprintln(os.Stderr, "Wrong operation type")
		os.Exit(WRONG_OP_INPUT)
	case SUM:
	case SUB:
	case MUL:
	case DIV:
	}
}

func parseArgs() (float64, string, float64) {
	stdin := bufio.NewReader(os.Stdin)
	var num_1 float64
	var num_2 float64
	var operation string

	fmt.Println("Enter the first number:")
	_, err := fmt.Fscan(stdin, &num_1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Wrong number input")
		os.Exit(WRONG_NUM_INPUT)
	}

	fmt.Println("Choose the operation (+, -, *, /):")
	fmt.Fscan(stdin, &operation)

	checkOperation(operation)

	fmt.Println("Enter the second number:")
	_, err = fmt.Fscan(stdin, &num_2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Wrong number input")
		os.Exit(WRONG_NUM_INPUT)
	}

	return num_1, operation, num_2
}

func calcResult(num_1 float64, operation string, num_2 float64) float64 {
	var result float64

	switch operation {
	case "+":
		result = num_1 + num_2
	case "-":
		result = num_1 - num_2
	case "*":
		result = num_1 * num_2
	case "/":
		if num_2 == 0 {
			fmt.Fprintln(os.Stderr, "Division by zero")
			os.Exit(ZERO_DIV)
		}
		result = num_1 / num_2
	default:
		fmt.Fprintln(os.Stderr, "Unknown error")
		os.Exit(UNKNOWN_ERROR)
	}

	return result
}

func main() {

	num_1, operation, num_2 := parseArgs()

	res := calcResult(num_1, operation, num_2)
	fmt.Println("Result :", res)
}
