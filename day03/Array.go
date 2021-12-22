package main

import "fmt"

//数组是值类型 var name [len] type
var arr0 [5]int = [5]int{1, 2, 3}
var arr1  = [5]int{1, 2, 3}
var arr2  = [...]int{1, 2, 3}
var arr3  = [5]string{3:"hello",4:"world"}

func main() {
	fmt.Println(arr0,arr1,arr2,arr3)

	a:=[3]int{1,2}
	b:=[...]int{3,5}
	c:=[...]struct{
		name string
		age uint8
	}{
		{"user1",10},
		{"user2",20},
	}

	fmt.Println(a,b,c)
	fmt.Println(len(a),cap(b))  //长度
	fmt.Println("------------------------------------------------------------------")

	//切片是引用类型 只是少了声明数值
	var s1 []int

	s2 := []int{}

	//make([]type,len,cap)
    var s3 []int=make([]int ,2)

	slice := []int{0, 1, 2, 4, 8: 100}
    fmt.Println(slice,len(slice),cap(slice))

	var s4 []int = make([]int, 3, 3)

	s5 := []int{1, 2, 3}

    fmt.Println(s1,s2,s3,s4,s5)
	fmt.Println("--------------------------------------------------------------------")


  //map[KeyType]ValueType
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//声明map时直接填充元素
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)

	//contain :  value, ok := map[key]
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//遍历 k,v
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}


	//遍历k
	for k := range scoreMap {
		fmt.Println(k)
	}
}
