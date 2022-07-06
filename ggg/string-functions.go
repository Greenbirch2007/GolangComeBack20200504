package main

//http://books.studygolang.com/gobyexample/string-functions/
import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p(s.Contains("test", "es"))
	p(s.Count("test", "t"))
	p(s.HasPrefix("test", "te"))
	p(s.HasSuffix("test", "es"))
	p(s.Index("test", "e"))
	p(s.Join([]string{"a", "b"}, "-"))
	p(s.Repeat("a", 5))
	p(s.Replace("foo", "o", "0", -1))
	p(s.Replace("foo", "o", "0", 1))
	p(s.Split("a-b-c", "-"))
	p(s.ToLower("AFD"))
	p(s.ToLower("test"))
	p()
}
