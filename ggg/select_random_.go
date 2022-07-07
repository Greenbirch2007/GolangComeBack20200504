package main

import "fmt"

var p = fmt.Println

//Golang中select的四大用法
// https://blog.51cto.com/u_10919694/3089601

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch6 := make(chan int, 1)
	var i1, i2 int
	select {
	case i1 = <-ch1:
		p("接收到了管道1的一条数据:", i1)
	case ch2 <- i2:
		p("向管道2发送了一条数据:", i2)
	case i3, ok := <-ch3:
		if ok {
			p("收到了管道3的数据:", i3)
		} else {
			p("管道3已被关闭")
		}
	case ch6 <- 2342:
		p("向管道6发送了一条数据:", 271828)
	default:
		p("以上case皆不可运行,即没有进行通信")
	}
}
