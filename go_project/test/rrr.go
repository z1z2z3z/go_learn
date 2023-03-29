/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
)

func mainr() {
	ch := make(chan int)
	go test(ch)
	if v, ok := <-ch; ok {
		fmt.Println("get val: ", v, ok)
	}
	// for {
	// 	if v, ok := <-ch; ok {
	// 		fmt.Println("get val: ", v, ok)
	// 	} else {
	// 		break
	// 	}

	// }

}

func test(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
