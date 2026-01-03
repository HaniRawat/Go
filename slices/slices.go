package main

import (
	"fmt"
	"slices"
)

func main() {
	//slice is a dynamic array like vector in cpp
	//its the most used construct in go
	//provides more useful methods than simple arrays

	//uninitialized slices are nil (nil == null)
	// var nums []int

	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))


	// make([]int, x, y)  x-> size of slice, y->capacity of slice 
	// var nums = make([]int, 2, 3) //the values are 0 by default in "make" method
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//output -> [0, 0], len = 2, cap = 3

	// nums = append(nums, 1)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//output -> [0, 0, 1], len = 3, cap = 3

	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	//output -> [0, 0, 1, 2], len = 4, cap = 6
	//capacity of slice doubles after we exceed the mentioned capacity

	//to create an empty slice which doesnt have any 0s we can uses make([]int, 0, 3)

	// arr := make([]int, 0, 3)
	// arr = append(arr, 1)
	// arr = append(arr, 2)
	// arr = append(arr, 3)

	// fmt.Println(arr)

	//one more method to make slice
	// nums := []int{}
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	//insertion using indices
	// nums := make([]int, 2, 3)
	// nums[0] = 1
	// nums[1] = 2
	// fmt.Println(nums)

	//copy function
	// arr := make([]int, 0, 3)
	// arr = append(arr, 1)
	// nums := make([]int, len(arr))

	// copy(nums, arr)
	// fmt.Println(arr, nums)

	// slice operator 
	// var nums = []int {1,2,3}
	// fmt.Println(nums[0:2])  //2 index is excluded
	// fmt.Println(nums[:2])
	// fmt.Println(nums[:])

	//2d matrix

	// nums := [][]int {{1,2}, {3,4}, {5,6}}

	// for i:=range(2) {
	// 	for j:=range(2) {
	// 		fmt.Println(nums[i][j])
	// 	}
	// }

	//slice package
	var nums1 = []int{1,2}
	var nums2 = []int{1,2}
	fmt.Println(slices.Equal(nums1, nums2)) //to check if two slices are equal or not

	

}