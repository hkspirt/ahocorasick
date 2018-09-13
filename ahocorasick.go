//----------------
//Func  : Aho Corasick Word Match 敏感词匹配
//Author: xjh
//Date  : 2018/09/13
//Note  : 基于github.com/gansidui/ahocorasick
//        支持线程安全
//        支持中文(UTF8)
//        比使用正则匹配regexp有上千倍性能提升
//----------------

package ahocorasick

import (
	"container/list"
)

type trieNode struct {
	count int
	fail  *trieNode
	child map[rune]*trieNode
	index int
}

func newTrieNode() *trieNode {
	return &trieNode{
		count: 0,
		fail:  nil,
		child: make(map[rune]*trieNode),
		index: -1,
	}
}

type ACMatcher struct {
	root *trieNode
	size int
}

func NewMatcher(dict []string) *ACMatcher {
	m := &ACMatcher{
		root: newTrieNode(),
		size: 0,
	}

	for i := range dict {
		m.insert(dict[i])
	}
	m.build()
	return m
}

//包含敏感词位置、个数
func (m *ACMatcher) Match(s string) []int {
	curNode := m.root
	var p *trieNode = nil
	ret := make([]int, 0)
	mark := make(map[int]bool)
	for _, v := range s {
		for curNode.child[v] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = m.root
		}
		p = curNode
		for p != m.root && p.count > 0 && !mark[p.index] {
			mark[p.index] = true
			for i := 0; i < p.count; i++ {
				ret = append(ret, p.index)
			}
			p = p.fail
		}
	}
	return ret
}

//是否包含敏感词，查找到任意立即返回
func (m *ACMatcher) Has(s string) bool {
	curNode := m.root
	var p *trieNode = nil
	for _, v := range s {
		for curNode.child[v] == nil && curNode != m.root {
			curNode = curNode.fail
		}
		curNode = curNode.child[v]
		if curNode == nil {
			curNode = m.root
		}
		p = curNode
		for p != m.root && p.count > 0 {
			return true
		}
	}
	return false
}

func (m *ACMatcher) Size() int {
	return m.size
}

func (m *ACMatcher) build() {
	ll := list.New()
	ll.PushBack(m.root)
	for ll.Len() > 0 {
		temp := ll.Remove(ll.Front()).(*trieNode)
		var p *trieNode = nil
		for i, v := range temp.child {
			if temp == m.root {
				v.fail = m.root
			} else {
				p = temp.fail
				for p != nil {
					if p.child[i] != nil {
						v.fail = p.child[i]
						break
					}
					p = p.fail
				}
				if p == nil {
					v.fail = m.root
				}
			}
			ll.PushBack(v)
		}
	}
}

func (m *ACMatcher) insert(s string) {
	curNode := m.root
	for _, v := range s {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode()
		}
		curNode = curNode.child[v]
	}
	curNode.count++
	curNode.index = m.size
	m.size++
}
