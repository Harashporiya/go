package main

import (
	"fmt"
	// "slices"
)

// slice is dynamic array
// most used construct in go
// + useful method

func main(){
	// null in go -> nil
	// uninitialized slice is nil
     //var nums []int
//    fmt.Println(nums==nil)
//   fmt.Print(len(nums))

    // var nums = make([]int, 1, 5)

	// fmt.Print(nums)
    // capacity -> maximum number  of elements can it

	// nums:=[]int{}

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)

	// nums [0]=1
	
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// copy function

	// var nums = make([]int, 0, 5)
	// nums=append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy(nums2,nums)

	// fmt.Println(nums,nums2)

	// Slice Operator

	// var num = []int{1,2,3,4,5,6,7}

	// fmt.Println(num[3:4])
	// fmt.Println(num[:4])
	// fmt.Println(num[5:])

	// slice
	// var num1=[]int{1,2,3}
	// var num2=[]int{1,2,3,4}

	// fmt.Println(slices.Equal(num1,num2))


	// 2d slice
	var nums = [][]int{{1,2},{3,4}}

	fmt.Println(nums)

}