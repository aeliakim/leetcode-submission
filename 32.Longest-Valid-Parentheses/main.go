package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func longestValidParentheses(s string) int {
	result := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
				length := i - stack[len(stack)-1]
				if length > result {
					result = length
				}
			} else {
				stack[0] = i
			}
		}
	}
	return result
}

func main() {
	fmt.Printf("Input the parentheses: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", input)
	fmt.Print(longestValidParentheses(input))

}
