package associate

import (
	"strings"
	"unicode/utf8"
)

//Trie根节点
type Trie struct {
	root *Node //根节点,root字段中的node包含所有属于该根结点数据
}

//Node节点
type Node struct {
	char   rune           //字符
	childs map[rune]*Node //所有的子节点用map来存
	Data   interface{}    //自定义数据
	deep   int            //深度
	isTerm bool           //是否是一个字符串的结尾(完整的字符串)
}

//构造一个Trie树
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(' ', 1),
	}
}

//创造一个节点,传入字符,深度
func NewNode(char rune, deep int) *Node {
	return &Node{
		char:   char,                     //当前节点的字符
		childs: make(map[rune]*Node, 16), //保存子节点的map
		deep:   deep,                     //深度
	}
}

func (t *Trie) Add(key string, data interface{}) {

	var parent *Node = t.root
	allChars := []rune(key)

	for _, char := range allChars {
		node, ok := parent.childs[char]
		if !ok {
			node = NewNode(char, parent.deep+1)
			parent.childs[char] = node
		}
		parent = node
	}
	parent.Data = data
	parent.isTerm = true
}

//从Trie中查找,前缀搜索
func (t *Trie) prefixSearch(key string, limit int) (nodes []*Node) {
	var (
		node  = t.root
		queue []*Node //队列
	)
	allChars := []rune(key)
	for _, char := range allChars {
		child, ok := node.childs[char]
		if !ok {
			return
		}
		node = child
	}
	queue = append(queue, node)
	for len(queue) > 0 {
		var q2 []*Node
		for _, n := range queue {
			if n.isTerm == true {
				if len(nodes) >= limit {
					return
				} else {
					nodes = append(nodes, n)
				}
			}
			for _, v := range n.childs {
				q2 = append(q2, v)
			}
		}
		queue = q2
	}
	return
}

func (t *Trie) PrefixSearch(key string, limit int) (data []string) {
	nodes := t.prefixSearch(key, limit)
	for _, v := range nodes {
		data = append(data, (*v).Data.(string))
	}
	return
}

//敏感词过滤(字符串检索),查看该字符串中
func (t *Trie) Search(key string) bool {
	var node = t.root
	allChars := []rune(key)
	for _, char := range allChars {
		child, ok := node.childs[char]
		if !ok {
			return false
		}
		node = child
	}
	if node.isTerm == true {
		return true
	} else {
		return false
	}
}

//敏感词替换
func (t *Trie) Replace(text, replace string) (result string, hit bool) {

	chars := []rune(text)
	if t.root == nil {
		return
	}

	var left []rune
	node := t.root
	start := 0
	for index, v := range chars {
		ret, ok := node.childs[v]
		if !ok {
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = t.root
			continue
		}
		node = ret
		if ret.isTerm {
			hit = true
			node = t.root
			if ret.Data == nil {
				left = append(left, ([]rune(replace))...)
			} else {
				count := utf8.RuneCountInString(ret.Data.(string))
				str := strings.Repeat(replace, count)
				left = append(left, ([]rune(str))...)
			}
			start = index + 1
			continue
		}
	}
	result = string(left)
	return
}
