# Associate

智能提示,敏感词替换

#### 用法

```go

package main

import (
	"fmt"
	"github.com/jshzhj/associate"
)

func main(){
	obj := associate.NewTrie()
	//(1)前缀搜索,智能提示
	//第一个参数为要加入Trie中的字符串,第二个参数为字符串叶子节点保存的数据
	obj.Add("相框", "相框")
	obj.Add("相框摆", "相框摆")
	obj.Add("相框摆台", "相框摆台")
	obj.Add("相框挂", "相框挂")
	obj.Add("相框挂台", "相框挂台")

	data := obj.PrefixSearch("相框", 4)
	for _, v := range data {
		fmt.Println(v)
	//输出:相框 相框挂 相框摆 相框摆台
	}

	//(2)查看当前字符串是否在trie树中
	mgc := obj.Search("相")
	fmt.Println(mgc)
        //输出:false

	//(3)敏感词替换
	trie := associate.NewTrie()
	//将敏感词加入到Trie中,
	trie.Add("tm", "tm")
	trie.Add("电影","电影")
	trie.Add("他妈", nil)

	result, str := trie.Replace("这个电影真tm的是难看,好他妈难看啊", "*")
	fmt.Printf("result:%#v, str:%v\n", result, str)
    //输出:result:"这个**真**的是难看,好*难看啊", str:true

}


```