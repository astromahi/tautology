/*
Purpose 	  : Checking wheather the given expression is a tautology
File Name	  : verifier.go
Package		  : main
Date 		  : 03.07.2016
Author 		  : Mahendran Kathirvel
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"tautology/postfix"
)

func main() {

	var input string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Expression: ")
	if scanner.Scan() {
		input = scanner.Text()
	}

	if strings.TrimSpace(input) == "" {
		fmt.Print("***Please input the expression***\n")
		os.Exit(1)
	}

	start := time.Now()

	fmt.Print("\n-----------Tautology Verifier-----------\n")

	// Convert given expression into postfix
	exp, expvar := postfix.Convert(input)

	var data []rune
	len := len(expvar)

	var result []bool

	// constructing truth table for validating postfix expression
	for i := 0; i < (1 << uint(len)); i++ {
		data = nil
		for j := len - 1; j >= 0; j-- {
			if (i & (1 << uint(j))) != 0 {
				data = append(data, '1')
			} else {
				data = append(data, '0')
			}
		}

		// evaluates the expression for generated data
		result = append(result, postfix.Evaluate(exp, expvar, data))
	}

	var status bool
	// validating the results
	for i := 0; i < (1 << uint(len)); i++ {
		if !result[i] {
			status = false
			break
		} else {
			status = true
		}
	}

	elapsed := time.Now().Sub(start)

	// printing the result
	if status {
		fmt.Println("Given expression is a tautology")
	} else {
		fmt.Println("Given expression is not a tautology")
	}
	fmt.Printf("Time taken: %s\n", elapsed.String())
}
