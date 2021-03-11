package main

import (
	"fmt"

	"github.com/tsrines/go-examples/server"
)

type currentRecord struct {
}

func (currentRecord) check() {

}

func main() {
	s := []int{}
	for _, i := range s {
		fmt.Println("got here")
		i++
	}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)

	}
	server.Run()
}
