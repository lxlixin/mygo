package basicknowledge

import (
	"fmt"
	"math/rand"
	"sync"
)

// go运行时 runtime(go帮我们跟操作系统交流，在Java中需要开发者自己实现，go中平台已经帮我们实现)
// runtime 相当于我们在Java中自己写的线程池
// go 语言中使用go创建一个goroutine，一个goroutine对应一个函数
// main函数也是一个goroutine在跑，函数执行完毕，goroutine就结束

//sync.WaitGroup  add done wait 协调，保证主程序结束的时候所有协程都结束

var wg sync.WaitGroup

func worker(i int) {
	//告诉 wg 我这个协程结束了
	defer wg.Done()
	fmt.Printf("%d号工人做了工作\n", i)
}

//Factory 工厂介绍sync.WaitGroup 用法
func Factory() {
	//这里有是个工人，一次把他们放进了工作
	for i := 0; i < 10; i++ {
		//wg 监控有一个工人就绪
		wg.Add(1)
		//工人开始工作
		go worker(i)
	}
	//wg 等待工人都工作结束
	wg.Wait()
	defer fmt.Println("工作都结束了！")
}

// go 并发使用通信而不是共享内存
// channel go中的通信类型 是引用类型，用make初始化，分配内存
// 使用make初始化的还有 slice，map
// 通道的操作：发送，接收，关闭
// 通道可以执行关闭操作，从一个关闭的通道中取值，可以去到对应类型的剩余值，直到去到对应的零值，往关闭的通道中放值会报错
// 无缓冲通道和缓冲通道

//ChanFunc channel 的一些基本操作
func ChanFunc() {
	// 定义并初始化通道
	c := make(chan int, 1)
	// 往通道中放数据
	c <- 10
	// 从通道中取数据
	ret := <-c
	fmt.Printf("从通道中拿到的数据：%d \n", ret)
	// 关闭通道
	close(c)
	// 从关闭的通道中获取数据
	ret2 := <-c
	fmt.Printf("从通道中拿到的数据：%d \n", ret2)
	//往关闭的通道中放值
	// c <- 9
}

// 用goroutine和channel实现生产者消费者问题

type item struct {
	id  int
	num int
}

type result struct {
	item
	res int
}

func producer(c chan item) {
	var id int
	for {
		if len(c) < 10 {
			id++
			it := item{id, rand.Intn(1000)}
			c <- it
		}
	}
}

func consumer(s chan result, cc chan item) {
	maxcount := rand.Intn(10)
	var count int
	for count < maxcount {
		var sum int
		c := <-cc
		var x = c.num
		for x > 0 {
			sum = sum + x%10
			x = x / 10
		}
		s <- result{c, sum}
	}
}

//ChannerTest 测试是否正确
func ChannerTest() {
	c := make(chan item)
	s := make(chan result)
	go producer(c)
	for {
		go consumer(s, c)
		ss := <-s
		fmt.Printf("结果是：第%d个生成的数字%d，和是%d \n", ss.id, ss.num, ss.res)
	}
}

// select 多路复用

// TestSelect 通过select 打印10之内的偶数
func TestSelect() {
	var c = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case c <- i:
		case x := <-c:
			fmt.Println(x)
		}
	}
}

//单项通道，只能读或者只能写，便于阅读，便于管理，便于控制

// 并发控制，锁
// 数据竞争
// 定义一个互斥锁 var lock sync.Mutex
// 加锁 lock.Lock()
// 释放锁 lock.Unlock()
// 定义一个 读写互斥锁 var sync.RWMutex，读比写频率高很多的时候，可以使用读写锁
