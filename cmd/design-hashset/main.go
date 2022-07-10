package main

import (
	"fmt"
)

const MOD = 1007

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

type MyHashSet struct {
	arr [MOD]*Node
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (s *MyHashSet) Add(key int) {
	idx := key % MOD

	var pre *Node
	for node := s.arr[idx]; node != nil; pre, node = node, node.Next {
		if node.Val == key {
			return
		}
	}
	if pre == nil {
		s.arr[idx] = &Node{key, nil, nil}
	} else {
		pre.Next = &Node{key, pre, nil}
	}
}

func (s *MyHashSet) Remove(key int) {
	idx := key % MOD

	for node := s.arr[idx]; node != nil; node = node.Next {
		if node.Val == key {
			if node.Pre == nil {
				s.arr[idx] = node.Next
			} else {
				node.Pre.Next = node.Next
			}
			return
		}
	}
}

func (s *MyHashSet) Contains(key int) bool {
	idx := key % MOD
	for node := s.arr[idx]; node != nil; node = node.Next {
		if node.Val == key {
			return true
		}
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(1007)
	obj.Remove(1)
	obj.Remove(1)
	obj.Add(1)
	obj.Add(1)
	param_3 := obj.Contains(1007)
	if param_3 != true {
		fmt.Errorf("xx")
	}
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
