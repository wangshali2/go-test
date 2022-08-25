package main

//Go语言中通过struct来实现面向对象。
import (
	"fmt"
)

//1.定义结构体： type 类型名 struct {  字段名 字段类型  ... }
type Person struct {
	name, city string
	age        int8
}

func main() {

	fmt.Println(Person{"aaa", "上海", 18})
	fmt.Println(Person{
		name: "bbb",
		city: "北京",
	})

	//1.实例化结构体
	var p1 Person

	p1.name = "wsl"
	p1.age = 18
	p1.city = "sh"
	fmt.Printf("p1=%v\n", p1)  //p1={wsl 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"wsl", city:"北京", age:18}

	//2，匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "mlf"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//3.指针类型结构体
	var p2 = new(Person)

	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%T\n", p2)  //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}

	//使用键值对初始化
	p5 := Person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p8 := &Person{
		"pprof.cn",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8.name) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}

	printPerson(&p1)
}

func printPerson(use *Person) {
	fmt.Println(use.age)
	fmt.Println(use.name)
	fmt.Println(use.city)
}
