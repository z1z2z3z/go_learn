package main

import "fmt"

var s []string

type Person1 struct {
	age int
}

func maintt() {

	slice := make([]int,5,5)
	slice = append(slice, 3)
	fmt.Println(slice,cap(slice))

	i := 65
	// UTF-8 编码中，十进制数字 65 对应的符号是 A
	m := string(i)
	fmt.Println(m)

	var mm = make(map[string]int) //A
	mm["a"] = 1
	if v, ok := mm["b"]; ok { //B
		fmt.Println(v)
	}

	fmt.Println("1", f1())
	fmt.Println("2", f2())
	fmt.Println("3", f3())

	person := &Person1{28}

	// 1. 会把 28 缓存在栈中 输出28
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person1) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()
	fmt.Println("2323", f(3))

	person.age = 29

	defer func() {
		fmt.Println("qwr", "1")
	}()
	if aa == true {
		fmt.Println("yuu", "2")
		return
	}
	defer func() {
		fmt.Println("ttt", "3")
	}()

}

func f(n int) (r int) {
	defer func() {
		r += n
		fmt.Println("f", r)
		recover()
	}()

	var f func()

	f = func() {
		r += 2
	}
	defer f()
	return n + 1
}

var aa bool = true

func f1() (result int) {
	defer func() {
		// 这里result是全局的
		result++
	}()
	return 0
} //1  re先被赋值为0  然后re++ 变成了1

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
} //5

///当defer被声明时，其参数就会被实时解析
func f3() (r int) {

	////  因为是值传递  不会改变原本r=1
	// defer func(r int) {
	// 	r = r + 5
	// 	fmt.Println(r)
	// }(r)

	////  如果改成指针   就可以进行改变了
	defer func(r *int) {
		*r = *r + 5
		fmt.Println(*r)
	}(&r) ///  这里的值就已经定了就是默认值0
	return 1
} //1
