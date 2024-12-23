package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fizzBuzz(n int) []string {
	var result = []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func main() {
	fmt.Printf("Input the number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	// fmt.Printf("%d", input)
	fmt.Print(fizzBuzz(input))
}
