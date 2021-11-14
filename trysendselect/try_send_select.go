package trysendselect

import "fmt"

func Run() {
	type Book struct {
		id int
	}

	bookshelf := make(chan Book, 3)

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case bookshelf<-Book{i}:
			fmt.Println("succeeded to put book",i)
		default:
			fmt.Println("failed to put book")

		}
	}
	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case book:=<-bookshelf:
			fmt.Println("succeeded to get book",book)
		default:
			fmt.Println("failed to get book")

		}
	}
}
