package main

import "fmt"

//声明方法
func main() {

	//声明变量
	var name string = "wsl"
	fmt.Println("name:", name)

	var a, b int = 1, 2
	fmt.Println(a, b)

	// 因式分解关键字：用于声明全局变量
	var (
		c int
		d bool
	)
	fmt.Println(d, c)

	//:= 是一个声明语句,在函数体使用
	//name:="mlf" 如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
	name1 := "wsl"
	fmt.Println("name1", name1)

	//常量
	const age = 18
	fmt.Println(age)

	const (
		e = iota
		f = 2
		g = iota
	)

	fmt.Println(e, f, g)

}
