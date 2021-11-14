package firsresponsewins

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {
	rand.Seed(time.Now().UnixNano())

	startTime:= time.Now()

	c:= make(chan int32, 5)

	//1=>2s, 2=>1s,3=>4s,4=>1.5s,5=>5s
	for i:=0;i<cap(c);i++ {
		go source(c)
	}

	rnd:=<- c

	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}

func source(c chan<- int32) {
	ra,rb:=rand.Int31(),rand.Intn(3)+1
	time.Sleep(time.Duration(rb)*time.Second)
	c<-ra
}
