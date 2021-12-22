package main

import "fmt"

// 接受者的方法：接受者可以是任意类型，包括结构体
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {   函数体 }
// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//方法① 值类型的接收者
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
	fmt.Printf("%d岁！\n", p.age)
}

//方法②  指针类型的接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}


//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

func main() {
	p1 := NewPerson("wsl", 25)
	p1.Dream()
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
	p1.SetAge2(30) // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 25

	var m1 MyInt
	m1.SayHello()
	m1=100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}