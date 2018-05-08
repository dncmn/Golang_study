package package_learn

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func Test_Ioutil(T *testing.T){
	byte,err:=ioutil.ReadFile("../../testFiles/test_1.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(byte))
}
