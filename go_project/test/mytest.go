/*
 * @Author: zzz
 * @Date: 2023-02-09 09:37:49
 * @LastEditors: zzz
 * @LastEditTime: 2023-02-14 10:35:28
 * @Description: 请填写简介
 */
/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"runtime"

	// "runtime"
	"sync"
	// "time"
)

var wg sync.WaitGroup

func mainm() {

	var ch = make(chan string)
	quitChan := make(chan bool)
	for i := 0; i < 10; i++ {

		go sum(i, i+10, ch) // 这种用for循环i=0;i<n;i++  获取
	}

	// 多次写入数据，for 读取数据时，写入者注意关闭 channel 不然会死锁
	go func() {
		// 如果channel关闭了  会自动退出循环
		for val := range ch {
			fmt.Println(val)
			fmt.Println(runtime.NumGoroutine())
			/// 对于无缓存的  可以通过协程数来关闭channel
			/// 因为无法判断最后一个生产者  故在消费者出关闭channel
			if runtime.NumGoroutine() == 2 {
				close(ch)
			}
		}
		quitChan <- true
	}()

	<-quitChan

	///  消费者不能close(ch)

	// go func() {
	// 	defer wg.Done() ///如果不关  就无法done 就会死锁
	// 	for i := 0; i < 11; i++ {
	// 		if i == 10 {
	// 			close(ch)
	// 		} else {

	// 			e, ok := <-ch
	// 			fmt.Println(e, ok)
	// 		}
	// 	}

	// }()

	// fmt.Println("ghghghghghgh")
	// wg.Wait()

}

func sum(start, end int, ch chan string) {
	// defer wg.Done()
	var sum int = 0
	for i := start; i < end; i++ {
		sum += i
	}
	ch <- fmt.Sprintf("Sum from %d to %d is %d", start, end, sum)
}
