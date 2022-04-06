package main

//定义结点
type LinkTableNode struct {
	next *LinkTableNode
}

//定义链表
type LinkTable struct {
	len  int
	head *LinkTableNode
	tail *LinkTableNode
}

//创建链表
func CreateLinkTable() *LinkTable {
	var linktable *LinkTable = new(LinkTable)
	if linktable == nil {
		return nil
	}
	linktable.len = 0
	linktable.head = nil
	linktable.tail = nil
	return linktable
}

//销毁链表
func DestroyLinkTable(linktable *LinkTable) bool {
	if linktable == nil {
		return false
	}
	for linktable != nil {
		var temp *LinkTableNode = linktable.head
		linktable.head = linktable.head.next
		temp.next = nil
	}
	linktable.head = nil
	linktable.tail = nil
	linktable.len = 0
	return true
}

//添加节点到表尾
func AddLinkTableNode(linktable *LinkTable, node *LinkTableNode) bool {
	if linktable == nil || node == nil {
		return false
	}

	if linktable.len == 0 { //链表为空时
		linktable.head = node
		linktable.tail = node
		linktable.len = 1
	} else { //链表不为空时
		linktable.tail.next = node
		linktable.tail = node
		linktable.len++
	}
	return true
}

//获取链表头节点
func GetHeadNode(linktable *LinkTable) *LinkTableNode {
	if linktable == nil {
		return nil
	}
	return linktable.head
}

//获取链表长度
func (linktable LinkTable) GetLength() int {
	return linktable.len
}

//获取index位置的结点
func (linktable LinkTable) GetLinkTableNode(index int) *LinkTableNode {
	if linktable.len < index {
		return nil
	}
	temp := linktable.head
	for index > 1 {
		temp = temp.next
	}
	return temp.next
}

//查找节点是否存在
func FindLinkTableNode(linktable *LinkTable, node *LinkTableNode) bool {
	if linktable == nil || node == nil {
		return false
	}

	var temp *LinkTableNode = linktable.head
	for temp != nil {
		if temp == node {
			return true
		}
		temp = temp.next
	}
	return false
}

//删除指定节点
func DeleteLinkTableNode(linktable *LinkTable, node *LinkTableNode) bool {
	if linktable == nil || node == nil {
		return false
	}

	if linktable.head == node {
		linktable.head = linktable.head.next
		linktable.len--
		if linktable.len == 0 {
			linktable.tail = nil
		}
		return true
	}

	var temp *LinkTableNode = linktable.head
	for temp != nil {
		if temp.next == node {
			temp.next = temp.next.next
			linktable.len--
			if linktable.len == 0 {
				linktable.tail = nil
			}
			return true
		}
		temp = temp.next
	}

	return false
}
