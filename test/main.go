package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	//rune表示一个utf8字符
	char   rune
	Data   interface{}
	parent *Node
	Depth  int
	//childs 用来当前节点的所有孩子节点
	childs map[rune]*Node
	term   bool
}

type Trie struct {
	root *Node
	size int
}

func NewNode() *Node {
	return &Node{
		childs: make(map[rune]*Node, 32),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

//假如我要把 敏感词： “我操”
// Add("我操", nil)
// Add("色情片", nil)
func (p *Trie) Add(key string, data interface{}) (err error) {

	key = strings.TrimSpace(key)
	node := p.root
	runes := []rune(key)
	for _, r := range runes {
		ret, ok := node.childs[r]
		if !ok {
			ret = NewNode()
			ret.Depth = node.Depth + 1
			ret.char = r
			node.childs[r] = ret
		}

		node = ret
	}

	node.term = true
	node.Data = data
	return
}

// text = "我们都喜欢王八蛋"
// replace = "***"
func (p *Trie) Check(text, replace string) (result string, hit bool) {

	chars := []rune(text) //将字符串打散成字符数组
	if p.root == nil {
		return
	}

	var left []rune
	node := p.root
	start := 0

	for index, v := range chars {
		ret, ok := node.childs[v] //看这个字符是否在childs中
		if !ok {
			//如果不在,
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = p.root
			continue
		}
		fmt.Printf("%v", ret)
		os.Exit(1)
		node = ret
		if ret.term {
			hit = true
			node = p.root
			left = append(left, ([]rune(replace))...)
			start = index + 1
			continue
		}
	}

	result = string(left)
	return
}

func main() {

	trie := NewTrie()
	//将敏感词加入到Trie中
	trie.Add("黄色", nil)
	trie.Add("绿色", nil)
	trie.Add("蓝色", nil)

	result, str := trie.Check("我们这里有一个黄色的灯泡，他存在了很久。他是蓝色的。", "*")

	fmt.Printf("result:%#v, str:%v\n", result, str)

}
