/*
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Row struct {
	Id int `json:"id"`
}

type Report struct {
	Name string `json:"name"`
}

type Bag struct {
	Reports []Report `json:"reports"`
	I       int      `json:"i"`
	J       int      `json:"j"`
}

func GetRows() (rows []Row) {

	rows = make([]Row, 0)
	for i := 0; i < rand.Intn(10); i++ {
		rows = append(rows, Row{Id: i})
	}

	return
}

func GetReports(m, n int) (reports []Report) {
	reports = make([]Report, 0)
	for i := 0; i < rand.Intn(100); i++ {
		reports = append(reports, Report{Name: fmt.Sprintf("r: in %d rows[%d] report %d", m, n, i)})
	}
	return
}

func maingg() {

	var (
		cLimit = 3 // 线程限制
		wg     = sync.WaitGroup{}
	)

	c := make(chan int, cLimit)
	// c2 := make(chan int, 10)
	cReport := make(chan Bag, 100)
	defer close(c)
	// defer close(cReport)

	go func() {
		log.Println("in Rec Start")
		defer func() {
			log.Println("in Rec Finish")
		}()
	Rec:
		for {
			select {
			case r, ok := <-cReport:
				if !ok {
					break Rec
					// c2 <- 1
				}
				log.Printf("rec %v from %d rows[%d] len=%d\n ", ok, r.I, r.J, len(r.Reports))
			
			// default:
			// 	time.Sleep(100*time.Microsecond)

			}
		}
	}()

	// 通过主协程 写入channel 以及 在每个子协程 读取channel 达到限制进程数量  并且可以完成大量的遍历任务
	// 比如有9个任务  但是我每次只想用3个子协程工作
	// 下面这部分便可以实现
	for i := 0; i < 10; i++ {
		rows := GetRows()
		wg.Add(1)
		// 通过主协程 写入channel
		c <- 1
		log.Printf("add %d rows=%d", i, len(rows))
		go func(rows []Row, i int) {
			log.Println("in ", i)
			for j := 0; j < len(rows); j++ {
				rs := GetReports(i, j)
				log.Printf("in %d rows[%d] get rs:%d to send\n", i, j, len(rs))
				cReport <- Bag{Reports: rs, I: i, J: j}
				time.Sleep(time.Second * time.Duration(rand.Intn(7)))
			}
			defer func() {
				wg.Done()
				log.Println("leave ", i)
			}()
			fmt.Println("gorun--------------", runtime.NumGoroutine())
			// 在每个子协程 读取channel
			<-c
		}(rows, i)

		if i%cLimit == 0 {
			fmt.Println("3个")
		}
	}
	log.Println("after loop")
	// Rec:
	// 	for {
	// 		select {
	// 		case r, ok := <-cReport:
	// 			if !ok {
	// 				break Rec
	// 			}
	// 			log.Printf("rec %v from %d rows[%d] len=%d\n ", ok, r.I, r.J, len(r.Reports))
	// 		}
	// 	}
	log.Println("after work")
	wg.Wait()
	close(cReport)
	// select {
	// case v, ok := <-c2:
	// 	fmt.Println("v,ok = ", v, ok)
	// 	break
	// }
	// log.Println("Finish")
}
