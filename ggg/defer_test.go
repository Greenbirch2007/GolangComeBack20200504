
// http://books.studygolang.com/gobyexample/defer/






//在 closeFile 后得到一个文件对象，我们使用 defer通过 closeFile 来关闭这个文件。
//这会在封闭函数（main）结束时执行，就是 writeFile 结束后。


package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	f := createFile("C:\\mytest\\t.txt")
	defer closeFile(f)
	writeFile(f)
}


