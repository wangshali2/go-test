package main

//Go语言中通过struct来实现面向对象。
import (
	//"encoding/json"
	"encoding/json"
	"fmt"
)

//类型定义
type NewInt int

//类型别名
type MyInt = int

//1.定义结构体： type 类型名 struct {  字段名 字段类型  ... }
type Person struct {
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

}
