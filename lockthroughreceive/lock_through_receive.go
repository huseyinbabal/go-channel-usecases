package lockthroughreceive

import "fmt"

func Run() {
	mutex := make(chan struct{}, 1)
	mutex<- struct{}{}

	counter:=0
	increase:= func() {
		<-mutex
		counter++
		mutex<- struct{}{}
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
