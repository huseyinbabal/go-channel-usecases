package onetoonenotification

import (
	"crypto/rand"
	"fmt"
	"log"
	"sort"
)

func Run() {
	values := make([]byte, 32*1024*1024)

	_, err := rand.Read(values)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i]<values[j]
		})
		done<- struct{}{}
	}()

	<-done
	fmt.Println(values[0],values[len(values)-1])
}
