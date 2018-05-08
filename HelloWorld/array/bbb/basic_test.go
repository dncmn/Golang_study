package bbb

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"path"
	"testing"
	"time"
)

func Test_GenrageBase64Str(T *testing.T) {
	byte, err := ioutil.ReadFile("manan.wav")
	if err != nil {
		fmt.Println("read file error", err)
		return
	}

	str := base64.StdEncoding.EncodeToString(byte)
	fmt.Println(str)
}

func Test_Packge_Sort(T *testing.T) {
	name := ReFileName(".wav")
	fmt.Println(name)
}

func ReFileName(filePath string) string {
	dir := path.Dir(filePath)
	ext := path.Ext(filePath)
	cruTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%d_snap_%d", cruTime, GenerateRangeNum(10000, 99999)))
	newName := h.Sum(nil)
	return fmt.Sprintf("%s/%x%s", dir, newName, ext)
}

// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	return randNum
}

func Test_GetDataFromMap(T *testing.T) {
	var data = map[string]struct{}{
		"mp3": struct{}{},
		"wav": struct{}{},
	}

	fileName := "mpd3"
	_, ok := data[fileName]

	fmt.Println(ok)
}

func Test_Float32Demo(T *testing.T) {
	var data float32 = 0.93685
	fmt.Println(math.Round(float64(data * 100)))
	fmt.Println(data)
}

func Test_Base64Encoding(T *testing.T) {
	s := "helloWorld"
	byte := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("base64编码后")
	fmt.Println(byte)
	result, err := base64.StdEncoding.DecodeString(byte)
	if err != nil {
		fmt.Println("decode error", err)
		return
	}

	fmt.Println(string(result))
}

func Test_Path_Basic2(T *testing.T) {
	// path.Clean--将所有的双斜线变成单斜线
	url := "http://www.baidu.com/aaa.wav"
	fmt.Println(path.Clean(url))

	// path.Dir--找出文件的根目录
	url = "//home//mn/notes//golang_notes//hello.go"
	fmt.Println(path.Dir(url)) ///home/mn/notes/golang_notes

	// path.IsAbs---判断路径是否是绝对路径

	url = "./hello.go"
	abs := path.IsAbs(url) //  len(url)!=0 && path[0]!="/"
	fmt.Println(abs)

	fmt.Println(path.ErrBadPattern)

}

func Test_Path_Basic(T *testing.T) {
	sourceUrl := "http://www.baidu.com/team1/helloWorld.wav"
	fileName := path.Base(sourceUrl)
	extName := path.Ext(sourceUrl)
	fmt.Printf("fileName=%s,extName=%s\n", fileName, extName)

	s1 := "/home/mn"
	s2 := "notes/golang_notes"
	fmt.Println(path.Join(s1, s2))

	s3 := "//home//mn//notes//golang_notes"
	dir, file := path.Split(s3)
	fmt.Printf("dir=%s,file=%s\n", dir, file)
	fmt.Println(path.Split(sourceUrl))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(path.Clean("//home//mn"))

	fmt.Println("64" + path.Clean(path.Base(s3)))
}

func Test_Array_Sort(T *testing.T) {

	array := []int{7, 0, 12, 4, 8}
	//quick_sort(array, 0, len(array)-1)
	//array = merge_sort(array)
	//insert_sort(array, len(array))
	//shell_sort(array, len(array))
	select_sort(array, len(array))
	fmt.Println(array)
}

func select_sort(array []int, n int) {
	for i := 0; i < n-1; i++ {

		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		if i != min {
			array[i], array[min] = array[min], array[i]
		}
	}
}

func shell_sort(array []int, n int) {
	for step := n / 2; step > 0; step /= 2 {
		for i := step; i < n; i++ {
			key := array[i]
			j := i - step
			for j >= 0 && array[j] > key {
				array[j+step] = array[j]
				j -= step
			}
			array[j+step] = key
		}
	}
}

func insert_sort(array []int, n int) {
	for i := 1; i < n; i++ {

		key := array[i]
		j := i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}
}

func merge_sort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	mid := len(array) / 2
	left := merge_sort(array[:mid])
	right := merge_sort(array[mid:])
	return group_sort(left, right)
}

func group_sort(left, right []int) (result []int) {

	i, j, ll, lr := 0, 0, len(left), len(right)

	for i < ll && j < lr {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return
}

func quick_sort(array []int, start, end int) {
	low := start
	high := end
	key := array[start]
	for {
		for low < high {
			if array[high] < key {
				array[low] = array[high]
				break
			}
			high--
		}

		for low < high {
			if array[low] > key {
				array[high] = array[low]
				break
			}
			low++
		}

		if low >= high {
			array[low] = key
			break
		}
	}

	if low-1 > start {
		quick_sort(array, start, low-1)
	}

	if high+1 < end {
		quick_sort(array, high+1, end)
	}
}

func Test_Base64(T *testing.T) {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	content := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("after base64=", content)
	bytes, err = base64.StdEncoding.DecodeString(content)
	fmt.Println()
	fmt.Println(string(bytes))
}

func Test_ReadFile(T *testing.T) {

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	fmt.Println("fileContent=", string(data))
}
