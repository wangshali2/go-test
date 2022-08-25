package main

import "fmt"

//切片是对数组的抽象,动态数组
func main() {
	//切片是引用类型 只是少了声明数值
	var s1 []int
	fmt.Println(s1)

	s2 := []int{}
	fmt.Println(s2)

	//make([]type,len,capacity)
	var s3 []int = make([]int, 2)
	fmt.Println(s3)

	slice := []int{0, 1, 2, 4, 8: 100}
	fmt.Println(slice, len(slice), cap(slice))

	var s4 []int = make([]int, 3, 3)
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	fmt.Println(s5)
}
