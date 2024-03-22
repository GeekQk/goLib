package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"testing"
)

func TestRunTime(t *testing.T) {
	go func() {
		m := make(map[int]int)
		for i := 0; ; i++ {
			m[i] = i
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	<-sig
	//执行 `GODEBUG=gctrace=1 go run main.go`, 查看运行时的内存情况
}

func TestRunShed(t *testing.T) {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			runtime.Gosched() // 让出CPU时间片，重新等带安排任务
		}
	}("world")

	for i := 0; i < 2; i++ {
		fmt.Println("hello")
		runtime.Gosched() // 让出CPU时间片，重新等待安排任务
	}
}

var wg sync.WaitGroup

func TestRunExit(t *testing.T) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	wg.Wait() // 主goroutine等待子goroutine结束，主在结束
}

func TestSetCore(t *testing.T) {
	// 获取当前的GOMAXPROCS值
	originalGOMAXPROCS := runtime.GOMAXPROCS(0)
	fmt.Printf("Original GOMAXPROCS: %d\n", originalGOMAXPROCS)

	// 设置GOMAXPROCS为2，表示程序最多可以使用2个CPU核心
	runtime.GOMAXPROCS(2)

	// 执行一些并发的任务
	for i := 0; i < 4; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d is running on CPU: %d\n", id, runtime.NumCPU())
		}(i)
	}

	// 等待所有goroutine完成
	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func(id int) {
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	// 恢复原始的GOMAXPROCS值
	runtime.GOMAXPROCS(originalGOMAXPROCS)
	fmt.Printf("Restored GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	/*
		从Go 1.5版本开始，runtime.GOMAXPROCS 的设置实际上并不会限制goroutine的调度，而只是影响了Go运行时在调度goroutine时使用的P（Processor）的数量。
		这意味着即使你设置了 GOMAXPROCS 为1，Go仍然可以在多个CPU核心上调度goroutine，只是同时运行的goroutine数量受到了限制。
		因此，在Go 1.5及以后的版本中，runtime.GOMAXPROCS 主要用于控制并行执行的用户级goroutine的数量，而不是直接限制CPU核心的使用。
	*/
}
