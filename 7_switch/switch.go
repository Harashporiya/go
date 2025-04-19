package main

import "fmt"

// import "time"

func main(){
	// simple switch
	// i:=2

	// switch i {
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// case 3:
	// 	fmt.Println(3)
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it's workday")
	// }

	// time show fmt.Printf(time.Now().String())

	// type show

	 whoAmI := func (i interface{})  {
		switch i.(type) {
		case int:
			fmt.Println("Its an Interger")
		case string:
			fmt.Println("Its an string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("other")
		}
	}

	whoAmI(true)
}