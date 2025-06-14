package main

import (
	"fmt"
	//"time"
	// "go/doc"
	// "math/rand"
	// "time"
	// "time"
)

// send
// func processNum(numChan chan int){
// 	for num := range numChan {
// 		 fmt.Println("Processing number ",  num)
// 		 time.Sleep(time.Second)
// 	}

// }

// receive
// func sum(result chan int, num1 int, num2 int){
//   numResult := num1 + num2
//   result <- numResult
// }

// goroutine synchronizer
// func task(done chan bool){
//   defer func()  {done<-true}()

//   fmt.Println("processing...")
// }


// func emailSender(emailChan chan string, done chan bool){
// 	defer func()  {done<-true}()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main(){

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func ()  {
	 chan1 <- 10	
	}()

	go func ()  {
	 chan2 <- "Harash"	
	}()

	for i:=0; i<2; i++ {
		select {
		case chan1val :=<-chan1 :
			fmt.Println("received data from chan1 ", chan1val)
		case chan2val :=<-chan2 :
			fmt.Println("received data from chan2 ", chan2val)	
			
		}
	}

	// emailChan := make(chan string, 100)
    // done := make(chan bool)

	// go emailSender(emailChan,done)

	// for i:=0; i<4; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com",i)
		
	// }

	// fmt.Println("done sending.")

	// // This is important
    // close(emailChan)
	// <- done // block

	// emailChan <- "1@kslmledlfm"
	// emailChan <- "ksdknkn;c"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
   
	
	 
	// done := make(chan bool)
	// go task(done)

	// <- done // block


// 	result := make(chan int)
//    go sum(result,4,5)
//    res := <- result // blocking

//    fmt.Print(res)


// 	numChan := make(chan int)

//    go processNum(numChan)

//    for {
// 	 numChan <- rand.Intn(100)
//    }

//    numChan <- 5
  
//    time.Sleep(time.Second * 0)

	// messageChan := make(chan string)

	// messageChan <- "ping" // <- data channels ka ander send kr raha hai // channles blocking

	// mesg := <- messageChan // <- channels se data receive kr raha hai
   
	// fmt.Print(mesg)

}