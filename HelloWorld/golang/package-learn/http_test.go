package package_learn

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func Test_http(T *testing.T) {
	//client := http.DefaultClient
	//
	//resp, err := client.Get("https://www.baidu.com/")
	////http.Client{}
	//resp, err = http.Get("https://www.baidu.com/")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//resp.Header = map[string][]string{}
	//
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header)
	//fmt.Println(resp.Request.Method)

	req := &http.Request{}
	req.Method = "get"
	req.RemoteAddr = "https://www.baidu.com"
	params := make(map[string][]string)
	params["host"] = []string{"192.168.1.204"}
	req.Header = params
	fmt.Println(req.Header)

}

func Test_GetLocalIP(T *testing.T) {
	ids, err := LocalIPv4s()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ids)
}

func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	fmt.Println("first ips=", addrs)

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

func Test_Get(T *testing.T) {
	src := "http://qa-api.game.snaplingo.com/api/game_app/v2/anonymous/user_title/titles?userID=123456"
	resp, err := http.Get(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

/*
	发送一个get请求
				并且对请求体部分添加了比如"Content-type","token"等进行验证。
*/
func Test_ClientRequestGet(T *testing.T) {
	src := "http://qa-api.game.snaplingo.com/api/game_app/v2/anonymous/user_title/titles?userID=123456"
	req, err := http.NewRequest("", src, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("token", "c25hcGxpbmdvOnNuYXBsaW5nbw")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func Test_ClientRequestPost(T *testing.T) {
	client := &http.Client{}
	src := "http://qa-api.game.snaplingo.com/api/game_app/v2/anonymous/user_title/titles"

	params := url.Values{"userID": {"123456"},
		"titleID":  {"50001"},
		"isUsed":   {"false"},
		"progress": {"1"},
	}

	body, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", src, strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("创立post请求失败")
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "c25hcGxpbmdvOnNuYXBsaW5nbw")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
