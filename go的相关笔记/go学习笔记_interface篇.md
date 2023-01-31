### Go 接口 interface （接口定义、接口实现、接口调用、值接收者、指针接收者）

##### 1. 接口的定义

接口是和调用方的一种约定，它是一个高度抽象的类型，不用和具体的实现细节绑定在一起。接口要做的是定义好约定，告诉调用方自己可以做什么，但不用知道它的内部实现，这和我们见到的具体的类型如 int、map、slice 等不一样。

接口的定义和结构体稍微有些差别，虽然都以 type 关键字开始，但接口的关键字是 interface，表示自定义的类型是一个接口。

也就是说 person 是一个接口，它有两个方法 sayName() string 和 sayAge() int，整体如下面的代码所示:

```go
type person interface {
    sayName() string
    sayAge() int
}
```

针对 `person` 接口来说，它会告诉调用者可以通过它的 `sayName()` 方法获取姓名字符串，通过它的 `sayAge()` 方法获取年龄，这就是接口的约定。至于这个字符串怎么获得的，长什么样，接口不关心，调用者也不用关心，因为这些是由接口实现者来做的。

> 接口特点：
> 
> - 接口只有方法声明、没有实现，没有数据字段
> - 接口可以匿名嵌入其它接口，或者嵌入到结构中

> 接口是用来定义行为的类型，这些被定义的行为不由接口直接实现，而是由用户定义的类型实现，**一个实现了这些方法的具体类型是这个接口类型的实例。**

**如果用户定义的类型实现了某个接口类型声明的一组方法，那么这个用户定义的类型的值就可以赋给这个接口类型的值。这个赋值会把用户定义的类型存入接口类型的值。**

##### 2. 接口的实现

接口的实现者必须是一个具体的类型，以 `student` 结构体为例，

```go
type student struct {
    name string
    age  int
}
```

让它来实现 `person` 接口，如下代码所示：

```go
func (s student) sayName() string {
    fmt.Printf("name is %v\n", s.name)
    return s.name
}

func (s student) sayAge() int {
    fmt.Printf("name is %v\n", s.age)
    return s.age
}
```

给结构体类型 `student` 定义一个方法，这个方法和接口里方法的签名（名称、参数和返回值）一样，这样结构体 `student` 就实现了 `person`接口。

> 注意：如果一个接口有多个方法，那么需要实现接口的每个方法才算是实现了这个接口。

## [参考链接](https://blog.csdn.net/wohu1104/article/details/111242632?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522166865393816782429711078%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fblog.%2522%257D&request_id=166865393816782429711078&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~blog~first_rank_ecpm_v1~rank_v31_ecpm-27-111242632-null-null.nonecase&utm_term=go%E7%AC%94%E8%AE%B0&spm=1018.2226.3001.4450)

### 如何判断一个对象有没有实现interface

```go
package main

import "fmt"

type LessAddr interface {
	Less() Integer
	Add() Integer
}

type Integer int

func(a Integer) Less() Integer {
	return a
}

func(a *Integer) Add() Integer {
	return *a + 1
}

func main() {
	var b Integer = 1 
	var a *Integer = &b
	if _, ok := interface{}(a).(LessAddr); ok{
		fmt.Println("是interface接口")	
	}
		
   
}
///  interface{}.(type)  判断类型
///  interface{}("DFGD").(string)  判断为string类型
```

---

#### 对象对interface赋值

```go
package main

import "fmt"

type LessAddr interface {
	Less(b Integer) bool
	Add(b Integer) Integer
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//  需要改变量值就必须要用指针
func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

func main() {
	// 因为随便定义这样的类型都满足接口
	var c Integer = 1
	// interface 就不要指针了
	var inter LessAddr = &c
	fmt.Println(inter.Add(22))
}
```
