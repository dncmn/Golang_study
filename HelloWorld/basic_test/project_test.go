package basic_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"testing"
)

func Test_Float2Str(T *testing.T) {
	var f1 float32 = 2.4

	fmt.Println(f1)
}

func Test_Path(T *testing.T) {
	demoUrl := "http://static.test.snaplingo.com/msg/audio/71384_201803041207385504671.mp3"
	fileName := path.Base(demoUrl)
	fmt.Printf("fileName is %v", fileName)
}

func Test_Golang_Download(T *testing.T) {

	str := "http://static.test.snaplingo.com/msg/audio/71384_201803041207385504671.mp3"
	resp, err := http.Get(str)
	if err != nil {
		fmt.Println("get file data error", err)
		return
	}

	fileName := path.Base(str)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read file data err", err)
		return
	}

	err = ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		fmt.Println("write data to file error", err)
		return
	}

	fmt.Println("final")
}
