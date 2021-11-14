package encapsulatedchannel

import "fmt"

var counter = func(n int) chan<- chan<- int {
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				n++
			} else {
				request <- n
			}
		}
	}()
	return requests
}(0)

func Run() {
	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter<-nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)

	<-done;<-done

	request:=make(chan int, 1)


	counter <- request
	fmt.Println(<-request)
}
