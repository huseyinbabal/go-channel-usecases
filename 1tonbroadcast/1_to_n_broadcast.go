package _tonbroadcast

import (
	"fmt"
	"time"
)

func afterDuration(d time.Duration) <- chan struct{} {
	c:= make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func Run() {

	fmt.Println("Hi!")
	<- afterDuration(time.Second)
	fmt.Println("Hello!")
	<-afterDuration(time.Second)
	fmt.Println("Bye!")
}
