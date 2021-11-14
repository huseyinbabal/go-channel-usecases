package countingsemaphores

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (b Bar) serveCustomer(c int) {
	log.Print("customer#",c, " enters the bar")
	seat := <-b
	log.Print("++customer#",c," drinks at seat#",seat)
	time.Sleep(time.Second*time.Duration(2+rand.Intn(6)))
	log.Print("--customer#",c," frees sear#",seat)
	b<-seat

}

func Run() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar,10)
	for i:=0;i<cap(bar24x7);i++{
		bar24x7<-Seat(i)
	}

	for i:=0;;i++{
		time.Sleep(time.Second)
		go bar24x7.serveCustomer(i)
	}

	for{time.Sleep(time.Second)}
}
