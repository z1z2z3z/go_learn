##### Golang sync.Mutex 与 sync.RWMutex的介绍

两者的关系：RWMutex(读写锁)是基于Mutex(互斥锁)实现的

**1. sync.Mutex**

sync.Mutex 用于多个 goroutine 对共享资源的互斥访问。使用要点如下：
（1）使用 Lock() 加锁，Unlock() 解锁；
（2）对未解锁的 Mutex 使用 Lock() 会阻塞；
（3）对未上锁的 Mutex 使用 Unlock() 会导致 panic 异常。

**2. sync.RWMutex**

读写锁就是一个可以并发读但是不可以并发写的锁

sync.RWMutex 用于读锁和写锁分开的情况。使用时注意如下几点：
（1）RWMutex 是单写多读锁，该锁可以加多个读锁或者一个写锁；
（2）读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁；
（3）写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占；
（4）适用于读多写少的场景。


[参考链接](https://blog.csdn.net/K346K346/article/details/90476721)


