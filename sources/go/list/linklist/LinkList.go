/*
 *  单链表：这里有很多博客直接使用node节点来表示链表，笔者不采用，而是认为节点是节点，链表是链表
 */
package linklist

import (
	"errors"
	"fmt"
)

// 结点结构体
type node struct {
	data	interface{}
	next	*node
}

// 单链表结构体：笔者这里将该结构体指针作为了 头指针 的概念来使用
type LinkList struct {
	length	int 		// 元素个数
	head	*node		// 首元结点
}

// 构造表
func NewLinkList() *LinkList {
	return &LinkList{
		length: 0,
		head:  nil,
	}
}

// 获取长度
func (l *LinkList) Length() int {
	return l.length
}

// 显示表
func (l *LinkList) Display() {

	if l.length == 0 {
		fmt.Println("数据结构为空")
		return
	}

	fmt.Printf("数据元素显示：")
	currentnode := l.head
	for currentnode.next != nil {
		fmt.Printf("%v ", currentnode.data)
		currentnode = currentnode.next	
	}
	fmt.Printf("%v ", currentnode.data)
	fmt.Println("")

}

func (l *LinkList) Append(e interface{}) {

	// 构造要插入的节点
	insertnode := &node{
		data: e,
		next: nil,
	}

	// 当前循环到的节点
	currentnode := l.head

	// 第一次追加
	if currentnode == nil {
		l.head = insertnode
		l.length++
		return
	}

	// 常见追加情况
	for currentnode.next != nil {
		currentnode = currentnode.next
	}
	currentnode.next = insertnode
	l.length++
}

// 插入元素
func (l *LinkList) Insert(index int, e interface{}) error {

	if index < 1 || index > l.length + 1 {
		fmt.Println("插入位序不正确")
		return errors.New("插入位序不正确")
	}

	// 构造要插入的节点
	insertnode := &node{
		data: e,
		next: nil,
	}

	// 当前循环到的节点
	currentnode := l.head

	// 如果是在第一个节点插入
	if index == 1 {
		l.head = insertnode
		l.head.next = currentnode
		l.length++
		return nil
	}

	// 常见插入：找到插入位置的前一个节点
	i := 1
	for currentnode.next != nil {
		if i == index - 1 {
			break
		}
		i++
		currentnode = currentnode.next
	}

	// 执行插入
	insertnode.next = currentnode.next
	currentnode.next = insertnode
	l.length++
	return nil

}

// 删除 按照位序删除
func (l *LinkList) Delete(index int) error {

	if index < 1 || index > l.length {
		fmt.Println("删除位序非法")
		return errors.New("删除位序非法")
	}

	if l.length == 0 {
		fmt.Println("数据结构为空")
		return errors.New("数据结构为空")
	}

	// 如果删除的是第一个元素
	if index == 1  {
		if l.length == 1 {
			l.head = nil
		} else {
			l.head = l.head.next
		}
		l.length--
		return nil
	}

	// 常规删除：找到要删除元素的前一个元素
	i := 1
	currentnode := l.head
	for currentnode.next != nil {
		if i == index-1 {
			break
		}
		i++
		currentnode = currentnode.next
	}
	currentnode.next = currentnode.next.next
	l.length--
	return nil

}

// 修改 按照位序修改
func (l *LinkList) Update(index int, e interface{}) error {

	if index < 1 || index > l.length {
		fmt.Println("位序非法")
		return errors.New("位序非法")
	}

	i := 1
	currentnode := l.head
	for currentnode.next != nil {
		if i == index {
			break
		}
		i++
		currentnode = currentnode.next
	}

	currentnode.data = e
	return nil
}

// 查询 按照位序查询值
func (l *LinkList) GetElem(index int) (interface{}, error) {

	if index < 1 || index > l.length {
		fmt.Println("位序非法")
		return nil, errors.New("位序不合法")
	}

	i := 1
	currentnode := l.head
	for currentnode.next != nil {
		if i == index {
			break
		}
		i++
		currentnode = currentnode.next
	}
	return currentnode.data, nil
}

// 查询 按照值查询位序
func (l *LinkList) Locate(e interface{}) (int, error) {

	i := 1
	currentnode := l.head
	for currentnode.next != nil {
		if currentnode.data == e {
			break
		}
		i++
		currentnode = currentnode.next
	}

	if i == l.length && currentnode.data != e {
		fmt.Println("未找到元素")
		return 0, errors.New("未找到元素")
	}

	return i, nil
}

// 查询前驱
func (l *LinkList) PrevElem(e interface{}) (pe interface{}, err error) {

	if l.length <= 1 {
		fmt.Println("数据结构为空")
		err = errors.New("数据结构为空")
		return
	}

	if l.head.data == e {
		fmt.Println("首元素无前驱")
		err = errors.New("首元素无前驱")
		return
	}

	i := 1
	currentnode := l.head
	for currentnode.next != nil {
		if currentnode.next.data == e {
			pe = currentnode.data
			return
		}
		i++
		currentnode = currentnode.next
	}

	fmt.Println("元素未找到")
	err = errors.New("元素未找到")
	return
}

// 查询后继
func (l *LinkList) NextElem(e interface{}) (ne interface{}, err error) {

	if l.length <= 1 {
		fmt.Println("数据结构为空")
		err = errors.New("数据结构为空")
		return
	}

	i := 1
	currentnode := l.head
	for currentnode.next != nil {
		if currentnode.data == e {
			break
		}
		i++
		currentnode = currentnode.next
	}

	if i == l.length && currentnode.data != e {
		fmt.Println("元素未找到")
		err = errors.New("元素未找到")
		return
	}

	if i == l.length && currentnode.data == e {
		fmt.Println("最后元素无后继")
		err = errors.New("最后元素无后继")
		return
	}

	ne = currentnode.next.data
	return
}

// 清空
func (l *LinkList) Clear() {
	l.head = nil
	l.length = 0
}
