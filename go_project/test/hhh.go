/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
)
var wgg sync.WaitGroup
func main3() {
	// var wg *sync.WaitGroup
	// wg := new(sync.WaitGroup)
	wgg.Add(10)
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go doSomething(i, ch)
	}
	
	for i := 0; i < 10; i++ {
		dd := <-ch
		fmt.Println("from ch:" + strconv.Itoa(dd))
	}
	wgg.Wait()
	fmt.Println("all done")
}

func doSomething(index int, ch chan int) {
	defer wgg.Done()
	//fmt.Println("start done:" + strconv.Itoa(index))
	//time.Sleep(20 * time.Millisecond)
	ch <- index
}
