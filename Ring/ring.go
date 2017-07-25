package main

import (
	"container/ring"
	"fmt"
	"time"
)

func main() {
	println("Ring Example")

	coffee := []string{"kenya", "guatemala", "ethiopia"}

	// create a ring and populate it with some values
	r := ring.New(len(coffee))
	for i := 0; i < r.Len(); i++ {
		r.Value = coffee[i]
		r = r.Next()
	}

	// print all values of the ring, easy done with ring.Do()
	r.Do(func(x interface{}) {
		fmt.Print(x, ",")
	})

	i := 0
	// .. or each one by one.
	for _ = range time.Tick(time.Second * 1) {
		r = r.Next()
		fmt.Println(r.Value)
		i++
		if i > 10 {
			break
		}
	}
}
