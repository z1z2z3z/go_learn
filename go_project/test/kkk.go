/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"time"
)

// 代码片段5
   func main66() {
	ch := make(chan int)

	go writeChan(ch)

	for {
		val, ok := <-ch
		fmt.Println("read ch: ", val)
		if !ok {
			break
		}
	}

	time.Sleep(time.Second)
	fmt.Println("end")
}

func writeChan(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
	}
	defer close(ch)
}