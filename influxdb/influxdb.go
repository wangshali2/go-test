package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"time"
	//"time"
)

const (
	MyDB     = "wsl"
	username = "admin"
	password = "admin"
)

func main() {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://172.31.0.151:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		fmt.Println(err)
	}

	//database
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	if err != nil {
		fmt.Println(err)
	}

	tags := map[string]string{"name": "xc"}
	fields := map[string]interface{}{
		"id":   1,
		"sex":  1,
		"pass": 0707,
	}

	//table
	pt, err := client.NewPoint("myuser", tags, fields, time.Now())
	if err != nil {
		fmt.Println(err) //
	}

	bp.AddPoint(pt) //

	if err := conn.Write(bp); err != nil {
		fmt.Println(err)
	}
}
