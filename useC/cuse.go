/*
 * @Description: go调用c/c++  直接嵌套在go语言中
 */
package main

/* // C
#include <stdio.h>
int add(int a,int b) {
	return a*b;
} */
import "C"  // 这个必须顶在后面引入  不能有换行

import "fmt"

func main()  {
	fmt.Println(C.add(2,9))
}