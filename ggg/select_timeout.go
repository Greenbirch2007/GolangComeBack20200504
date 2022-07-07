package main

import (
	"fmt"
	"time"
)

//https://shockerli.net/post/golang-select-time-implement-timeout/

// Go 采用 time.After 实现超时控制

// 场景:

// 假设业务中需调用服务接口A，要求超时时间为5秒，那么如何优雅、简洁的实现呢？

// 我们可以采用select+time.After的方式，十分简单适用的实现。
var p = fmt.Println

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"

	}()

	select {
	case res := <-ch:
		p(res)
	case <-time.After(time.Second * 1):
		p("timeout")
	}
}
