package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {

		select {
		case c <- x:
			//如果c可写，则执行下面语句
			x, y = y, x+y
		case <-quit:
			//如果 quit 可读，则进入该case
			fmt.Println("quit")
			return
		}

	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacii(c, quit)
}
