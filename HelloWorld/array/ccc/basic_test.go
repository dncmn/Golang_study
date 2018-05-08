package ccc

import (
	"fmt"
	"testing"
)

func Test_Array_Sort(T *testing.T) {
	array := []int{8, 0, 2, 5, 1}
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
			array[min], array[i] = array[i], array[min]
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
			if array[low] > array[high] {
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
