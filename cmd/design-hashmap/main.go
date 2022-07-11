package main

const MOD = 100007

type Node struct {
	Val   int
	Value int
	Pre   *Node
	Next  *Node
}

type MyHashMap struct {
	arr [MOD]*Node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (s *MyHashMap) Put(key int, value int) {
	idx := key % MOD

	var pre *Node
	for node := s.arr[idx]; node != nil; pre, node = node, node.Next {
		if node.Val == key {
			node.Value = value
			return
		}
	}
	if pre == nil {
		s.arr[idx] = &Node{key, value, nil, nil}
	} else {
		pre.Next = &Node{key, value, pre, nil}
	}
}

func (s *MyHashMap) Get(key int) int {
	idx := key % MOD
	for node := s.arr[idx]; node != nil; node = node.Next {
		if node.Val == key {
			return node.Value
		}
	}
	return -1
}

func (s *MyHashMap) Remove(key int) {
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

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
