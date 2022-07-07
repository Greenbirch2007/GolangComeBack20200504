package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	i := 0
	ch := make(chan string, 0)
	defer func() {
		close(ch)
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			p(time.Now().Unix())
			p("当前i的值为:", i)
			i++
			select {
			case m := <-ch:
				p("输出为:", m)
				break
			default:
				p("执行了default代码块")
			}

		}
	}()
	time.Sleep(time.Second * 4)
	ch <- "stop"
}
