package main

import (
	"fmt"
)

type Counter struct {
	n int
}

func (counter *Counter) count() {
	counter.n++
	fmt.Printf("counter = %d\n", counter.n)	
}

// func main(){
// 	counter := new(Counter)
// 	counter.count()
// 	counter.count()
// 	counter.count()
// }