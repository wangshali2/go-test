package main

import "fmt"

func main() {

	//声明指针变量：指向一个值的内存地址
	var ip *int
	var a int = 20

	//&a 返回变量存储地址
	//*a 取出指针指向的内容
	ip = &a
	fmt.Println("ip的值为", &a)  //0xc00000a098
	fmt.Println("ip的值为", ip)  //0xc00000a098
	fmt.Println("ip的值为", *ip) //20

	var prt *int
	fmt.Println(prt) //nil 空指针
}
