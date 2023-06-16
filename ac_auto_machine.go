package acAutoMachine

import (
	"fmt"
)

type AcNode struct {
	fail      *AcNode
	isTail    bool
	children  map[rune]*AcNode
	prefix    []rune
	depth     int
}

func newAcNode() *AcNode {
	return &AcNode{
		fail: 		nil,
		isTail: 	false,
		children: 	map[rune]*AcNode{},
		prefix: 	[]rune{},
		depth: 	    0,
	}
}

type AcAutoMachine struct {
	root *AcNode
	patterns map[string]int
}

func NewAcAutoMachine() *AcAutoMachine {
	return &AcAutoMachine{
		root: newAcNode(),
		patterns: map[string]int{},
	}
}

func (ac *AcAutoMachine) Add(pattern string) {
	if _, ok := ac.patterns[pattern]; ok {
		return
	}
	chars := []rune(pattern)
	iter := ac.root
	for _, ch := range chars {
		if _, ok := iter.children[ch]; !ok {
			iter.children[ch] = newAcNode()
			iter.children[ch].prefix = append(iter.prefix, ch)
			iter.children[ch].depth = iter.depth + 1
		}
		iter = iter.children[ch]
	}
	iter.isTail = true
	ac.patterns[pattern] = 1
}

func (ac *AcAutoMachine) Build() {
	queue := make([]*AcNode, 0)
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		currNode := queue[0]
		queue = queue[1:]
		for ch, childNode := range currNode.children {
			if currNode == ac.root {
				childNode.fail = ac.root
			} else {
				failAcNode := currNode.fail
				for failAcNode != nil {
					if _, ok := failAcNode.children[ch]; ok {
						childNode.fail = failAcNode.children[ch]
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil {
					childNode.fail = ac.root
				}
			}
			queue = append(queue, childNode)
		}
	}
}

func (ac *AcAutoMachine) Search(target string) (result []string) {
	result = make([]string, 0)
	chars := []rune(target)
	iter := ac.root
	for _, ch := range chars {
		_, ok := iter.children[ch]
		for !ok && iter != ac.root {
			iter = iter.fail
			_, ok = iter.children[ch]
		}
		if ok {
			iter = iter.children[ch]
		} else {
			iter = ac.root
		}
		temp := iter
		for temp != ac.root {
			if temp.isTail {
				result = append(result, string(temp.prefix))
			}
			temp = temp.fail
		}
	}
	return
}

type HighLightMark struct {
	Location   string
	Length     string
	Color      string
}

func (ac *AcAutoMachine) HighlightSearch(target string, exceptPattern map[string]int) (result []string, highLightMark []HighLightMark, matchPattern map[string]int) {
	result = make([]string, 0)
	highLightMark = make([]HighLightMark, 0)
	matchPattern = make(map[string]int)
	if exceptPattern == nil {
		exceptPattern = make(map[string]int)
	}
	chars := []rune(target)
	iter := ac.root
	for index, ch := range chars {
		_, ok := iter.children[ch]
		for !ok && iter != ac.root {
			iter = iter.fail
			_, ok = iter.children[ch]
		}
		if ok {
			iter = iter.children[ch]
		} else {
			iter = ac.root
		}
		temp := iter
		for temp != ac.root {
			if temp.isTail {
				_, exist := exceptPattern[string(temp.prefix)]
				if !exist {
					result = append(result, string(temp.prefix))
					var targetMark = HighLightMark{
						Location: fmt.Sprintf("%d", index - temp.depth + 1),
						Length:   fmt.Sprintf("%d", temp.depth),
					}
					highLightMark = append(highLightMark, targetMark)
					matchPattern[string(temp.prefix)] = 1
				}
			}
			temp = temp.fail
		}
	}
	return
}