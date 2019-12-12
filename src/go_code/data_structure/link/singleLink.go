package link

import (
	"fmt"
)

//HeroNode 链表节点
type HeroNode struct {
	ID   int
	Name string
	next *HeroNode //指针
}

//InsertHeroNode 插入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	tmp := head
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	tmp.next = newHeroNode
}

//InsertHeroNodeByID 根据id从小到大插入
func InsertHeroNodeByID(head *HeroNode, newHeroNode *HeroNode) {
	tmp := head

	for {
		if tmp.next == nil {
			tmp.next = newHeroNode
			break
		}
		if tmp.next.ID > newHeroNode.ID {
			tmp2 := tmp.next
			tmp.next = newHeroNode
			newHeroNode.next = tmp2
			break
		} else if tmp.next.ID == newHeroNode.ID {
			fmt.Printf("已经存在id为%d的节点\n", tmp.next.ID)
			break
		} else {
			tmp = tmp.next
		}
	}

}

//DeleteHeroNode 删除
func DeleteHeroNode(head *HeroNode, ID int) {
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
			}
			break
		} else {
			tmp = tmp.next
		}
	}
}

//FindHeroNode 查找
func FindHeroNode(head *HeroNode, ID int) {
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

//ModifyHeroNode 修改
func ModifyHeroNode(head *HeroNode, ID int, changeName string) {
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

//ListHeroNode 显示信息
func ListHeroNode(head *HeroNode) {
	tmp := head
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
