package main

import "fmt"

type data interface{}

type LinkTableNode struct {
	value data           //存放数据
	next  *LinkTableNode //定义next结点
}

type LinkTable struct {
	head  *LinkTableNode //头节点
	tail  *LinkTableNode //尾节点
	count int            //节点总数
}

func CreateLinkTable() *LinkTable {
	pLinkTable := new(LinkTable)
	pLinkTable.head = nil
	pLinkTable.tail = nil
	pLinkTable.count = 0
	return pLinkTable
}

func IsEmpty(pLinkTable *LinkTable) bool { //直接根据count来判断是否为空
	return pLinkTable.count == 0
}

func DeleteLinkTable(pLinkTable *LinkTable) bool { //删除整个LinkTable
	if pLinkTable == nil {
		return false
	}
	for !IsEmpty(pLinkTable) { //从头结点开始一直删
		temp := pLinkTable.head
		pLinkTable.head = pLinkTable.head.next
		temp.next = nil
		temp.value = nil
		pLinkTable.count--
	}
	pLinkTable.head = nil
	pLinkTable.tail = nil
	return true
}

func InsertLinkTableNode(pLinkTableNode *LinkTableNode, pLinkTable *LinkTable) bool {
	if pLinkTable == nil || pLinkTableNode == nil {
		return false
	}
	if IsEmpty(pLinkTable) { //如果是空结点，那么直接把头结点和尾结点指向一个结点
		pLinkTable.head = pLinkTableNode
		pLinkTable.tail = pLinkTableNode
		pLinkTable.count++
	} else { //否则就向尾结点插入
		pLinkTable.tail.next = pLinkTableNode
		pLinkTable.count++
		pLinkTable.tail = pLinkTableNode
	}
	return true
}

func DeleteLinkTableNode(pLinkTableNode *LinkTableNode, pLinkTable *LinkTable) bool {
	//如果链表为空或者传入删除链表是空结点，返回false
	if pLinkTable == nil || pLinkTableNode == nil {
		return false
	}
	//如果删除的就是头结点
	if pLinkTableNode == pLinkTable.head {
		if pLinkTable.count == 1 { //如果链表只剩一个结点，head和tail都是一个结点
			pLinkTable.head = nil
			pLinkTable.tail = nil
		} else { //如果链表不止一个结点，就把头结点head指向下一个结点
			pLinkTable.head = pLinkTable.head.next
		}
		pLinkTable.count--
		return true
	}

	p := pLinkTable.head

	//如果删除结点是尾结点
	if p.next == pLinkTableNode {
		pLinkTable.tail = p
		pLinkTable.count--
		return true
	}

	//否则删除的结点就是中间结点
	for p.next != pLinkTable.tail {
		if p.next == pLinkTableNode {
			p = p.next.next
			pLinkTable.count--
			return true
		}
		p = p.next
	}

	return false
}

func SearchLinkTableNode(pLinkTableNode *LinkTableNode, pLinkTable *LinkTable) bool {
	if pLinkTable == nil || pLinkTableNode == nil {
		return false
	}
	p := pLinkTable.head //从头结点开始寻找
	for p != nil {
		if p == pLinkTableNode {
			return true
		}
		p = p.next
	}
	return false
}

func main() {
	pLinkTable := CreateLinkTable()
	pLinkTableNode1 := new(LinkTableNode)
	pLinkTableNode1.value = "example"
	InsertLinkTableNode(pLinkTableNode1, pLinkTable)

	pLinkTableNode2 := new(LinkTableNode)
	pLinkTableNode2.value = 2
	InsertLinkTableNode(pLinkTableNode2, pLinkTable)

	fmt.Println("count:", pLinkTable.count)
	fmt.Println("Head:", pLinkTable.head.value)
	fmt.Println("Tail:", pLinkTable.tail.value)
	fmt.Println("--------------------------------")
	DeleteLinkTableNode(pLinkTableNode2, pLinkTable)
	fmt.Println("count:", pLinkTable.count)
	fmt.Println("Head:", pLinkTable.head.value)
	fmt.Println("Tail:", pLinkTable.tail.value)
	fmt.Println("--------------------------------")
	fmt.Println("pLinkTableNode1 is in pLinkTable? ", SearchLinkTableNode(pLinkTableNode1, pLinkTable))
	fmt.Println("--------------------------------")
	DeleteLinkTable(pLinkTable)
	fmt.Println("count:", pLinkTable.count)
	// fmt.Println("Head:", pLinkTable.head.value)
	// fmt.Println("Tail:", pLinkTable.tail.value)
}
