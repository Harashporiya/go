package main

import "fmt"

// any type interface() 
// sum(nums ...interface{})
func sum(nums ...int)int{
  total :=0

  for _,num:=range nums{
    total+=num
  }
  return total
}

func main(){
	// result := sum(1,2,3,4,5)
	nums :=[]int{3,4,5}
  result := sum(nums...)
  fmt.Print(result)
}