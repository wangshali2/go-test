package main

//Go语言中通过struct来实现面向对象。
import (
	"encoding/json"
	"fmt"
)

//类型定义
type NewInt int

//类型别名
type MyInt = int

//1.定义结构体： type 类型名 struct {  字段名 字段类型  ... }
type person struct {
	name, city string
	age        int8
}

//Tag
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	//结构实例化：  var 结构体实例 结构体类型
	var p1 person

	p1.name = "wsl"
	p1.age = 18
	p1.city = "sh"
	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}

	//2，匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "mlf"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//3.指针类型结构体
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%T\n", p2)  //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

	fmt.Println("---------------------------------------初始化---------------------------------------------")

	var p4 person              //初始化①
	fmt.Printf("p1=%#v\n", p4) //p1=main.person{name:"", city:"", age:0}

	//使用键值对初始化
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}

	p8 := &person{
		"pprof.cn",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8.name) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}
	fmt.Println("---------------------------------------------Tag---------------------------------------")
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}

	fmt.Println("-------------------------------------------构造函数-----------------------------------------")
	p9 := newPerson("wsl", "测试", 90)
	fmt.Printf("%#v\n", p9)

}

//构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
