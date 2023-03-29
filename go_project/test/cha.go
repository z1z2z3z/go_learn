/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"time"
)

// 只能写入：var send chan<- int
// 只能读取：var recv <-chan int
func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")

		t := time.NewTicker(time.Millisecond * 500)

		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case <-t.C:
				fmt.Println("Working .")
			}
		}

	}()
	// return
}

func main1() {
	stopCh := make(chan struct{})
	worker(stopCh)

	time.Sleep(time.Second * 2)

	close(stopCh)

	time.Sleep(time.Second)
	fmt.Println("mian exit")
}
