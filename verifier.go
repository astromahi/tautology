/*
Purpose 	  : Checking wheather the given expression is a tautology
File Name	  : verifier.go
Package		  : main
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
	fmt.Println("Postfix is: ", string(exp))

	var data []rune
	len := len(expvar)

	result := make(chan bool)

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
		// deploy goroutine for each data set
		go func() {
			result <- postfix.Evaluate(exp, expvar, data)
		}()
	}

	var status bool

	// validating the results received from channel
	for i := 0; i < (1 << uint(len)); i++ {
		if <-result {
			status = true
		} else {
			status = false
			break
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
