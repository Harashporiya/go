package main

import "fmt"

// iterating over data structures
func main() {
//  nums := []int{6,7,8}
 
//  for i:=0; i<len(num);i++{
// 	fmt.Print(num[i]," ")
//  }
//  sum:=0
//   for i, num := range nums {
	// sum+=num
// 	fmt.Println(num," ",i)
//   }
//   fmt.Print(sum)

//   m := map[string]string{"fname":"harash","lname":"Poriya Jaat"}

//   for k, v:=range m {
// 	fmt.Println(k," ", v)
//   }

// for k:=range m {
// 	fmt.Println(k," ")
//   }


  // unicode code point rune
  // starting byte of rune
  // 255 -> 1 byte , 2 byte 
   for i, c := range "Harash Poriya"{
	fmt.Println(i,string(c))
	fmt.Println(i,c)
   }

  
	
}