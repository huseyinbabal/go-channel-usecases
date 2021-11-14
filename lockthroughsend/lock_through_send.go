package lockthroughsend

import "fmt"

func Run() {
	mutex := make(chan struct{}, 1)
	counter := 0
	increase := func() {
		mutex <- struct{}{}
		counter++
		<-mutex
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done;<-done
	fmt.Println(counter)
}
