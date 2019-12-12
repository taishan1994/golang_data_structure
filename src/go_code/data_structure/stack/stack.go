package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int     //栈最大可以存放的数量
	Top    int     //栈顶
	Arr    [20]int //模拟栈
}

func (s *Stack) Push(val int) (err error) {
	//先判断栈是否满了
	if s.Top == s.MaxTop-1 {
		fmt.Println("栈满了")
		return errors.New("栈满了")
	}
	s.Top++
	s.Arr[s.Top] = val
	return
}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		fmt.Println("栈已空")
		return -1, errors.New("栈已空")
	}
	val = s.Arr[s.Top]
	s.Arr[s.Top] = 0
	s.Top--
	return val, nil

}

func (s *Stack) Show() {
	if s.Top == -1 {
		fmt.Println("栈为空")
		return
	}
	tmp := s
	for i := tmp.Top; i >= 0; i-- {
		fmt.Printf("Arr[%d]=%v\n", i, tmp.Arr[i])
	}
	return
}

//判断一个字符是数字还是加减乘除
func (s *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	//乘法
	case 42:
		res = num2 * num1
	//加法
	case 43:
		res = num2 + num1
	//减法
	case 45:
		res = num2 - num1
	//除法
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

//定义优先级
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}
