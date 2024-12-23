package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    result := []int{}
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                result = append(result, i)
                result = append(result, j)
            }
        }
    }
	
	// best runtime
	// result := make(map[int]int)
	// for indexNums, valueNums := range nums {
		// 	findValue := target - valueNums

		// 	index, found := result[findValue]
		// 	if found {
		// 		return []int{index, indexNums}
		// 	}
		// 	result[valueNums] = indexNums
	// }

    return result

}

func main() {
	// test case
	nums := []int{2,7,11,15}
	target := 9

	fmt.Println(twoSum(nums, target))
}