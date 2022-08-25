package main

import (
	"encoding/base64"
	"fmt"
	"github.com/imroc/req"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type InfluxdbItem struct {
	Measurement string                 `json:"metric"`
	Tags        map[string]string      `json:"tags"`
	Fileds      map[string]interface{} `json:"fileds"`
	Timestamp   int64                  `json:"timestamp"`
}

func main() {
	//tags := make(map[string]string)
	//tags["t1"] = "1"
	//tags["t2"] = "2"
	//fileds := make(map[string]interface{})
	//fileds["time"] = atom.Time
	//var data InfluxdbItem
	//data.Measurement = "test"
	//data.Tags = tags
	//data.Fileds = fileds
	//data.Timestamp = time.Now().UnixMilli()

	//// influx db -> tdengine
	//jsonData, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//reader := strings.NewReader(string(jsonData))
	//sql="insert into "

	/*reader := strings.NewReader("INSERT INTO  wsl2.tb VALUES (now, 'wsl');")

	url1, _ := url.Parse("http://172.31.0.151:6041/rest/sql")
	request, err := http.NewRequest("POST", url1.String(), reader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "text/plain")
	request.SetBasicAuth("root", "taosdata")

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	println(response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)

	println(string(body))

	err = response.Body.Close()
	if err != nil {
		return
	}*/

	//println(request.Body)

	//httpPost()
	//httpGet()

	//tdenginetHttpPost()
	//tdenginetHttpGet()
	//influxdbToTdengine()

	//get()

}
func influxdbToTdengine() {

	str := base64.StdEncoding.EncodeToString([]byte("root:taosdata")) //加密:cm9vdDp0YW9zZGF0YQ==
	println(str)
}

func get() {
	url1 := "http://172.31.0.151:6041/rest/sql"
	header := make(http.Header)

	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("Authorization", "Basic cm9vdDp0YW9zZGF0YQ==")
	r, _ := req.Post(url1, header, "create table if not exists wsl.tb (ts timestamp,metric nchar(50));")
	fmt.Println(r)

}

func tdenginetHttpGet() {

	reader := strings.NewReader("select * from wsl.meters;")
	url1, _ := url.Parse("http://172.31.0.151:6041/rest/sql")

	request, err := http.NewRequest("GET", url1.String(), reader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth("root", "taosdata") //todo 权限校验

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	println(response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)

	println(string(body))

	err = response.Body.Close()
	if err != nil {
		return
	}
}

//TDengine 直接通过 HTTP POST 请求 BODY 中包含的 SQL 语句来操作数据库
func tdenginetHttpPost() {

	reader := strings.NewReader("INSERT INTO  wsl.meters VALUES (now, 'wsl');")
	url1, _ := url.Parse("http://172.31.0.151:6041/rest/sql")

	request, err := http.NewRequest("POST", url1.String(), reader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded") //text/plain   //application/x-www-form-urlencoded
	request.SetBasicAuth("root", "taosdata")                                //todo 权限校验

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	println(response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)

	println(string(body))

	err = response.Body.Close()
	if err != nil {
		return
	}
}

//封装好的http
func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

//封装好的http
func httpPost() {
	resp, err := http.Post("usdp-o3tbdsfp-monitor1:6041/rest/sql", "application/json", strings.NewReader("INSERT INTO  wsl2.tb VALUES (now, 'wsl2');")) // Content-Type post请求必须设置
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
