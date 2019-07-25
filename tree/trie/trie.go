package trie

import (
	"fmt"
	"sort"
	"sync"
)

// 节点
type KeyWordTreeNode struct {
	KeyWordIDs            map[int64]bool
	Char                  string
	ParentKeyWordTreeNode *KeyWordTreeNode            // 父节点
	SubKeyWordTreeNodes   map[string]*KeyWordTreeNode // 子节点集合
}

// 创建节点
func NewKeyWordTreeNode() *KeyWordTreeNode {
	return &KeyWordTreeNode{
		KeyWordIDs:            make(map[int64]bool, 0),
		Char:                  "",
		ParentKeyWordTreeNode: nil,
		SubKeyWordTreeNodes:   make(map[string]*KeyWordTreeNode, 0),
	}
}

// 创建节点
func NewKeyWordTreeNodeWithParams(ch string, parent *KeyWordTreeNode) *KeyWordTreeNode {
	return &KeyWordTreeNode{
		KeyWordIDs:            make(map[int64]bool, 0),
		Char:                  ch,
		ParentKeyWordTreeNode: parent,
		SubKeyWordTreeNodes:   make(map[string]*KeyWordTreeNode, 0),
	}
}

// 字典树
type KeyWordTree struct {
	root        *KeyWordTreeNode // 根节点
	kv          KeyWordKV        // 映射关系
	charBeginKV CharBeginKV      // 开始映射
	rw          *sync.RWMutex    // 读写锁
}

// 创建字典树
func NewKeyWordTree() *KeyWordTree {
	return &KeyWordTree{
		root:        NewKeyWordTreeNode(),
		kv:          KeyWordKV{},
		charBeginKV: CharBeginKV{},
		rw:          new(sync.RWMutex),
	}
}

//
func (tree *KeyWordTree) Debugout() {
	fmt.Println("st.kv", tree.kv)
	tempRoot := tree.root
	dfs(tempRoot)
}

// 遍历
func dfs(root *KeyWordTreeNode) {
	if root == nil {
		return
	}
	fmt.Println("s.root=", root.Char)
	fmt.Println("s.KeyWordIds=", root.KeyWordIDs)
	for _, v := range root.SubKeyWordTreeNodes {
		dfs(v)
	}
}

// 字符串压入树中
func (tree *KeyWordTree) Put(id int64, keyword string) {
	tree.rw.Lock()
	defer tree.rw.Unlock()

	tree.kv[id] = keyword       // 保存
	tempRoot := tree.root       // 备份root
	for _, v := range keyword { // 循环每一个字符
		ch := string(v) // 字符转成字符串
		if tempRoot.SubKeyWordTreeNodes[ch] == nil {
			node := NewKeyWordTreeNodeWithParams(ch, tempRoot) // 创建新节点插入
			tempRoot.SubKeyWordTreeNodes[ch] = node
			tree.charBeginKV[ch] = append(tree.charBeginKV[ch], node) //
		} else {
			node := tempRoot.SubKeyWordTreeNodes[ch]
			node.KeyWordIDs[id] = true                  // 生效
			tempRoot = tempRoot.SubKeyWordTreeNodes[ch] // 向前推进
		}
	}
}

// 搜索提示，limit限制深度
func (tree *KeyWordTree) Search(keyword string, limit int) []string {
	tree.rw.Lock()
	defer tree.rw.Unlock()

	ids := make(map[int64]int64, 0)
	for pos, v := range keyword {
		ch := string(v)
		begins := tree.charBeginKV[ch] // 取得映射字符的所有节点
		for _, begin := range begins { // 循环每一个节点
			keyWordTempPtr := begin // 备份地址
			nextPos := pos + 1      // 标记下一个位置
			for len(keyWordTempPtr.SubKeyWordTreeNodes) > 0 && nextPos < len(keyword) {
				nextCh := string(keyword[nextPos]) // 下一个字符
				if keyWordTempPtr.SubKeyWordTreeNodes[nextCh] == nil {
					break
				}
				keyWordTempPtr = keyWordTempPtr.SubKeyWordTreeNodes[nextCh] // 递推前进
				nextPos++
			}
			for id := range keyWordTempPtr.KeyWordIDs { // 保存结果
				ids[id] = ids[id] + 1
			}
		}
	}

	list := PairList{}
	for id, count := range ids {
		list = append(list, Pair{id, count})
	}

	if !sort.IsSorted(list) {
		sort.Sort(list) // 排序
	}

	if len(list) > limit {
		list = list[:limit]
	}

	ret := make([]string, 0)
	for _, item := range list {
		ret = append(ret, tree.kv[item.K])
	}

	return ret
}

//
func (tree *KeyWordTree) Suggestion(keyword string, limit int) []string {

	return nil
}
