package main

import "fmt"


//  func add(a int,b int) int{
//    return a+b
//  }

// func add(a, b int) int{
// 	return a+b
//   }

//   func getLanguages()(string, string,bool){
// 	return "golang", "javascript", true
//   }

//   func processIt(fn func(a int)int){
// 	fn(1)
//   }

func processIt() func(a int)int{
	return func(a int)int{
		return 4
	}
  }

 func main(){
//    result := add(2,3)
//    fmt.Print(result)
//    fmt.Println(getLanguages())
     
//    lang1, lang2, lang3 := getLanguages()
//    fmt.Println(lang1,lang2,lang3)
    //  fn:=func(a int) int{
	// 	return 2
	//  }
   
    fn:=processIt()
	fmt.Print(fn(6))
 
 }