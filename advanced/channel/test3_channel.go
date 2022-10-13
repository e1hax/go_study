package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {
			c <- i
		}

		//存放完毕后，关闭channel
		close(c)
	}()

	//循环从channel中取数据
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main Finished...")

}
