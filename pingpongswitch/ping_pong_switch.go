package pingpongswitch

import (
	"fmt"
	"os"
	"time"
)

type Ball uint8

func play(playerName string, table chan Ball, serve bool) {
	var receive, send chan Ball
	if serve{
		receive,send = nil,table
	}else {
		receive,send=table,nil
	}

	var lastValue Ball = 1
	for {
		select {
		case send <- lastValue:
		case value := <-receive:
			fmt.Println(playerName, value)
			value += lastValue
			if value < lastValue {
				os.Exit(0)
			}
			lastValue = value
		}
		receive, send = send, receive
		time.Sleep(time.Second)
	}
}

func Run() {
	table:=make(chan Ball)
	go play("A:",table, false)
	play("B:",table, true)
}