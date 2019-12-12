package queue

import (
	"errors"
	"fmt"
)

//Queue 使用结构体管理队列
type Queue struct {
	MaxSize int
	Array   [5]int //模拟队列
	Front   int    //队列首位
	Rear    int    //队列尾部
}

//AddQueue 向队列中添加一个值
func (q *Queue) AddQueue(val int) (err error) {
	q.MaxSize = 4
	//先判断队列是否已满
	if q.Rear == q.MaxSize-1 {
		return errors.New("队列已满")
	}
	q.Rear++
	q.Array[q.Rear] = val
	return
}

//GetQueue 得到一个值
func (q *Queue) GetQueue() (val int, err error) {
	if q.Front == q.Rear {
		return -1, errors.New("队列已空")
	}
	q.Front++
	val = q.Array[q.Front]
	return val, err
}

//ShowQueue 显示队列
func (q *Queue) ShowQueue() {
	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("queue[%d]=%v\t", i, q.Array[i])
	}
}
