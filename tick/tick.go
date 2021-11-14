package tick

import (
	"fmt"
	"time"
)

func tick(d time.Duration) <-chan struct{}{
	c:=make(chan struct{},1)

	go func() {
		for {
			time.Sleep(d)
			select {
			case c<- struct{}{}:
			default:


			}
		}
	}()
	return c
}
func Run() {
	t := time.Now()
	for range tick(time.Second){
		fmt.Println(time.Since(t))
	}
}
