package main

import "fmt"

var p = fmt.Println

func main() {
	ch := make(chan int, 10)
	ch <- 66666
	select {
	case ch <- 32424:
		p("通道的值为:", <-ch)
		p("channel vaule is:", <-ch)
	default:
		p("通道已经被阻塞")
	}
}
