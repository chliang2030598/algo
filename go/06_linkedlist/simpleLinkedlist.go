package _6_linkedlist

import (
	"errors"
	"fmt"
)

// 单链表的实现LRU的实现

// LinkedListNode 单链表节点的数据结构
type LinkedListNode struct {
	Data string          // 节点对应的数据,Data类型以string为例
	Next *LinkedListNode // 指向下一个节点的指针
}

// LinkedList 单链表的数据结构定义
type LinkedList struct {
	Head     *LinkedListNode // 头部节点，链表的基地址
	Length   int             // 链表的长度
	Capacity int             // 链表容量
}

// NewLinkedList 生成一个新的链表;限制链表长度在3以上，包含头，普通，尾部，不考虑特殊情况
func NewLinkedList(data []string, capacity uint) (*LinkedList, error) {
	if len(data) < 3 {
		return nil, errors.New("the linkedlist length less than three")
	} else if capacity < uint(len(data)) {
		return nil, errors.New("the data length less than capacity")
	}

	//先构建出相同长度的LinkedListNode节点，然后添加每个node的值和指针
	linkedListNodeArray := make([]LinkedListNode, len(data))
	for i := 0; i < len(data); i++ {
		//  处理尾部节点,将尾部节点的指针指向空
		if i == len(data)-1 {
			linkedListNodeArray[i].Data = data[i]
			linkedListNodeArray[i].Next = nil
		}
		linkedListNodeArray[i].Data = data[i]
		linkedListNodeArray[i].Next = &linkedListNodeArray[i+1]
	}

	linkedList := new(LinkedList)
	// 头部节点为linkedListNodeArray[0]
	linkedList.Head = &linkedListNodeArray[0]
	linkedList.Length = len(data)
	linkedList.Capacity = int(capacity)
	return linkedList, nil
}

// InsertAfterNode value表示要插入的值，nodevalue表示已经存在节点的值(如果值有重复，默认第一个节点)，方法表示在nodevalue后插入vaule的值
func (linked *LinkedList) InsertAfterNode(data string, nodevalue string) error {
	if linked.isFull() {
		return errors.New("the linkedlist capacity is full")
	}
	linkedListNode, err := linked.Find(nodevalue)
	if err != nil {
		return err
	}
	// 判断查找到的nodevalue节点是否为尾部节点
	if linkedListNode.Next == nil {
		newLinkedListNode := new(LinkedListNode)
		newLinkedListNode.Data = data
		//设置新的尾部节点
		newLinkedListNode.Next = nil
		//将原来的尾部节点指向新节点
		linkedListNode.Next = newLinkedListNode

		return nil
	}

	// 创建新的linkListNode,并将新的指针指向linkedListNode.Next
	newLinkedListNode := new(LinkedListNode)
	newLinkedListNode.Data = data
	newLinkedListNode.Next = linkedListNode.Next

	//将linkedListNode.Next重定向
	linkedListNode.Next = newLinkedListNode

	linked.Length++

	return nil
}

func (linked *LinkedList) isFull() bool {
	if linked.Length == int(linked.Capacity) {
		return true
	}
	return false
}

// Find 根据data值遍历单链表，找到相应的节点
func (linked *LinkedList) Find(data string) (*LinkedListNode, error) {
	result := findByData(data, linked.Head)
	if result == nil {
		return nil, errors.New("the linkedlistNode is not exist")
	}
	return result, nil
}

// 通过递归方式找到node节点，linked.Next == nil 表示到单链表尾部节点仍未找到值对应的节点
func findByData(value string, linkedListNode *LinkedListNode) *LinkedListNode {
	if linkedListNode.Data == value {
		return linkedListNode
	} else if linkedListNode.Next == nil {
		return nil
	}
	return findByData(value, linkedListNode.Next)
}

func findByPtr(startLinkedListNode, targetLinkedListNode *LinkedListNode) (*LinkedListNode, error) {
	if startLinkedListNode == targetLinkedListNode {
		return startLinkedListNode, nil
	} else if startLinkedListNode.Next == nil {
		return nil, errors.New("the linkedListNode is not exist")
	}
	return findByPtr(startLinkedListNode.Next, targetLinkedListNode)
}

func (linked *LinkedList) Print() {

}

func print(linkedListNode *LinkedListNode) *LinkedListNode {
	if linkedListNode.Next == nil {
		return nil
	}
	fmt.Printf("")
	return print(linkedListNode.Next)
}