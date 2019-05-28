package binary_tree

import (
	"fmt"
	"testing"
)

func PrintTree(root *Node, layer int) {
	if root == nil {
		return
	}
	for i := 0; i < layer; i++ {
		fmt.Printf("--")
	}
	fmt.Printf("%d \n", root.Data)
	layer++
	PrintTree(root.Left, layer)
	PrintTree(root.Right, layer)
}

func TestBinaryTree(t *testing.T) {
	bst := NewBinaryTree()
	bst.Add(4)
	bst.Add(2)
	bst.Add(6)
	bst.Add(1)
	bst.Add(3)
	bst.Add(5)
	bst.Add(7)
	//fmt.Println("节点个数:", bst.Size)
	//PrintTree(bst.Root, 0)
	//fmt.Println(bst.String())
	//fmt.Println(bst.IsIn(100))
	//fmt.Println(bst.IsIn(7))
	//fmt.Println("最大值:", bst.Max())
	//fmt.Println("最小值:", bst.Min())
	//fmt.Println("前序遍历:")
	//bst.Preorder()
	//fmt.Println()
	//fmt.Println("中序遍历:")
	//bst.Inorder()
	//fmt.Println()
	//fmt.Println("后序遍历:")
	//bst.Postorder()
	//fmt.Println()
	//fmt.Println("删除最小值")
	//bst.RemoveMin()
	//fmt.Println(bst.String())
	//fmt.Println("删除最大值")
	//bst.RemoveMax()
	//fmt.Println(bst.String())
	bst.Remove(4)
	fmt.Println(bst.String())
}
