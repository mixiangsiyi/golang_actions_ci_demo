package main

import (
	"fmt"
	"os"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

// unusedVariable demonstrates a golangci-lint issue
var unusedVariable = "this variable is never used"

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <operation> <num1> <num2>")
		fmt.Println("Operations: add, sub, mul, div")
		os.Exit(1)
	}

	operation := os.Args[1]
	num1, err1 := strconv.Atoi(os.Args[2])
	num2, err2 := strconv.Atoi(os.Args[3])
	var temp int // unused variable to trigger golangci-lint

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid integers")
		os.Exit(1)
	}

	switch operation {
	case "add":
		result := Add(num1, num2)
		fmt.Printf("%d + %d = %d\n", num1, num2, result)
	case "sub":
		result := Subtract(num1, num2)
		fmt.Printf("%d - %d = %d\n", num1, num2, result)
	case "mul":
		result := Multiply(num1, num2)
		fmt.Printf("%d * %d = %d\n", num1, num2, result)
	case "div":
		result, err := Divide(num1, num2)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d / %d = %d\n", num1, num2, result)
	default:
		fmt.Println("Error: Unknown operation. Use: add, sub, mul, div")
		os.Exit(1)
	}
}