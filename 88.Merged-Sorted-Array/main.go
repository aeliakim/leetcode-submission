package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	// mergedArray := []int{}
	last := m + n - 1

	p1 := m - 1
	p2 := n - 1

	for i := last; i >= 0; i-- {
		if p1 < 0 {
			nums1[i] = nums2[p2]
			p2--
		} else if p2 < 0 {
			nums1[i] = nums1[p1]
			p1--
		} else {
			if nums1[p1] > nums2[p2] {
				nums1[i] = nums1[p1]
				p1--
			} else {
				nums1[i] = nums2[p2]
				p2--
			}
		}
	}

	fmt.Println(nums1)
// for i := 0; i < m; i++ {
// 	if nums1[i] == 0 {
// 		continue
// 	} else {
// 		mergedArray = append(mergedArray, nums1[i])
// 	}
// }

// fmt.Println(mergedArray)

// if n > 0 {
// 	for j := 0; j < n; j++ {
// 		if nums2[j] == 0 {
// 			continue
// 		} else {
// 			mergedArray = append(mergedArray, nums2[j])
// 		}
// 	}
// }

// fmt.Println(mergedArray)

// for k := range mergedArray {
// 	for p := k + 1; p < len(mergedArray); p++ {
// 		if mergedArray[k] > mergedArray[p] {
// 			temp := mergedArray[k]
// 			mergedArray[k] = mergedArray[p]
// 			mergedArray[p] = temp
// 		}
// 	}
// }

// fmt.Println(mergedArray)
// nums1 = mergedArray[0:]
// fmt.Println(nums1)
}

func main() {

	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3

	merge(nums1, m, nums2, n)
}
