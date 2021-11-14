package requestwithtimeout

import (
	"errors"
	"fmt"
	"time"
)

func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	go func(chan int) {
		c<-5
	}(c)
	select {
	case data:=<-c:
		return data, nil
	case <-time.After(timeout):
		return 0,errors.New("timeout")


	}
}
func Run() {
	timeout, err := requestWithTimeout(time.Second * 5)
	fmt.Println(timeout)
	fmt.Println(err)
}
