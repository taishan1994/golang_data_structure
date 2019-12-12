package sort

import (
	"fmt"
)

//SelectSort 选择排序
func SelectSort(arr *[7]int) {
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		index := i
		for j := i + 1; j < len(arr); j++ {
			if (*arr)[j] < tmp {
				tmp = (*arr)[j]
				index = j
			}
		}
		if index != i {
			(*arr)[index], (*arr)[i] = (*arr)[i], (*arr)[index]
		}
		fmt.Printf("第%d次选择后的结果是：%v\n", i, *arr)
	}
}

//InsertSort 插入排序
func InsertSort(arr *[7]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := (*arr)[i]
		inserIndex := i - 1
		for inserIndex >= 0 && (*arr)[inserIndex] > insertVal {
			(*arr)[inserIndex+1] = (*arr)[inserIndex]
			inserIndex--
		}
		//插入
		if (inserIndex + 1) != i {
			(*arr)[inserIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后的结果是：%v\n", i, *arr)
	}
}

//BubbleSort 冒泡排序
func BubbleSort(arr *[7]int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if (*arr)[j] > (*arr)[i] {
				(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			}
		}
		fmt.Printf("第%d趟的结果为：%v\n", len(arr)-1-i, *arr)
	}
}

//QuickSort 快速排序
func QuickSort(left int, right int, arr *[7]int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	tmp := 0
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		tmp = arr[l]
		arr[l] = arr[r]
		arr[r] = tmp
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}
}
