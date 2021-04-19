package main

import (
	"fmt"

	"github.com/Volfee/gollections/deque"
)

func main() {
	queue := deque.New()
	queue.Append(1)
	fmt.Println(queue.First())
}
