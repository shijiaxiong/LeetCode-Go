package main

import "fmt"

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

// 双向链表 头部存储较新访问的结点，尾部是当前频次最旧的结点
type Node struct {
	Key, Val   int
	Prev, Next *Node
}

// 使用哨兵首位节点简化流程
func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		head: head,
		tail: tail,
		Keys: make(map[int]*Node),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		// 更新值
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	}

	if len(this.Keys) == this.Cap {
		delete(this.Keys, this.tail.Prev.Key)
		this.Remove(this.tail.Prev)
	}

	node := &Node{Key: key, Val: value}
	this.Keys[key] = node
	this.Add(node)

	return
}

// 头(新)           	 中			尾(旧)
// head(node1)   	node2		tail(node3)

// 移除节点
func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	return
}

// head  tail是链表的首尾
// 将节点添加在首
func (this *LRUCache) Add(node *Node) {
	head := this.head

	// 先把node和head.Next连接，再连接node和head
	node.Next = head.Next
	head.Next.Prev = node

	node.Prev = head
	head.Next = node
}

func main() {
	cache := Constructor(1)
	cache.Put(0, 0)
	cache.Put(1, 1)

	fmt.Println(cache)
}
