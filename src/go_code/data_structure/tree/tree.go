package tree

import (
	"fmt"
)

type TreeNode struct {
	ID    int
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrder(root.Right)
	}
}

func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}
