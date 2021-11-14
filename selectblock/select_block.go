package selectblock

import "runtime"

func doSomething() {
	for {
		runtime.Gosched()
	}
}
func Run() {
	go doSomething()
	go doSomething()
	select {}
}
