package main

import (
	"fmt"
)

func init() {
	fmt.Println("init...")
}

func main() {
	// %d 表示整型数字，%s 表示字符串

	var name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	fmt.Println(s)
}
