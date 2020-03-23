package associate

//Ac自动机根结点
type AcAutoMachine struct {
	root *AcNode
}

//Ac节点
type AcNode struct {
	fail      *AcNode
	isPattern bool
	next      map[rune]*AcNode
}

func NewAcAutoMachine() *AcAutoMachine {
	return &AcAutoMachine{
		root: newAcNode(),
	}
}

func newAcNode() *AcNode {
	return &AcNode{
		fail:      nil,
		isPattern: false,
		next:      map[rune]*AcNode{},
	}
}

//构造trie树
func (ac *AcAutoMachine) AddPattern(pattern string) {
	chars := []rune(pattern)
	iter := ac.root
	for _, c := range chars {
		if _, ok := iter.next[c]; !ok {
			iter.next[c] = newAcNode()
		}
		iter = iter.next[c]
	}
	iter.isPattern = true
}

func (ac *AcAutoMachine) Build() {
	queue := []*AcNode{}
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		parent := queue[0]
		queue = queue[1:]

		for char, child := range parent.next {
			if parent == ac.root {
				child.fail = ac.root
			} else {
				failAcNode := parent.fail
				for failAcNode != nil {
					if _, ok := failAcNode.next[char]; ok {
						child.fail = parent.fail.next[char]
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil {
					child.fail = ac.root
				}
			}
			queue = append(queue, child)
		}
	}
}

func (ac *AcAutoMachine) Search(content string) (results []string) {
	chars := []rune(content)
	iter := ac.root
	var start, end int
	for i, c := range chars {
		_, ok := iter.next[c]
		for !ok && iter != ac.root {
			iter = iter.fail
		}
		if _, ok = iter.next[c]; ok {
			if iter == ac.root {
				start = i
			}
			iter = iter.next[c]
			if iter.isPattern {
				end = i
				results = append(results, string([]rune(content)[start:end+1]))
			}
		}
	}

	return
}
