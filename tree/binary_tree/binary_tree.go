package binary_tree

import (
	"bytes"
	"fmt"
	"strconv"
)

// 节点
type Node struct {
	Data  int   // 数据
	Left  *Node // 执行左子树
	Right *Node // 指向右子树
}

// 二叉树
type BinaryTree struct {
	Root *Node // 根节点
	Size int   // 节点数量
}

// 创建二叉树
func NewBinaryTree() *BinaryTree {
	bst := &BinaryTree{
		Size: 0,
		Root: nil,
	}
	return bst
}

// 获得二叉树大小
func (bst *BinaryTree) GetSize() int {
	return bst.Size
}

// 二叉树是否为空
func (bst *BinaryTree) IsEmpty() bool {
	return bst.Size == 0
}

// 根节点插入
func (bst *BinaryTree) Add(data int) {
	bst.Root = bst.add(bst.Root, data)
}

// 从某个节点开始插入
// 参数：
// 		node - 开始的节点
// 		data - 数据
// 返回值：
// 		新插入的节点
func (bst *BinaryTree) add(node *Node, data int) *Node {
	if node == nil {
		bst.Size++
		return &Node{Data: data}
	} else {
		if data < node.Data {
			node.Left = bst.add(node.Left, data)
		} else {
			node.Right = bst.add(node.Right, data)
		}
		return node
	}
}

// 判断数据是否存在
func (bst *BinaryTree) IsIn(data int) bool {
	return bst.isIn(bst.Root, data)
}

// 递归查找
func (bst *BinaryTree) isIn(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if data == node.Data {
		return true
	} else if data < node.Data {
		return bst.isIn(node.Left, data)
	} else {
		return bst.isIn(node.Right, data)
	}
}

// 最大值
func (bst *BinaryTree) Max() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.max(bst.Root).Data
}

func (bst *BinaryTree) max(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return bst.max(node.Right)
}

// 最小值
func (bst *BinaryTree) Min() int {
	if bst.Size == 0 {
		panic("二叉树为空")
	}
	return bst.min(bst.Root).Data
}

func (bst *BinaryTree) min(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return bst.min(node.Left)
}

// 先序遍历
func (bst *BinaryTree) Preorder() {
	bst.preorder(bst.Root)
}

func (bst *BinaryTree) preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	bst.preorder(node.Left)
	bst.preorder(node.Right)
}

// 中序遍历
func (bst *BinaryTree) Inorder() {
	bst.inorder(bst.Root)
}

func (bst *BinaryTree) inorder(node *Node) {
	if node == nil {
		return
	}
	bst.inorder(node.Left)
	fmt.Printf("%d ", node.Data)
	bst.inorder(node.Right)
}

// 后序遍历
func (bst *BinaryTree) Postorder() {
	bst.postorder(bst.Root)
}

func (bst *BinaryTree) postorder(node *Node) {
	if node == nil {
		return
	}
	bst.postorder(node.Left)
	bst.postorder(node.Right)
	fmt.Printf("%d ", node.Data)
}

// 二叉树转string
func (bst *BinaryTree) String() string {
	var buffer bytes.Buffer
	bst.generateString(bst.Root, 0, &buffer)
	return buffer.String()
}

//
func (bst *BinaryTree) generateString(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		//buffer.WriteString(bst.generateDepthString(depth) + "nil\n") // 空节点
		return
	}

	bst.generateString(node.Left, depth+1, buffer)
	buffer.WriteString(bst.generateDepthString(depth) + strconv.Itoa(node.Data) + "\n")
	bst.generateString(node.Right, depth+1, buffer)
}

//
func (bst *BinaryTree) generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

// 删除最大节点
func (bst *BinaryTree) RemoveMax() int {
	max := bst.Max()
	bst.Root = bst.removeMax(bst.Root)
	return max
}

//
func (bst BinaryTree) removeMax(node *Node) *Node {
	if node.Right == nil {
		left := node.Left
		bst.Size--
		return left
	}
	node.Right = bst.removeMax(node.Right)
	return node
}

// 删除最小节点
func (bst *BinaryTree) RemoveMin() int {
	min := bst.Min()
	bst.Root = bst.removeMin(bst.Root)
	return min
}

//
func (bst *BinaryTree) removeMin(node *Node) *Node {
	if node.Left == nil {
		right := node.Right // 备份右子树
		bst.Size--

		return right
	}
	node.Left = bst.removeMin(node.Left)
	return node
}
