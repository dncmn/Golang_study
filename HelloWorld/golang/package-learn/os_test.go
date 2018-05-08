package package_learn

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func Test_Exec_Demo1(T *testing.T) {

	// 拼接cmd struct
	/*
		只是对cmd.Path和cmd.Args进行了赋值,前提是对这个wget进行了验证,判断是不是可执行文件
		判断是不是可执行文件的依据就是:从$PATH下面查找,看看对应的目录下面是否存在这个exectable binary file
	*/
	cmd := exec.Command("wget", "http://www.baidu.com")
	if err := cmd.Run(); err != nil {
		fmt.Println("exex command error", err)
		return
	}

}

// test get environment param from $PATH
func Test_GetEnvironment(T *testing.T) {
	fmt.Println(os.Getenv("GOPATH"))
}

// test range value in $PATH
func Test_FilePath(T *testing.T) {
	path := os.Getenv("PATH")

	// 找出了path中的环境变量
	for _, dir := range filepath.SplitList(path) {
		fmt.Println(dir)
	}
}

func Test_LookPath(T *testing.T) {
	//res, err := exec.LookPath("ssh")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(res)

	//fmt.Println(os.Environ())
	//
	//for key, value := range os.Environ() {
	//	fmt.Printf("key=%v,value=%v", key, value)
	//}
}
