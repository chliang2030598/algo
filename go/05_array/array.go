package _5_array

import (
	"errors"
	"fmt"
)

// Arrayer 为Array的抽象接口，定义了Array需要实现的方法
type Arrayer interface {
	isFull() bool
	isEmpty() bool
	Insert(value int, n uint) error
	Delete(n uint) error
	Find(n uint) error
	Print()
}

// Array 自定义数组结构类型
type Array struct {
	Value    []int //数组值
	Length   uint  //数组长度
	Capacity uint  //数组容量
}

// NewArray 通过golang自带类型数组构建自定义数组结构类型
func NewArray(input []int, capacity uint) (*Array, error) {
	if capacity < uint(len(input)) {
		return nil, errors.New("the capacity is less the array length")
	}

	newArray := make([]int, len(input), capacity)
	newArray = append(newArray, input...)

	array := new(Array)
	array.Value = newArray
	array.Length = uint(len(array.Value))
	array.Capacity = uint(cap(array.Value))
	return array, nil
}

// 判断数组当前容量和数组长度的关系，如果length=Capacity-1，说明数组已满，无法插入
func (array *Array) isFull() bool {
	if array.Length == array.Capacity {
		return true
	}
	return false
}

func (array *Array) isEmpty() bool {
	if array.Length == 0 {
		return true
	}
	return false
}

// Insert 数组插入:value为插入值，n为插入位置
func (array *Array) Insert(value int, n uint) error {
	if array.isFull() {
		return errors.New("The array is full")
	}
	//判断插入位置属否合法
	if n > array.Length {
		return errors.New("The index is illegal")
	} else if n == array.Length {
		// 插入到末尾位置
		array.Value = append(array.Value, value)
		array.Length++
		return nil
	}
	// 插入到开头位置
	if n == 0 {
		//创建新数组，长度加1
		newArray := make([]int, array.Length+1, array.Capacity)
		//将元素插入到开头
		newArray = append(newArray, value)
		//将其他元素加入到新数组中
		newArray = append(newArray, array.Value[:]...)
		//将新的数组赋值给array对象
		array.Value = newArray

		array.Length++
		return nil
	}

	// 插入到中间位置
	newArray := make([]int, array.Length+1, array.Capacity)
	index := int(n)
	newArray = append(newArray, array.Value[:index]...)
	newArray = append(newArray, value)
	newArray = append(newArray, array.Value[index:]...)
	array.Value = newArray

	array.Length++
	return nil
}

// Delete 数组删除位置
func (array *Array) Delete(n uint) error {
	if array.isEmpty() {
		return errors.New("The array is Empty")
	}
	if n >= array.Length {
		return errors.New("The index is out of bound")
	}
	// 删除末尾元素
	if n == array.Length-1 {
		newArray := make([]int, array.Length-1, array.Capacity)
		newArray = append(newArray, array.Value[:array.Length-1]...)

		array.Value = newArray
		array.Length--

		return nil
	}

	// 删除首元素
	if n == 0 {
		newArray := make([]int, array.Length-1, array.Capacity)
		newArray = append(newArray, array.Value[1:]...)
		array.Value = newArray
		array.Length--
	}

	// 删除中间元素
	newArray := make([]int, array.Length-1, array.Capacity)
	index := int(n)
	newArray = append(newArray, array.Value[:index]...)
	newArray = append(newArray, array.Value[index+1:]...)

	array.Value = newArray
	array.Length--
	return nil
}

// Find 根据数据索引查找数组
func (array *Array) Find(n uint) (int, error) {
	if n > array.Length-1 {

	}
	return 0, nil
}

// Print 打印数组
func (array *Array) Print() {
	result := ""
	for i, v := range array.Value {
		temp := fmt.Sprintf("array[%d]=%d  ", i, v)
		result = result + temp
	}
	fmt.Printf("%s-----len=%d-----cap=%d\n", result, array.Length, cap(array.Value))
}
