package main

import (
	"encoding/json"
	"fmt"
)

func main() {
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
		fmt.Println("逮到了", v)
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

	fmt.Println("--------------------------------map(k,object)---------------------------------------")

	m := make(map[string]interface{}) //Map<String,Object>
	m["name"] = "simon"
	m["age"] = 12
	m["addr"] = "China"
	fmt.Println(m)

	data, _ := json.Marshal(m) //序列化对象  object->二进制
	fmt.Println("序列化：", data)

	m1 := make(map[string]interface{})
	_ = json.Unmarshal(data, &m1) //反序列化
	fmt.Println("反序列化：", m1)

	if m1["name"] != nil {
		fmt.Println(m1["name"].(string))
	}

	var m2 = map[string]int{"abc": 3, "ccd": 4} //Map<String,int>
	i := m2["abc"]
	fmt.Println(i)

	const name, age1 = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age1)
	fmt.Println(s)

}
