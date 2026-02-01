package main

import (
	"fmt"
)

// slice -> dynamic array (same as vector in c++)
// most used construct in go
// useful methods -> append, len, cap

func main(){
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))


  // make(array, length, capacity)
	// var nums = make([]int, 0, 5)

	// capacity -> maximum number of elements slice can hold before resizing

	// fmt.Println(nums) // -> zero value

	// fmt.Println(cap(nums))

	// nums = append(nums, 10)
	// nums = append(nums, 20)
	// nums = append(nums, 30)
	// nums = append(nums, 40)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))


	// nums := []int{}


	// var nums = []int{1,2,3,4,5}

	// // copy slice

	// var nums2 = make([]int, len(nums))

	// copy(nums2, nums)

	// nums = append(nums, 2)
	// fmt.Println(nums, nums2)





	// slice operator

	// var nums = []int{1, 2, 3}

	// fmt.Println((nums[1:2]))
	// fmt.Println((nums[1:]))
	// fmt.Println((nums[:2]))


	// slice
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1,2}

	// fmt.Println(slices.Equal(nums1, nums2))


	var nums = [][]int{{1,2}, {3,4}}
	fmt.Println(nums)
	
}