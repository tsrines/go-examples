package server

import (
	"fmt"
)

// Run starts our api.
func Run() {
	fmt.Println("Got here")
	NewDeckFromFile()

}
