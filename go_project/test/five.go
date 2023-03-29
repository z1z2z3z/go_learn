package main

import "fmt"

type MyInt1 int

// 别名
type MyInt2 = int

const (
	x = iota
	_
	y
	z = "zz"
	k
	m
	r = 3.15
	h
	p = iota
)

func GetValue() interface{} {
	return 1
}

func main44() {
	ii := GetValue()
	a := ii.(int) //变成对应类型
	dd := float32(a) + 2.3
	fmt.Println(a,dd)
	switch ii.(type) { //获取对应的类型
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

	var i int = 0
	var i1 MyInt1 = i // MyInt1 类型 和int类型
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
