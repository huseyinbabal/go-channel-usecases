package ratelimit

import (
	"fmt"
	"time"
)

type Request interface {

}

func handle(r Request)  {
	fmt.Println(r.(int))
}

const RateLimitPeriod = time.Minute
const RateLimit = 30

func handleRequests(requests <-chan Request)  {
	quotas:= make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod/RateLimit)
		defer tick.Stop()

		for t:= range tick.C{
			select {
			case quotas<-t:
			default:

			}
		}
	}()

	for r:=range requests{
		<-quotas
		go handle(r)
	}
}

func Run() {
	requests:=make(chan Request)
	go handleRequests(requests)
	for i:= 0;;i++{
		requests<-i
	}
}
