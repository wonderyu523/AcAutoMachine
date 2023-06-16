package acAutoMachine

import (
	"fmt"
	"testing"
)

func TestAc1(t *testing.T) {
	content := "还没为你把红豆，熬成缠绵的伤口，然后一起分享，会更明白相思的哀愁。"
	ac := NewAcAutoMachine()
	ac.Add("红豆")
	ac.Add("一起")
	ac.Add("相思")
	ac.Build()
	results := ac.Search(content)
	fmt.Println("内容: " + content)
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestAc2(t *testing.T) {
	content := "还没为你把红豆，熬成缠绵的伤口，然后一起分享，会更明白相思的哀愁。"
	ac := NewAcAutoMachine()
	ac.Add("红豆")
	ac.Add("一起")
	ac.Add("相思")
	ac.Build()
	exceptWordMap := make(map[string]int)
	exceptWordMap["红豆"] = 1
	_, highLightMarkList, _:= ac.HighlightSearch(content, exceptWordMap)
	fmt.Println("内容: " + content)
	for _, highLight := range highLightMarkList {
		fmt.Printf("Highlight Location:%s, Length:%s\n", highLight.Location, highLight.Length)
	}
}