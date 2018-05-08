package project_test

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func Test_Iap_Order(T *testing.T){

}

type GameTitle struct{

	Roletitledesc string `json:"roletitledesc"`
	Roletitleicon string `json:"roletitleicon"`
	Roletitleid uint64 `json:"roletitleid"`
	Roletitlelevel string `json:"roletitlelevel"`
}

type GameCourse struct{
	Courseid int `json:"courseID"`
	Levelids string  `json:"levelIDs"`
}


// 解析json文件,插入到数据库
func TestReadFile(T *testing.T){
	byte,err:=ioutil.ReadFile("/home/mn/notes/game_document/CourseConfig.json")
	if err!=nil{
		fmt.Println(err)
		return
	}

	var res =make(map[uint64]GameCourse)


	err=json.Unmarshal(byte,&res)
	if err!=nil{
		fmt.Println(err)
		return
	}


	for key,value:=range res{
		fmt.Printf("Id=%v,vaue=%v",key,value)


	}
}