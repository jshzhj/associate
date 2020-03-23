package associate

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	obj := NewTrie()
	//第一个参数为要加入Trie中的字符串,第二个参数为字符串叶子节点保存的数据
	obj.Add("相框", "相框")
	obj.Add("相框摆", "相框摆")
	obj.Add("相框摆台", "相框摆台")
	obj.Add("相框挂", "相框挂")
	obj.Add("相框挂台", "相框挂台")

	//前缀搜索,自动推荐
	data := obj.PrefixSearch("相框", 3)
	for _, v := range data {
		fmt.Println(v)
	}

	//查看当前字符串是否在trie树中
	mgc := obj.Search("相")
	fmt.Println(mgc)

	//Ac自动机-todo
	//ac := core.NewAcAutoMachine()
	//ac.AddPattern("垃圾")
	//ac.AddPattern("文章")
	//ac.AddPattern("真")
	//ac.Build()

	//敏感词替换
	trie := NewTrie()
	//将敏感词加入到Trie中,
	trie.Add("tm", "tm")
	trie.Add("他妈", nil)

	result, str := trie.Replace("这个电影真tm的是难看,好他妈难看啊", "*")

	fmt.Printf("result:%#v, str:%v\n", result, str)
}
