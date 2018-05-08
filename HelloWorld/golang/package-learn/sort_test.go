package package_learn

import (
	"fmt"
	"sort"
	"testing"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }           // 返回长度
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] } // 交换位置
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }      // 比较大小

func Test_Type_Convert(T *testing.T) {
	array := []int{9, 0, 3, 1, 7}
	fmt.Println(IntSlice(array))
	aa := sort.IntSlice(array)
	sort.Sort(aa)
	fmt.Println(aa)

	fmt.Println(sort.SearchInts(array, 0))

}

func Test_SortPackage(T *testing.T) {

	// sort int
	array := []int{4, 9, 1, 0, 2}
	//sort.Sort(array)
	fmt.Println(sort.IntsAreSorted(array))
	sort.Ints(array)
	fmt.Println(array)
	fmt.Println(sort.IntsAreSorted(array))

	// sort string
	str := []string{"ddd", "uuu", "bbb", "aaa"}
	sort.Strings(str)
	fmt.Println(str)

	// sort float
	farrary := []float64{1.1, 3.0, 7, 0, 2}
	sort.Float64s(farrary)
	fmt.Println(farrary)
}
