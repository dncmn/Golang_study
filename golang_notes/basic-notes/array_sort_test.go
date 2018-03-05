package sort_demo2

import (
	"fmt"
	"testing"
)

func Test_Array_sort(T *testing.T) {
	array := []int{3, 5, 2, 0, 9, 7}
	//n := len(array)
	//array = bubble_sort(array, n) // 冒泡排序--内部不稳定排序
	//array = select_sort(array, n) // 选择排序-内部不稳定排序
	//array = merge_sort(array, n) // 归并排序--内部排序稳定排序
	//array = insert_srot(array, n) // 插入排序--内部稳定排序
	array = shell_sort(array) // 希尔排序--内部稳定排序
	fmt.Println(array)
}

func shell_sort(array []int) []int {

	if len(array) == 1 {
		return array
	}

	mid := len(array) / 2
	left := shell_sort(array[:mid])
	right := shell_sort(array[mid:])

	return group_sort(left, right)
}

func group_sort(left, right []int) (result []int) {

	i, j := 0, 0

	lf := len(left)
	lr := len(right)

	for i < lf && j < lr {
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

func insert_srot(array []int, n int) []int {

	for i := 1; i < n; i++ {

		index := i - 1
		key := array[i]

		for index >= 0 && array[index] > key {
			array[index+1] = array[index]
			index--
		}

		array[index+1] = key
	}

	return array
}

func merge_sort(array []int, n int) []int {

	for step := n / 2; step > 0; step /= 2 { // 将原数组分组
		for i := step; i < n; i++ { // 每一组进行排序
			key := array[i]
			index := i - step

			for index >= 0 && array[index] > key {

				array[index+step] = array[index]
				index -= step
			}

			array[index+step] = key
		}

	}

	return array
}

func select_sort(array []int, n int) []int {

	for i := 0; i < n-1; i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[index] {
				index = j
			}
		}

		if index != i {
			array[i], array[index] = array[index], array[i]
		}
	}
	return array
}

func bubble_sort(array []int, n int) []int {

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j+1] < array[j] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
