package main

import (
	"fmt"
	"github.com/jshzhj/associate/core"
)

func main() {
	//obj := core.NewTrie()
	//obj.PrefixAdd("相框")
	//obj.PrefixAdd("相框摆")
	//obj.PrefixAdd("相框摆台")
	//obj.PrefixAdd("相框挂")
	//obj.PrefixAdd("相框挂台")
	//
	////前缀搜索,自动推荐
	//data := obj.PrefixSearch("相框", 3)
	//for _, v := range data {
	//	fmt.Println(v)
	//}
	//
	////查看当前字符串是否在trie树中
	//mgc := obj.Search("相")
	//fmt.Println(mgc)

	//Ac自动机-todo
	//ac := core.NewAcAutoMachine()
	//ac.AddPattern("垃圾")
	//ac.AddPattern("文章")
	//ac.AddPattern("真")
	//ac.Build()
	//
	//content := "这篇文章真的是好垃圾啊"
	//results := ac.Search(content)
	//for _, result := range results {
	//	fmt.Println(result)
	//}

	//敏感词替换

	trie := core.NewTrie()
	//将敏感词加入到Trie中
	trie.Add("黄色", "黄色啊a你好")
	trie.Add("绿色", "黄色啊a你好")
	trie.Add("蓝色", "黄色啊a你好")

	result, str := trie.Replace("我们这里有一个黄色的灯泡，他存在了很久。他是蓝色的。是黄色色的","*")

	fmt.Printf("result:%#v, str:%v\n", result, str)

}
