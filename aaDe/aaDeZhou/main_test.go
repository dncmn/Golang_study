package main

import (
	"time"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"aaDe/aaDeZhou/models"
	"testing"
)

func Test_go(t *testing.T) {

	for i := 0; i < 1; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				go sendGet(j)
				time.Sleep(time.Microsecond * 200)
			}
		}()

	}

	time.Sleep(time.Second * 300)
}

func sendGet(i int) {
	var res models.Result
	res.Name = fmt.Sprint("tom", i)
	res.Password = fmt.Sprint("pwd", i)
	btr, err := json.Marshal(res)
	if err != nil {
		log.Println("转换json失败", err)
		return
	}
	uri := fmt.Sprint("http://127.0.0.1:8080/testSng?data=", string(btr))
	log.Println("uri=====================", uri)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("获取数据失败.......", err)
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	// log.Println("解析数据失败......", err)
	//}
	//
	//log.Println(string(body))
}