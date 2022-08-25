package main

import (
	"fmt"
)

//声明数组:  数组是值类型 var name [len] type
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3}
var arr2 = [...]int{1, 2, 3} //数组长度不确定，可以使用 ... 代替数组的长度
var arr3 = [5]string{3: "hello", 4: "world"}

func main() {
	fmt.Println(arr0, arr1, arr2, arr3)

	a := [3]int{1, 2}

	b := [...]int{3, 5}

	c := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(a, b, c)
	fmt.Println(len(a), cap(b)) //长度

}
