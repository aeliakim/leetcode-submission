package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calPoints(operations []string) int {
	// var res int
	// result := 0
	// var res int
	// record := []int{}
	// for index := 0; index < len(operations); index++ {
	// 	numberRegex := `^-?\d+(\.\d+)?$`

	// 	matched, _ := regexp.MatchString(numberRegex, operations[index])

	// 	if matched {
	// 		var num, err = strconv.Atoi(operations[index])
	// 		if err == nil {
	// 			record = append(record, num)
	// 		}
	// 	} else {
	// 		if operations[index] == "C" {
	// 			var lenRec = len(record) - 1
	// 			record = record[:lenRec]
	// 		} else if operations[index] == "D" {
	// 			var lenRec = len(record) - 1
	// 			var num = record[lenRec] * 2
	// 			record = append(record, num)
	// 		} else if operations[index] == "+" {
	// 			var lenRec = len(record) - 1
	// 			res = record[lenRec] + record[lenRec-1]
	// 			record = append(record, res)
	// 		}
	// 	}
	// }

	result := 0
	record := make([]int, 0, len(operations))

	for _, op := range operations {
		switch op {
		case "C":
			result -= record[len(record)-1]
			record = record[0 : len(record)-1]
		case "D":
			record = append(record, record[len(record)-1]*2)
			result += record[len(record)-1]
		case "+":
			record = append(record, record[len(record)-1]+record[len(record)-2])
			result += record[len(record)-1]
		default:
			val, _ := strconv.Atoi(op)
			record = append(record, val)
			result += record[len(record)-1]
		}
	}
	for i := range record {
		result += record[i]
	}
	return result
}

func main() {
	ops := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Input the record: ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		ops = append(ops, input)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(calPoints(ops))

}
