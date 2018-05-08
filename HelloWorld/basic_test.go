package HelloWorld

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"testing"
)

func Test_UrlDemo(T *testing.T) {
	data := "http://www.baidu.com   da"

	escape := url.PathEscape(data)
	fmt.Println("pathEscape-result=", escape)
	res, err := url.PathUnescape(escape)
	if err != nil {
		fmt.Println("pathUnescape error", err)
	}
	fmt.Println(res)
}

func Test_Base64(T *testing.T) {
	str := base64.StdEncoding.EncodeToString([]byte("helloWorld"))

	fmt.Println("encode result=", str)
	array, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("decode error", err)
		return
	}
	fmt.Println(string(array))

}

func Test_Path_Join_Demo(T *testing.T) {
	//fmt.Println(path.Join("/home/mn", "notes/golang_notes"))
	str := "//hello.mp3"
	fmt.Println(strings.Replace(str, path.Ext(str), ".wav", -1))
}

func TestGetFileName(T *testing.T) {
	tmp := "http://www.badiu.com/team1/helloWorld.wav"
	fmt.Println(path.Base(tmp))
}

func Test_Path_Ext(T *testing.T) {
	src := "hello.txt"
	fmt.Println(path.Ext(src))
}

func Test_Run_Download_Command(T *testing.T) {
	cmd := exec.Command("wget", "-P", ".", "http://static.test.snaplingo.com/msg/audio/71384_201803041207385504671.mp3")
	err := cmd.Run()
	if err != nil {
		fmt.Println("获取资源失败:", err)
		return
	}

}

func Test_Url(T *testing.T) {
	// 1.测试url是否合法
	str := "http://asdfasdf.wav"

	_, err := url.ParseRequestURI(str)
	if err != nil {
		fmt.Println("请求的uri不是正确的")
		return
	}

	if !strings.Contains(str, ".wav") && !strings.Contains(str, ".mp3") {
		fmt.Println("uri的路径错误")
		return
	}

	fmt.Println("url合法")
	// 2.从指定的uri中下载视频

}

func Test_Strings_ContainsDemo(T *testing.T) {
	//fmt.Println(strings.Contains("hellWorld",";"))

	str := ""
	fmt.Println(len(strings.Trim(str, " ")))
	fmt.Println(str == "")
}

const (
	One int = iota + -1
	Two
	Three
	Four
)

func Test_Strings_Contains(T *testing.T) {
	msg := "helloWorld"
	son := "helloWorld"
	fmt.Println(strings.Contains(msg, son))
}

func Test_Map_Nil(T *testing.T) {
	OutPutResult(map[string]string{"one": "first"})
}

func OutPutResult(data interface{}) {
	fmt.Printf("the result is=%v", data)
}

func Test_Fmt(T *testing.T) {
	fmt.Println("helloWorld")
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
}

// Ip2long 转化ip成整形ip
func Ip2long(ip string) (ipInt int64) {
	ips := strings.Split(ip, ".")
	ip1, _ := strconv.ParseInt(ips[0], 10, 32)
	ip2, _ := strconv.ParseInt(ips[1], 10, 32)
	ip3, _ := strconv.ParseInt(ips[2], 10, 32)
	ip4, _ := strconv.ParseInt(ips[3], 10, 32)
	ipTmp := fmt.Sprintf("%08b%08b%08b%08b", ip1, ip2, ip3, ip4)
	ipInt, _ = strconv.ParseInt(ipTmp, 2, 64)
	return
}

func Test_Ip2long(T *testing.T) {
	ip := Ip2long("192.168.125.124")
	fmt.Println("final-ip=", ip)

}
