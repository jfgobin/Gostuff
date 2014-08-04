package main

import (
	"fmt"
	"time"
)


func processors(id int, ci chan int) {
// Gets a number from the channel, displays it and then wait
// if the number is -1, exits
	for ifc := 0; ifc > -1; ifc = <- ci {
		fmt.Printf("Routine %d got %d\n",id, ifc)
		time.Sleep(time.Duration(id)*time.Second)
	}  
	fmt.Printf("Routine %d terminating\n",id)
}

func main() {
	var i int

	ci := make(chan int, 10)
	go processors(1,ci)
	go processors(2,ci)
	go processors(3,ci)

	for i=0; i<10 ; i++ {
		ci <- i
	}
	ci <- -1
	ci <- -1
	ci <- -1
	fmt.Println("Going to sleep for 60 seconds ... waiting for child to die")
	time.Sleep(60*time.Second)
}
