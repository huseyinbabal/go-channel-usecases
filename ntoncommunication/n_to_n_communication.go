package ntoncommunication

import (
	"log"
	"time"
)

type T = struct {

}

func worker(id int, ready <- chan T, done chan<- T) {
	<-ready

	log.Print("Worker#",id, " starts")

	time.Sleep(time.Second*time.Duration(id+1))
	log.Print("Worker#",id," job done")
	done<-T{}

}

func Run() {
	log.SetFlags(0)
	ready,done:= make(chan T),make(chan T)

	go worker(0,ready,done)
	go worker(1,ready,done)
	go worker(2,ready,done)
	time.Sleep(time.Second*3/2)

	ready <- T{};ready<-T{}; ready<-T{}

	<-done;<-done;<-done
}
