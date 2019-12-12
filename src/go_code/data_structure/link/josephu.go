package link

import (
	"fmt"
)

//Kid 小孩
type Kid struct {
	ID   int
	next *Kid
}

//AddKid 添加一个
func AddKid(num int) *Kid {
	first := &Kid{}
	cur := &Kid{}

	if num < 1 {
		fmt.Println("不合法")
		return first
	}
	for i := 1; i <= num; i++ {
		kid := &Kid{
			ID: 1,
		}
		if i == 1 {
			first = kid
			cur = kid
			cur.next = first
		} else {
			tmpKid := &Kid{
				ID: i,
			}
			cur.next = tmpKid
			cur = tmpKid
			cur.next = first

		}
	}
	return first
}

//ShowKid 显示列表
func ShowKid(first *Kid) {
	if first.next == nil {
		fmt.Println("链表已空")
	}
	cur := first
	for {
		fmt.Printf("小孩编号：%d\n", cur.ID)
		if cur.next == first {
			break
		}
		cur = cur.next
	}
}

//Play 开始数数
func Play(first *Kid, start int, count int) {

	if first.next == nil {
		fmt.Println("空链表")
		return
	}
	tail := first
	for {
		//到最后一个节点了
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	//删除就以frst为主，让first移动到要删除的位置
	for i := 1; i <= start-1; i++ {
		first = first.next
		tail = tail.next
	}
	//开始数，然后进行删除
	for {
		for i := 1; i <= count-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("编号为：%d  的出列\n", first.ID)
		first = first.next
		tail.next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("编号为：%d  的出列\n", first.ID)
}
