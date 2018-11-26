package _11_sorts

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println("------测试冒泡排序")
	a := []int{3, 5, 4, 2, 1, 6}
	PrintArray(a)
	sortedArray := BubbleSort(a)
	PrintArray(sortedArray)
}

func TestInsertionSort(t *testing.T) {
	fmt.Println("------测试插入排序")
	a := []int{3, 5, 4, 2, 1, 6}
	PrintArray(a)
	sortedArray := InsertionSort(a)
	PrintArray(sortedArray)
}

func TestSelectionSort(t *testing.T) {
	fmt.Println("------测试选择排序")
	a := []int{3, 5, 4, 2, 1, 6}
	PrintArray(a)
	sortedArray := SelectionSort(a)
	PrintArray(sortedArray)
}
