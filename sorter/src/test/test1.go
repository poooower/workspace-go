package test

import (
	"fmt"
	"math/rand"
	"time"
)

type A struct {
}

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")

}

func Timeout(b chan bool) {
	s := time.Duration(rand.Int31n(6)) * time.Second
	time.Sleep(s)
	b <- true
}

func DoSomething(ch chan int) {
	s := time.Duration(rand.Int31n(6)) * time.Second
	time.Sleep(s)
	ch <- 1
}

func (a *A) Do() {
	// chn := make([]chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("make", i)
	// 	chn[i] = make(chan int)
	// 	go Count(chn[i])
	// }
	//
	// for i, ch := range chn {
	// 	fmt.Println("range", i)
	// 	<-ch
	// }

	for {
		ch := make(chan int)
		go DoSomething(ch)
		b := make(chan bool)
		go Timeout(b)

		select {
		case <-ch:
			fmt.Println("case <-ch:")
		case <-b:
			fmt.Println("timeout")
		}

		close(ch)
		close(b)
	}

	fmt.Println("finish")
}
