package main

import (
	"fmt"
)

//接口 它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
type Phone interface {
	call() //方法
}

type NokiaPhone struct {
}

type IPhone struct {
}

//方法实现
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

//方法实现
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
