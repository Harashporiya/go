package main

import "fmt"

// by value
// func changeNum(num int){
// 	num=5
// 	fmt.Println("In ChangeNum",num)
// }


// by reference

func changeNum(num *int){
	*num = 5 // dereference
	fmt.Println("In ChangeNum",*num)
}

func main(){
  num:=1
  changeNum(&num)
//   fmt.Println("Memory Address", &num)
  fmt.Print("After changeNum in main ",num)
}