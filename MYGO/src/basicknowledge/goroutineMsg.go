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

//无缓冲channel 示例，w无缓冲通道，必须有接收才可以发送，否则一直阻塞
func aha1() {
	ch := make(chan int)
	// 无缓冲区的通道，没有任何接收，100就发送不出去
	ch <- 100
	fmt.Println("此时的ch是无缓冲区的，将会发生死锁，因为没有地方接收")
}

func aha2() {
	ch := make(chan int)
	//此处有接收者，并且是异步的，所以不会阻塞
	f := func(x chan int) {
		<-ch
	}
	go f(ch)
	ch <- 100
}

//有缓冲通道，超过容量就会阻塞

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

//单项通道，只能读或者只能写，多用于函数参数，便于阅读，便于管理，便于控制

// select 多路复用，同一时间，可以对多个通道进行发送和接收操作
func aha3() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	//如果case1和case2都满足条件，则随机执行一个，跟代码顺序无关
	//select 只能执行一次，所以一般跟for配合使用
	select {
	case ch1 <- 100:
		fmt.Println("111")
	case <-ch2:
		fmt.Println("222")
	default:
		fmt.Println("ahahhahaha !")
	}
}

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

//通道是线程安全的

// 并发控制，锁
// 数据竞争
// 定义一个互斥锁 var lock sync.Mutex
// 加锁 lock.Lock()
// 释放锁 lock.Unlock()
// 定义一个 读写互斥锁 var sync.RWMutex，读比写频率高很多的时候，可以使用读写锁

//sync.Map  原生的Map不是并发安全的，并发场景下推荐sync.Map
// sync.Once，闭包的应用
func aha4(x int) func() {
	return func() {
		fmt.Printf("因为once的do方法中的参数不能有参数，使用闭包可以间接设置参数:%d", x)
	}
}

func aha5() {
	f := aha4(994903)
	o := sync.Once{}
	o.Do(f)
}

// 原子操作，go语言中的基础类型，共语言提供了sync/atomic，为整数类型提供了原子操作
// eg: i++ 不是并发安全的，使用atomic.AddInt64()就是线程安全的，自己使用lock也可以实现，但是性能低于内置的原子操作
