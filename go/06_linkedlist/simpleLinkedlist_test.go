package _6_linkedlist

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	dataArray := []string{"a", "b", "c"}
	linkedList, err := NewLinkedList(dataArray, 5)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("------初始化数据 ")
	linkedList.Print()

	fmt.Println("------调用LRU a")
	linkedList.LRU("a")
	linkedList.Print()

	fmt.Println("------调用LRU d,e;使链表满容量")
	linkedList.LRU("d")
	linkedList.LRU("e")
	linkedList.Print()

	fmt.Println("------调用LRU f;使链表超容量")
	linkedList.LRU("f")
	linkedList.Print()
}
