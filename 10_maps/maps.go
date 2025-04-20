package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main(){
  // creating map

//   m := make(map[string]string)

  // setting an elements
//   m["name"]="golang"
//   m["Harash"]="Jaat"

  // get an element
//   fmt.Println(m["Harash"])
  // iMP: if key does not exists in the map thenit return zero value
//   fmt.Println(m["Haras"])

//    m:=make(map[string]int)

//    m["age"]=30
//    m["price"]=30

//    fmt.Println(m["age"])
    // fmt.Println(len(m))

    // delete(m,"age")
	// clear(m)

	// fmt.Print(m)

	// m := map[string]int{"Price":23,"age":21,"Phones":3}
	// fmt.Print(m)

//    k, ok := m["age"]
//     fmt.Println(k)
//    if ok {
// 	fmt.Println("All Ok")
//    }else  {
// 	fmt.Println("Not Ok")
//    }

m1 := map[string]int{"Price":23,"age":21,"Phones":3}
m2 := map[string]int{"Price":23,"ae":21,"Phones":3}
  
 fmt.Print(maps.Equal(m1,m2))



}