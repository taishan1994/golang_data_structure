package array

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "strconv"
)

func OriginArr() [11][11]int {
	//创建原始数组
	var chessMap [11][11]int

	chessMap[1][2] = 1
	chessMap[2][3] = 2
	return chessMap
}

func PrintArr(chessMap [11][11]int) {
	//打印数组
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

//存储行、列、值
type valNode struct {
	row int
	col int
	//这里是接口类型，所以值可以是任意类型，不仅仅包含整型和字符型
	value interface{}
}

var sparseArr []valNode

func DoParseArr(chessMap [11][11]int) []valNode {
	//稀疏数组
	//遍历数组，如果某个值不为零，则将其放置在对应的结构体中
	val := valNode{
		//原来数组的行和列以及值
		row:   11,
		col:   11,
		value: 0,
	}
	//初始化存储稀疏数组
	sparseArr = append(sparseArr, val)

	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				var val valNode
				val.row = i
				val.col = j
				val.value = v2
				sparseArr = append(sparseArr, val)
			}
		}
	}

	for _, j := range sparseArr {
		fmt.Printf("第%d行，第%d列的值是%d\n", j.row, j.col, j.value.(int))

	}
	return sparseArr
}

func WriteParseArr(sparseArr []valNode, filepath string) {
	//将稀疏数组存储
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	for _, j := range sparseArr {
		//因为读取到的整型，需要转为字符串再进行写入
		//将接口赋值给一个变量需要进行类型断言
		str := strconv.Itoa(j.row) + " " + strconv.Itoa(j.col) + " " + strconv.Itoa((j.value.(int))) + "\n"
		wriiter := bufio.NewWriter(file)
		wriiter.WriteString(str)
		wriiter.Flush()
		// fmt.Printf("第%d行，第%d列的值是%d\n", j.row, j.col, j.value.(int))

	}
}

func ReadParseArr(filepath string) (newChessMap [11][11]int) {
	//初始化数组
	//读取存储的文件,并将每行转成
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		tmp := strings.Split(strings.TrimRight(str, "\n"), " ")
		// fmt.Println(strings.Split(strings.TrimRight(str, "\n"), " "))
		r, _ := strconv.Atoi(tmp[0])
		c, _ := strconv.Atoi(tmp[1])
		v, _ := strconv.Atoi(tmp[2])
		if r == 11 {
			continue
		}
		newChessMap[r][c] = v
		if err == io.EOF {
			break
		}
	}
	return newChessMap
}
