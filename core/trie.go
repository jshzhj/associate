package core

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

	var parent *Node = t.root //1:将根结点的node赋值给parent(父节点,第一次时候,父节点就是根节点)
	//将根结点指针地址赋给parent,那么对parent的修改就是对根结点的修改
	allChars := []rune(key) //2:将字符串转成字符,遍历它

	for _, char := range allChars {
		node, ok := parent.childs[char] //查看当前字符是否在根结点的childs中
		if !ok {
			//如果不在,则新创建这个字符的node,深度=父节点的deep+1

			node = NewNode(char, parent.deep+1)
			parent.childs[char] = node //修改根结点,将这个新建的node加入根结点的childs中
		}
		//如果在,则继续往下找,将查找到的这个node,赋值给parent变量,作为下一次遍历的根结点
		//下一次直接从新的parent这个变量中继续查找,一直到循环结束,将这个新增的字符串遍历完
		parent = node
	}
	//将所有字符遍历完成之后,这个字符串就已经加入Trie树了
	//这时,parent这个变量为最后一个字符的node
	parent.Data = data   //尾节点保存数据
	parent.isTerm = true //是一个字符串的结束
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

	chars := []rune(text) //将字符串打散成字符切片
	if t.root == nil {
		return
	}

	var left []rune
	node := t.root //第一次为根节点的root,
	start := 0
	//遍历字符数组
	for index, v := range chars {
		ret, ok := node.childs[v] //看这个字符是否在根节点childs中,ret是*Node,ok判断是否在map中
		if !ok {
			//如果不在,加入到left切片中, start:截取的起始下标(含),index+1:结束的截取下标(不含),
			//三个点,两个切片合并,第二个切片打散加入第一个切片中
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = t.root
			continue
		}
		//如果在trie中,将这个在根节点clilds中的字符给node,用*号替代,之后继续
		node = ret
		if ret.isTerm {
			//在trie中找到完整的词,之后将node恢复原位,到根的root,将*放入left中
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
