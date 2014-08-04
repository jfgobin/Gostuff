package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nothing(id int64) {
	var i int32
	rep := rand.Int31n(10)+5
	wa  := rand.Int31n(5)
	for i=0;i<rep;i++ {
		fmt.Printf("Boucle %d avant %d secondes d'attente\n",id,wa)
		time.Sleep(time.Duration(wa)*time.Second)
	}
}

func main() {

	go nothing(1)
	go nothing(2)
	go nothing(3)
	time.Sleep(200*time.Second)
}
