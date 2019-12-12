package link

import (
	"fmt"
)

//HerosNode 链表节点
type HerosNode struct {
	ID   int
	Name string
	pre  *HerosNode //指针
	next *HerosNode //指针
}

//InsertHerosNode 插入
func InsertHerosNode(head *HerosNode, newHerosNode *HerosNode) {
	tmp := head
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = newHerosNode
	newHerosNode.pre = tmp
}

//InsertHerosNodeByID 根据id从小到大插入
func InsertHerosNodeByID(head *HerosNode, newHerosNode *HerosNode) {
	tmp := head

	for {
		if tmp.next == nil {
			tmp.next = newHerosNode
			newHerosNode.pre = tmp
			break
		}
		if tmp.next.ID > newHerosNode.ID {
			tmp.next.pre = newHerosNode
			newHerosNode.pre = tmp
			tmp2 := tmp.next
			tmp.next = newHerosNode
			newHerosNode.next = tmp2
			break
		} else if tmp.next.ID == newHerosNode.ID {
			fmt.Printf("已经存在id为%d的节点\n", tmp.next.ID)
			break
		} else {
			tmp = tmp.next
		}
	}

}

//DeleteHerosNode 删除
func DeleteHerosNode(head *HerosNode, ID int) {
	tmp := head
	for {
		if tmp.next == nil {
			fmt.Println("链表中没有该id")
			break
		}
		if tmp.next.ID == ID {
			if tmp.next.next == nil {
				tmp.next = nil
			} else {
				tmp2 := tmp.next.next
				tmp.next = tmp2
				tmp2.pre = tmp
			}
			break
		} else {
			tmp = tmp.next
		}
	}
}

//FindHerosNode 查找
func FindHerosNode(head *HerosNode, ID int) {
	tmp := head
	for {
		if tmp.next == nil {
			fmt.Println("链表中没有该id")
			break
		}
		if tmp.next.ID == ID {
			fmt.Println("找到了该id")
			break
		} else {
			tmp = tmp.next
		}
	}
}

//ModifyHerosNode 修改
func ModifyHerosNode(head *HerosNode, ID int, changeName string) {
	tmp := head
	for {
		if tmp.next == nil {
			fmt.Println("链表中没有该id")
			break
		}
		if tmp.next.ID == ID {
			tmp.next.Name = changeName
			break
		} else {
			tmp = tmp.next
		}
	}
}

//ForListHerosNode 显示信息
func ForListHerosNode(forHead *HerosNode) {
	fmt.Println("正向打印所有信息")
	tmp := forHead
	if tmp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("节点信息如下：id=%d,name=%s\n", tmp.next.ID, tmp.next.Name)
		tmp = tmp.next
		if tmp.next == nil {
			fmt.Println("已显示所有信息")
			break
		}
	}
}

//BackListHerosNode 显示信息
func BackListHerosNode(head *HerosNode) {
	fmt.Println("----------------------")
	fmt.Println("反向打印所有信息")
	tmp := head
	if tmp.next == nil {
		fmt.Println("链表为空")
		return
	}
	var backHead *HerosNode
	for {
		tmp = tmp.next
		if tmp.next == nil {
			backHead = tmp
			break
		}
	}
	for {
		fmt.Printf("节点信息如下：id=%d,name=%s\n", backHead.ID, backHead.Name)
		backHead = backHead.pre
		if backHead.pre == head {
			fmt.Printf("节点信息如下：id=%d,name=%s\n", backHead.ID, backHead.Name)
			fmt.Println("已显示所有信息")
			break
		}
	}
}
