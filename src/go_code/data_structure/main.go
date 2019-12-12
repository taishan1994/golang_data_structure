package main

import (
	"fmt"
	"go_code/data_structure/tree"
)

func main() {

	node7 := &tree.TreeNode{
		ID:    7,
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node6 := &tree.TreeNode{
		ID:    6,
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node5 := &tree.TreeNode{
		ID:    5,
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &tree.TreeNode{
		ID:    4,
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node3 := &tree.TreeNode{
		ID:    3,
		Val:   3,
		Left:  node6,
		Right: node7,
	}
	node2 := &tree.TreeNode{
		ID:    2,
		Val:   2,
		Left:  node4,
		Right: node5,
	}

	node1 := &tree.TreeNode{
		ID:    1,
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	fmt.Println("先序遍历")
	tree.PreOrder(node1)
	fmt.Println()
	fmt.Println("中序遍历")
	tree.InOrder(node1)
	fmt.Println()
	fmt.Println("后序遍历")
	tree.PostOrder(node1)

	// key := ""
	// id := 0
	// name := ""
	// var hashTable hash.HashTable
	// for {
	// 	fmt.Println("==========员工菜单==========")
	// 	fmt.Println("insert 表示添加员工")
	// 	fmt.Println("show   表示显示员工")
	// 	fmt.Println("find   表示查询员工")
	// 	fmt.Println("exit   表示退出员工")
	// 	fmt.Println("请输入你的选择：")
	// 	fmt.Scanln(&key)
	// 	switch key {
	// 	case "insert":
	// 		fmt.Println("请输入员工id：")
	// 		fmt.Scanln(&id)
	// 		fmt.Println("请输入员工名字：")
	// 		fmt.Scanln(&name)
	// 		emp := &hash.Emp{
	// 			ID:   id,
	// 			Name: name,
	// 		}
	// 		hashTable.Insert(emp)
	// 	case "show":
	// 		hashTable.Show()
	// 	case "find":
	// 		fmt.Println("请输入你要查找的id：")
	// 		fmt.Scanln(&id)
	// 		emp := hashTable.Find(id)
	// 		if emp == nil {
	// 			fmt.Printf("id=%d的员工不存在\n", id)
	// 		} else {
	// 			//显示雇员信息
	// 			emp.ShowMe()

	// 		}
	// 	case "exit":
	// 		os.Exit(0)
	// 	}
	// }

	// maze.Maze()

	// str := "30+2*6-600/2"
	// n := len(str)
	// numStack := &stack.Stack{
	// 	MaxTop: n,
	// 	Top:    -1,
	// }
	// operStack := &stack.Stack{
	// 	MaxTop: n,
	// 	Top:    -1,
	// }
	// index := 0
	// num1 := 0
	// num2 := 0
	// oper := 0
	// result := 0
	// keepNum := ""
	// for {
	// 	ch := str[index : index+1]
	// 	//字符对应的ASCII码值
	// 	tmp := int([]byte(ch)[0])
	// 	if operStack.IsOper(tmp) {
	// 		if operStack.Top == -1 {
	// 			operStack.Push(tmp)
	// 		} else {
	// 			//判断栈顶的优先级
	// 			if operStack.Priority(operStack.Arr[operStack.Top]) >=
	// 				operStack.Priority(tmp) {
	// 				num1, _ = numStack.Pop()
	// 				num2, _ = numStack.Pop()
	// 				oper, _ = operStack.Pop()
	// 				result = operStack.Cal(num1, num2, oper)
	// 				//将计算的结果重新入栈
	// 				numStack.Push(result)
	// 				//当前符号重新入栈
	// 				operStack.Push(tmp)
	// 			} else {
	// 				operStack.Push(tmp)
	// 			}
	// 		}

	// 	} else {
	// 		//判断型如30等数字
	// 		keepNum += ch

	// 		if index == n-1 {
	// 			val, _ := strconv.ParseInt(keepNum, 10, 64)
	// 			numStack.Push(int(val))
	// 		} else {
	// 			//向index后面探以位
	// 			if operStack.IsOper(int([]byte(str[index+1 : index+2])[0])) {
	// 				val, _ := strconv.ParseInt(keepNum, 10, 64)
	// 				numStack.Push(int(val))
	// 				keepNum = ""
	// 			}
	// 		}
	// 		//入栈的数要是int型
	// 		// val, _ := strconv.ParseInt(ch, 10, 64)
	// 		// numStack.Push(int(val))
	// 	}
	// 	if index == n-1 {
	// 		break
	// 	}
	// 	//继续扫描
	// 	index++
	// }

	// //如果表达式扫描完毕，依次取出符号和两位数进行运算
	// for {
	// 	if operStack.Top == -1 {
	// 		break
	// 	}
	// 	num1, _ = numStack.Pop()
	// 	num2, _ = numStack.Pop()
	// 	oper, _ = operStack.Pop()
	// 	result = operStack.Cal(num1, num2, oper)
	// 	//将计算的结果重新入栈
	// 	numStack.Push(result)
	// }

	// res, _ := numStack.Pop()
	// fmt.Printf("计算表达式 %v 的值是：%v\n", str, res)

	// s := &stack.Stack{
	// 	MaxTop: 5,
	// 	Top:    -1, // 当栈顶为-1时，表示栈顶为空
	// }
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)
	// s.Push(4)
	// s.Push(5)
	// for i := 0; i < 3; i++ {
	// 	val, err := s.Pop()
	// 	if err != nil {
	// 		fmt.Println("出栈错误err=", err)
	// 		return
	// 	}
	// 	fmt.Println("出栈的值是：", val)
	// }
	// s.Push(6)
	// s.Show()

	// arr := [7]int{5, 2, 8, 1, 4, 9, 3}
	// fmt.Println("原始数组为：", arr)
	// sort.SelectSort(&arr)
	// sort.InsertSort(&arr)
	// sort.BubbleSort(&arr)
	// sort.QuickSort(0, len(arr)-1, &arr)
	// fmt.Println("排序后结果：", arr)

	// first := link.AddKid(20)
	// link.ShowKid(first)
	// link.Play(first, 1, 3)

	// head := &link.CatNode{}
	// cat1 := &link.CatNode{
	// 	ID:   1,
	// 	Name: "tom",
	// }
	// cat2 := &link.CatNode{
	// 	ID:   2,
	// 	Name: "jack",
	// }
	// cat3 := &link.CatNode{
	// 	ID:   3,
	// 	Name: "bob",
	// }
	// cat4 := &link.CatNode{
	// 	ID:   4,
	// 	Name: "mike",
	// }

	// link.InserCatNode(head, cat1)
	// link.InserCatNode(head, cat2)
	// link.InserCatNode(head, cat3)
	// link.InserCatNode(head, cat4)
	// link.ListCatNode(head)
	// fmt.Println("------------------------------")
	// fmt.Println("删除id=1后的结果是：")
	// h1 := link.DeleteCatNode(head, 1)
	// link.ListCatNode(h1)
	// fmt.Println("------------------------------")
	// fmt.Println("删除id=4后的结果是：")
	// h2 := link.DeleteCatNode(h1, 4)
	// link.ListCatNode(h2)
	// fmt.Println("------------------------------")
	// fmt.Println("删除id=3后的结果是：")
	// h3 := link.DeleteCatNode(h2, 3)
	// link.ListCatNode(h3)
	// fmt.Println("------------------------------")
	// fmt.Println("删除id=2后的结果是：")
	// h4 := link.DeleteCatNode(h3, 2)
	// link.ListCatNode(h4)

	// head := &link.HerosNode{}
	// hero1 := &link.HerosNode{
	// 	ID:   1,
	// 	Name: "宋江",
	// }
	// hero2 := &link.HerosNode{
	// 	ID:   2,
	// 	Name: "李逵",
	// }
	// hero4 := &link.HerosNode{
	// 	ID:   4,
	// 	Name: "林冲",
	// }
	// hero3 := &link.HerosNode{
	// 	ID:   3,
	// 	Name: "武松",
	// }
	// // link.InsertHerosNode(head, hero1)
	// // link.InsertHerosNode(head, hero2)
	// // link.InsertHerosNode(head, hero4)
	// // link.InsertHerosNode(head, hero3)

	// link.InsertHerosNodeByID(head, hero2)
	// link.InsertHerosNodeByID(head, hero1)
	// link.InsertHerosNodeByID(head, hero4)
	// link.InsertHerosNodeByID(head, hero3)
	// link.DeleteHerosNode(head, 3)
	// link.FindHerosNode(head, 4)
	// link.ModifyHerosNode(head, 4, "我是修改后的英雄")
	// link.ForListHerosNode(head)
	// link.BackListHerosNode(head)

	// head := &link.HeroNode{}
	// hero1 := &link.HeroNode{
	// 	ID:   1,
	// 	Name: "宋江",
	// }
	// hero2 := &link.HeroNode{
	// 	ID:   2,
	// 	Name: "李逵",
	// }
	// hero4 := &link.HeroNode{
	// 	ID:   4,
	// 	Name: "林冲",
	// }
	// hero3 := &link.HeroNode{
	// 	ID:   3,
	// 	Name: "武松",
	// }
	// // link.InsertHeroNode(head, hero1)
	// // link.InsertHeroNode(head, hero2)
	// // link.InsertHeroNode(head, hero4)
	// // link.InsertHeroNode(head, hero3)

	// link.InsertHeroNodeByID(head, hero2)
	// link.InsertHeroNodeByID(head, hero1)
	// link.InsertHeroNodeByID(head, hero4)
	// link.InsertHeroNodeByID(head, hero3)
	// link.DeleteHeroNode(head, 3)
	// link.FindHeroNode(head, 4)
	// link.ModifyHeroNode(head, 4, "我是修改后的英雄")
	// link.ListHeroNode(head)

	// var key string
	// var val int
	// q := &queue.CircleQueue{
	// 	MaxSize: 5,
	// 	Front:   0,
	// 	Rear:    0,
	// }
	// for {
	// 	fmt.Println("------------------------------")
	// 	fmt.Println("1.输入push表示添加数据到队列")
	// 	fmt.Println("2.输入pop表示从队列中获取数据")
	// 	fmt.Println("3.输入show表示显示队列")
	// 	fmt.Println("4.输入exit表示退出")
	// 	fmt.Println("------------------------------")
	// 	fmt.Scanln(&key)
	// 	switch key {
	// 	case "push":
	// 		fmt.Println("请输入要添加的值：")
	// 		fmt.Scanln(&val)
	// 		err := q.Push(val)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("添加成功")
	// 			fmt.Println("Rear：", q.Rear)
	// 		}
	// 	case "pop":
	// 		val, err := q.Pop()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("得到的值为：", val)
	// 			fmt.Println("Front：", q.Front)
	// 		}

	// 	case "show":
	// 		q.Show()
	// 		fmt.Println()
	// 	case "exit":
	// 		os.Exit(0)
	// 	}
	// }

	// var key string
	// var val int
	// q := &queue.Queue{
	// 	MaxSize: 4,
	// 	Front:   -1,
	// 	Rear:    -1,
	// }
	// for {
	// 	fmt.Println("------------------------------")
	// 	fmt.Println("1.输入add表示添加数据到队列")
	// 	fmt.Println("2.输入get表示从队列中获取数据")
	// 	fmt.Println("3.输入show表示显示队列")
	// 	fmt.Println("4.输入exit表示退出")
	// 	fmt.Println("------------------------------")
	// 	fmt.Scanln(&key)
	// 	switch key {
	// 	case "add":
	// 		fmt.Println("请输入要添加的值：")
	// 		fmt.Scanln(&val)
	// 		err := q.AddQueue(val)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("添加成功")
	// 			fmt.Println("Rear：", q.Rear)
	// 		}
	// 	case "get":
	// 		val, err := q.GetQueue()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		} else {
	// 			fmt.Println("得到的值为：", val)
	// 			fmt.Println("Rear：", q.Front)
	// 		}

	// 	case "show":
	// 		q.ShowQueue()
	// 		fmt.Println()
	// 	case "exit":
	// 		os.Exit(0)
	// 	}
	// }

	// chessMap := array.OriginArr()
	// array.PrintArr(chessMap)
	// sparseArr := array.DoParseArr(chessMap)
	// filepath := "array/data.txt"
	// array.WriteParseArr(sparseArr, filepath)
	// newChessMap := array.ReadParseArr(filepath)
	// array.PrintArr(newChessMap)

}
