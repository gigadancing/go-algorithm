package trie

// 映射结构
type Pair struct {
	K int64
	V int64
}

type KeyWordKV map[int64]string // map

type CharBeginKV map[string][]*KeyWordTreeNode // 字典树结构

type PairList []Pair

// 实现排序接口
func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].V > p[j].V
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
