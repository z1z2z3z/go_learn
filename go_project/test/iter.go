/*
 * @Author: zzz
 * @Date: 2023-02-14 11:04:37
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-03-13 17:12:09
 * @Description: 请填写简介
 */
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

type AA []int

func (a AA) Len() int {
	return len(a)
}
func (a AA) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AA) Less(i, j int) bool {
	return a[i] < a[j]
}

func mainSort() {

	var k int = 8
	AAA(&k)
	fmt.Println("DFg", &k)

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	// sort 排序任意数据结构 因为可以自己写方法
	hh := []int{6, 4, 5, 9, 12, 1, 5, 78}

	// 左边的几种方式都是重新创建了切片
	ff := hh[3:] // :数 ；省略冒号(一个数) ；省略冒号后面的数
	fmt.Println(ff)

	sort.Sort(AA(hh))
	// sort 自己的排序方法(int float string)
	sort.Ints(hh)
	fmt.Println(people)
	// sort.Sort(ByAge(people))
	fmt.Println(people)
	b.Add1(a)
	fmt.Println(b.Result())

	num, ok := b.(*Integer)
	if ok {
		fmt.Println(*num)
	} else {
		fmt.Println(111)
	}

	// // switch x.(type) 断言
	// var value interface{} // Value provided by caller.
	// switch str := value.(type) {
	// case string:
	// 	fmt.Println(str)
	// 	return //type of str is string
	// case int:
	// 	fmt.Println(str)
	// 	return //type of str is int
	// }
	// // 语句switch中的value必须是接口类型，变量str的类型为转换后的类型。

}

type Integer int

func (a Integer) Less1(b Integer) bool {
	return a < b
}
func (a *Integer) Add1(b Integer) {
	*a += b
}
func (a *Integer) Result() Integer {
	return *a
}

type LessAdder interface {
	Less1(b Integer) bool
	Add1(b Integer)
	Result() Integer
}

var a Integer = 1

// *Integer 才实现了接口的所有方法
var b LessAdder = &a

func AAA(r *int) {
	*r = 5
}
