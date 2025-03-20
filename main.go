package main

import (
	"fmt"
	"github.com/go-pprof-practices/country"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func main() {
	// 只有1个cpu线程来执行Goroutine，串行
	runtime.GOMAXPROCS(1)
	//  Go 运行时（runtime）中的一个函数，它用于控制互斥锁（mutex）争用事件在互斥锁（mutex）性能分析（profile）中的采样频率
	//  控制 mutex 争用事件（即多个 Goroutine 竞争同一个锁）被记录到互斥锁性能分析中的比例
	//  rate = 1，则所有争用事件都会被记录
	//  rate = 10，则大约每 10 次争用事件记录 1 次
	runtime.SetMutexProfileFraction(1)
	// 阻塞事件采样评率，记录所有阻塞事件
	runtime.SetBlockProfileRate(1)

	// pprof
	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("启动pprof")
		}
		os.Exit(0)
	}()

	for {
		for _, v := range country.AllCountries {
			v.Exist()
		}
		time.Sleep(time.Second)
	}
}
