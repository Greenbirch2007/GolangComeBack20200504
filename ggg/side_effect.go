

//http://books.studygolang.com/the-way-to-go_ZH_CN/06.2.html
//1.4. 6.2.4 改变外部变量（outside variable）

//传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回。如下的例子，reply 是一个指向 int 变量的指针，通过这个指针，我们在函数内修改了这个 int 变量的数值。

package main

import (
    "fmt"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
    *reply = a * b
}

func main() {
    n := 0
    reply := &n
    Multiply(10, 5, reply)
    fmt.Println("Multiply:", *reply) // Multiply: 50
