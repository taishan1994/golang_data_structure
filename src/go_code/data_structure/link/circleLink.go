package link

import (
	"fmt"
)

//CatNode 猫
type CatNode struct {
	ID   int
	Name string
	next *CatNode
}

//InserCatNode 插入
func InserCatNode(head *CatNode, newCatNode *CatNode) {
	//初始化链表
	//头结点一开始是空的，当插入第一条数据时，进行初始化
	if head.next == nil {
		head.ID = newCatNode.ID
		head.Name = newCatNode.Name
		head.next = head
		return
	}
	//定义一个临时变量，找到环形的末尾，方便以后进行插入
	tmp := head
	for {
		if tmp.next == head {
			tmp.next = newCatNode
			newCatNode.next = head
			break
		} else {
			tmp = tmp.next
		}
	}
}

//DeleteCatNode 删除
func DeleteCatNode(head *CatNode, id int) *CatNode {
	//建立一个节点指向头结点
	tmp := head
	//如果头结点.next为空，说明是空链表
	if tmp.next == nil {
		fmt.Println("空链表")
		return head
	}
	//如果头结点.next就是它自己，说明只有一个元素
	if tmp.next == head {
		//判断该元素是否是要删除的，如果是，则将头结点置为空
		if tmp.ID != id {
			fmt.Println("要删除的id不存在")
		} else {
			head.next = nil
			return head
		}
	}
	//奖励一个辅助指针指向头结点
	helper := head
	//如果头结点正好是我们要删除的
	if tmp.ID == id {
		//如果头结点.next不是指向它自己，说明除了头结点之外还存在其它节点
		if tmp.next != head {
			//此时若想删除头结点，我们必须获得一个新的头结点
			tmp = head.next
			//将helper遍历到头结点的前一位
			for {
				if helper.next != head {
					helper = helper.next
				} else {
					//同时删除掉原来的头结点
					helper.next = head.next
					return tmp
				}
			}
		} else {
			//说明只有一个头结点，且是我们要删除的，直接将其置为空
			tmp.next = nil
		}
		//如果头结点不是我们要删除的
	} else {
		for {
			//如果找到一个节点是我们要删除的
			if tmp.next.ID == id {
				//删除该节点
				tmp2 := tmp.next
				tmp.next = tmp2.next
				break
				//如果找不到则继续遍历下一个节点
			} else {
				tmp = tmp.next
				//如果下一个节点是头结点，则表明完成遍历，找不到要删除的节点，并退出
				if tmp.next == head {
					fmt.Println("未找到该条记录")
					break
				}
			}
		}
	}

	return head
}

//ListCatNode 显示
func ListCatNode(head *CatNode) {
	tmp := head
	if tmp.next == nil {
		fmt.Println("空环形链表")
		return
	}
	for {
		fmt.Printf("猫的信息为：id=%d,name=%s\n", tmp.ID, tmp.Name)
		if tmp.next == head {
			break
		} else {
			tmp = tmp.next
		}

	}

}
