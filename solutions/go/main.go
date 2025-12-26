package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Select problem
	fmt.Print("Number of problems Project Euler: ")
	problemInput, _ := reader.ReadString('\n')
	problemInput = strings.TrimSpace(problemInput)
	problemNumber, err := strconv.Atoi(problemInput)

	if err != nil {
		fmt.Println("Error: invalid problem number")
		return
	}

	solver, exists := Solutions[problemNumber]
	if !exists {
		fmt.Println("Problem not found")
		return
	}

	//Write input data
	fmt.Print("Enter number (n): ")
	dataInput, _ := reader.ReadString('\n')
	dataInput = strings.TrimSpace(dataInput)
	n, err := strconv.Atoi(dataInput)

	if err != nil {
		fmt.Println("Error: input data must be a number")
		return
	}

	// Solve the problem
	result := solver(n)
	fmt.Println("Answer:", result)
}