package associate

import "testing"

func TestAcAutoMachine(t *testing.T) {
	//Ac自动机-todo
	ac := NewAcAutoMachine()
	ac.AddPattern("垃圾")
	ac.AddPattern("文章")
	ac.AddPattern("真")
	ac.Build()

}
