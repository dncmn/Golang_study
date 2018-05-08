package project

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTransLate(T *testing.T) {
	aa := ""
	res, err := strconv.Atoi(aa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Test_OpenGoroutine(T *testing.T) {
	//go func() {
	//
	//	// 1.从数据库表中判断是否有数据存在,没有的话,程序就休眠2分钟
	//	res := []int{1, 2, 3, 4, 5, 6}
	//
	//	// 查找数据的时候,就是记录过滤条件,只查询本地条件异常的
	//	for _, v := range res {
	//
	//		fmt.Println(v)
	//		// 2.判断当前记录如果尝试到了五次就跳过这一条记录,没有的话,就继续
	//
	//		// 3.再次验证,直接实验五次,失败的话,就更改数据库中的记录
	//		// 4.					成功的话,该如何处理???????
	//	}
	//}()

}

var ch = make(chan bool, 1)

func Test_ChannelNoCache(T *testing.T) {
	fmt.Println("before")
	<-ch
	close(ch)
	fmt.Println("end")
}

func Test_SendSign(T *testing.T) {
	ch <- true
}

func Producer(queue chan<- int) {

	for i := 0; i < 10; i++ {

		queue <- i

	}

}

func Consumer(queue <-chan int) {

	for i := 0; i < 10; i++ {

		v := <-queue

		fmt.Println("receive:", v)

	}

}

func Test_ProduceAndConsume(T *testing.T) {
	queue := make(chan int, 1)

	go Producer(queue)

	go Consumer(queue)

	time.Sleep(1e9) //让Producer与Consumer完成

}

func TestReadChannel(T *testing.T) {
	ch := make(chan int, 1)
	_, ok := <-ch

	if ok {
		fmt.Println("exist")
	}

	fmt.Println("not exist")
}
