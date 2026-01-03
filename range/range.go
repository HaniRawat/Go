package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sum := 0

	// for i := range len(nums) {
	// 	fmt.Println(nums[i], i)
	// 	sum += nums[i]
	// }
	// fmt.Println(sum)

	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	// m := map[string]string {
	// 	"fname" : "john",
	// 	"lname" : "doe",
	// } 

	// for k, v := range m {
	// 	fmt.Println(k)
	// 	fmt.Println(v)
	// }

	//unicode code point rune
	// i->starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}