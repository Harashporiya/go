package main

import (
	"fmt"
	"time"
)

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time  // Time -> nanosecond precision

}

// receiver type
func (o *order) changeStatus(status string){
  o.status = status
}

func (o order) getAmount() float32{
	return o.amount 
}

func main(){
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
    myorder := order{
		id: "1",
		amount: 23.4,
		status: "received",
		createdAt: time.Now(),
	}

	myorder.changeStatus("confrimed")
	
	fmt.Print(myorder.getAmount())

	// myorder2 := order{
	// 	id: "2",
	// 	amount: 123.4,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// myorder.createdAt = time.Now()

	// fmt.Println(myorder.amount)

	// fmt.Println("Order Struct ",myorder)
}