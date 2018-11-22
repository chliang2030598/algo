package _5_array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	//测试越界插入
	fmt.Println("------测试越界插入")
	testdata := []int{0, 1, 2, 3, 4, 5}
	testArray, err := NewArray(testdata, uint(10))
	if err != nil {
		t.Fatal(err.Error())
	}
	testArray.Print()
	err = testArray.Insert(10, 7)
	if err == nil {
		t.Fatal("test for index out bound is faild")
	}
	testArray.Print()

	//测试插入超过容量测试
	fmt.Println("------测试插入超过容量测试")
	testArray, _ = NewArray(testdata, uint(6))
	testArray.Print()
	err = testArray.Insert(10, 1)
	if err == nil {
		t.Fatal("test for index out of cap is faild")
	}
	testArray.Print()

	//测试插入结尾
	fmt.Println("------测试插入结尾")
	testArray, _ = NewArray(testdata, uint(10))
	testArray.Print()
	err = testArray.Insert(10, testArray.Length)
	if err != nil {
		t.Fatal(err.Error())
	}
	testArray.Print()

	//测试插入开头
	fmt.Println("------测试插入开头")
	testArray, _ = NewArray(testdata, uint(10))
	testArray.Print()
	err = testArray.Insert(10, 0)
	if err != nil {
		t.Fatal(err.Error())
	}
	testArray.Print()

	//测试插入中间
	fmt.Println("------测试插入中间")
	testArray, _ = NewArray(testdata, uint(10))
	testArray.Print()
	err = testArray.Insert(10, 2)
	if err != nil {
		t.Fatal("test for index out of cap is faild")
	}
	testArray.Print()
}
func TestDelete(t *testing.T) {
	// 测试删除开头
	data := []int{0, 1, 2, 3, 4, 5}

	fmt.Println("------测试删除开头")
	testArray, _ := NewArray(data, uint(10))
	testArray.Print()
	err := testArray.Delete(0)
	if err != nil {
		t.Fatal(err.Error())
	}
	testArray.Print()
	// 测试删除任意
	fmt.Println("------测试删除任意")
	testArray, _ = NewArray(data, uint(10))
	testArray.Print()
	err = testArray.Delete(testArray.Length - 1)
	if err != nil {
		t.Fatal(err.Error())
	}
	testArray.Print()

	// 测试删除不存在索引
	fmt.Println("------测试删除不存在索引")
	testArray, _ = NewArray(data, uint(10))
	testArray.Print()
	err = testArray.Delete(testArray.Length + 1)
	if err == nil {
		t.Fatal("Delete not exist index error")
	}
	testArray.Print()
}

func TestFind(t *testing.T) {
	// 测试查询不存在索引
	data := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("------测试查询不存在索引")
	testArray, _ := NewArray(data, uint(10))
	_, err := testArray.Find(testArray.Length)
	if err == nil {
		t.Fatal("test for not exist index find error")
	}
	// 测试查询存在的索引
	fmt.Println("------测试查询存在的索引")
	result, err := testArray.Find(testArray.Length - 1)
	if err != nil {
		t.Fatal("test for  exist index find error")
	}
	fmt.Printf("testArray[%d]=%d\n", testArray.Length-1, result)
}
