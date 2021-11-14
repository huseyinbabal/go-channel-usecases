package sendonlychannels

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan<- int32)  {

	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a+b*b
}

func Run() {
	now := time.Now()
	rand.Seed(time.Now().UnixNano())

	ra,rb:= make(chan int32),make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(rb)
	fmt.Println(sumSquares(<-ra,<-rb))
	fmt.Println(time.Since(now))
}