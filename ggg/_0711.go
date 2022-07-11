

//创建：数组和切片

arr1 := new([len]type)

slice1 := make([]type, len)

// 数组和切片

// arr1 := [...]type{i1, i2, i3, i4, i5}

// arrKeyValue := [len]type{i1: val1, i2: val2}

// var slice1 []type = arr1[start:end]

// 1. 18.3 映射
// 创建： map1 := make(map[keytype]valuetype)

// 初始化： map1 := map[string]int{"one": 1, "two": 2}

// （1）如何使用for或者for-range遍历一个映射：

// for key, value := range map1 {
// …
// }
// （2）如何在一个映射中检测键key1是否存在：

// val1, isPresent = map1[key1]

// 返回值：键key1对应的值或者0, true或者false

// （3）如何在映射中删除一个键：

// delete(map1, key1)



1. 18.4 结构体
创建：

type struct1 struct {
    field1 type1
    field2 type2
    …
}
ms := new(struct1)
初始化：

ms := &struct1{10, 15.5, "Chris"}
当结构体的命名以大写字母开头时，该结构体在包外可见。 通常情况下，为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体（参考例10.2）：

ms := Newstruct1{10, 15.5, "Chris"}
func Newstruct1(n int, f float32, name string) *struct1 {
    return &struct1{n, f, name} 
}


1. 18.5 接口
（1）如何检测一个值v是否实现了接口Stringer：

if v, ok := v.(Stringer); ok {
    fmt.Printf("implements String(): %s\n", v.String())
}
（2）如何使用接口实现一个类型分类函数：

func classifier(items ...interface{}) {
    for i, x := range items {
        switch x.(type) {
        case bool:
            fmt.Printf("param #%d is a bool\n", i)
        case float64:
            fmt.Printf("param #%d is a float64\n", i)
        case int, int64:
            fmt.Printf("param #%d is an int\n", i)
        case nil:
            fmt.Printf("param #%d is nil\n", i)
        case string:
            fmt.Printf("param #%d is a string\n", i)
        default:
            fmt.Printf("param #%d’s type is unknown\n", i)
        }
    }
}

