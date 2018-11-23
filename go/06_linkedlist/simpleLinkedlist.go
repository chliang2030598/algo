package _6_linkedlist

import (
	"errors"
	"fmt"
)

// LinkedLister  LinkedList的接口定义
type LinkedLister interface {
	InsertAfterNode(data string, nodevalue string) error
	InsertBeforeNode(data string, nodevalue string) error
	Find(data string) (*LinkedListNode, error)
	Delete(data string) error
	Print()
}

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
			break
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

		linked.Length++
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

// InsertBeforeNode value表示要插入的值，nodevalue表示已经存在节点的值(如果值有重复，默认第一个节点)，方法表示在nodevalue前插入vaule的值
func (linked *LinkedList) InsertBeforeNode(data string, nodevalue string) error {

	if linked.isFull() {
		return errors.New("the linkedList is full")
	}

	curLinkedListNode, err := linked.Find(nodevalue)
	if err != nil {
		return nil
	}
	//判断是否为头部节点,如果在头部节点前插入，需要重置头部节点
	if curLinkedListNode == linked.Head {
		newLinkedListNode := new(LinkedListNode)
		newLinkedListNode.Data = data
		newLinkedListNode.Next = curLinkedListNode

		linked.Head = newLinkedListNode

		linked.Length++
		return nil
	}

	//其他节点前插入
	preLinkedListNode, _ := FindByPreLinkedListNode(linked.Head, curLinkedListNode)

	newLinkedListNode := new(LinkedListNode)
	newLinkedListNode.Data = data
	newLinkedListNode.Next = curLinkedListNode

	preLinkedListNode.Next = newLinkedListNode
	linked.Length++
	return nil
}

// Delete 删除单链表节点：分为几种情况：1 删除头部节点 2 删除一般节点 3 删除尾部节点
func (linked *LinkedList) Delete(data string) error {
	deleteLinkedListNode, err := linked.Find(data)
	if err != nil {
		return err
	}

	// 判断删除的是否是头节点
	if deleteLinkedListNode == linked.Head {
		secondLinkedListNode := linked.Head.Next
		//设置新的头部节点
		linked.Head = secondLinkedListNode

		linked.Length--
		return nil
	} else if deleteLinkedListNode.Next == nil {
		// 查找尾部节点的上一个节点
		preLinkedListNode, _ := FindByPreLinkedListNode(linked.Head, deleteLinkedListNode)
		fmt.Printf("%v\n", preLinkedListNode)
		// 重新设置尾部节点
		preLinkedListNode.Next = nil
		linked.Length--
		return nil
	}

	// 删除一般节点
	preLinkedListNode, _ := FindByPreLinkedListNode(linked.Head, deleteLinkedListNode)

	// 获取删除节点的下一个节点
	afterLinkedListNode := deleteLinkedListNode.Next

	// 建立新的连接
	preLinkedListNode.Next = afterLinkedListNode

	linked.Length--

	// 疑问：要不要将deletenode的next置空
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

// FindByPreLinkedListNode 	从startLinkedListNode查找targetLinkedListNode的上一个节点
func FindByPreLinkedListNode(startLinkedListNode, targetLinkedListNode *LinkedListNode) (*LinkedListNode, error) {
	if startLinkedListNode.Next == targetLinkedListNode {
		return startLinkedListNode, nil
	} else if startLinkedListNode.Next == nil {
		return nil, errors.New("the linkedListNode is not exist")
	}
	return FindByPreLinkedListNode(startLinkedListNode.Next, targetLinkedListNode)
}

// Print 打印整个链表
func (linked *LinkedList) Print() {
	print(linked.Head)
}

// GetTailLinkedListNode 获取尾部节点
func (linked *LinkedList) GetTailLinkedListNode() *LinkedListNode {
	return getTail(linked.Head)
}

func getTail(linkedListNode *LinkedListNode) *LinkedListNode {
	if linkedListNode.Next == nil {
		return linkedListNode
	}
	return getTail(linkedListNode.Next)
}

func print(linkedListNode *LinkedListNode) *LinkedListNode {
	if linkedListNode.Next == nil {
		fmt.Printf("%p,%v\n", linkedListNode, linkedListNode)
		return nil
	}
	fmt.Printf("%p,%v\t", linkedListNode, linkedListNode)
	return print(linkedListNode.Next)
}

// LRU 单链表实现LRU
func (linked *LinkedList) LRU(data string) {
	existNode, _ := linked.Find(data)
	// 节点存在
	if existNode != nil {
		// 删除该节点
		linked.Delete(existNode.Data)

		// 将该节点插入到头
		linked.InsertBeforeNode(data, linked.Head.Data)

		return
	}

	// 节点不存在，判断容量是否已满，如果满，需要删除尾部节点，再插入到头部；不满直接插入到头部
	if linked.isFull() {
		tailNode := linked.GetTailLinkedListNode()
		linked.Delete(tailNode.Data)

		linked.InsertBeforeNode(data, linked.Head.Data)

		return
	}
	linked.InsertBeforeNode(data, linked.Head.Data)
	return
}
