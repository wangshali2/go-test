package main

import (
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v2/taosSql"
	"strconv"
	"time"
)

func main() {
	var taosuri = "root:taosdata@/tcp(172.31.0.151:6030)/"
	taos, err := sql.Open("taosSql", taosuri)

	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return
	}
	defer taos.Close()

	_, err = taos.Exec("create database if not exists wsl2")

	if err != nil {
		fmt.Println("failed to create database, err:", err)
		return
	}

	taos.Exec("use wsl2")
	taos.Exec("create table if not exists tb (ts timestamp, a int,b nchar(10))")

	var j JsonMetaData
	j.Timestamp = 1639989549
	j.Step = 4

	_, err = taos.Exec(fmt.Sprintf("insert into tb values( %d, %d,'e')", (j.Timestamp)*1000, j.Step))
	_, err = taos.Exec("insert into tb values(" + strconv.FormatInt(time.Now().UnixMilli(), 10) + ", 0,'a')(now+1s,1,'a')(now+2s,2,'b')(now+3s,3,'c')")
	if err != nil {
		fmt.Println("failed to insert, err:", err)
		return
	}

	rows, err := taos.Query("select * from tb2")
	if err != nil {
		fmt.Println("failed to select from table, err:", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var r struct {
			ts time.Time
			a  int
		}
		err := rows.Scan(&r.ts, &r.a)
		if err != nil {
			fmt.Println("scan error:\n", err)
			return
		}
		fmt.Println(r.ts, r.a)
	}
}

type JsonMetaData struct {
	Metric      string `json:"metric"`
	Endpoint    string `json:"endpoint"`
	Timestamp   int64  `json:"timestamp"`
	Step        int64  `json:"step"`
	Value       map[string]interface{}
	CounterType string `json:"counterType"`
	Tags        string `json:"tags"`
}
