package _11_sorts

import "fmt"

// BubbleSort 冒泡算法的实现
func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		isMove := false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				isMove = true
			}
		}
		if !isMove {
			break
		}
	}
	return a
}

// InsertionSort 插入排序的实现
func InsertionSort(a []int) []int {
	if len(a) < 0 {
		return nil
	}
	for i := 1; i < len(a); i++ {
		// value 为要插入的值
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
	return a
}

// SelectionSort 选择排序的实现
func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

// PrintArray 打印数组
func PrintArray(a []int) {
	result := ""
	for i, v := range a {
		result = result + fmt.Sprintf("a[%d]=%d,", i, v)
	}
	fmt.Println(result)
}
