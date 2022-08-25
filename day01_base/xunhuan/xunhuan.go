package main

import "fmt"

func main() {
	var a int = 10

	// if
	if a < 20 {
		println("a < 20\n")
	}
	println("a的值：%d 哼", a)

	//默认匹配后直接break，需要执行后面的case用fallthrough
	var grade string = "B"
	var mark int = 90

	// switch
	switch mark {
	case 90:
		grade = "A"
		fallthrough
	case 80:
		grade = "C"

	}
	println(grade)

	//type switch 判断interface的类型
	var x interface{}

	switch x.(type) {
	case nil:
		println(x)
	case string, bool:
		println()
	default:
		println("unknow")
	}

	// select
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	//for [init]; condition; [post] { }
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	println(sum)

	for sum <= 10 {
		sum += sum
	}

	for sum <= 10 {
		sum += sum
	}

	/*	for {
		sum++
	}*/ //无限循环

	//For i,[i] 循环
	str := []string{"wsl", "hello"}
	for i, s := range str {
		println(i, s)
	}

	num := [6]int{1, 2, 3, 5}
	for i, s := range num {
		println(i, s)
	}

	//跳转到指定的代码行
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
	}

}
