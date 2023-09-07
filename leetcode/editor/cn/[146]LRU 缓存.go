package main

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	linked  *Linked
	hashMap HashMap
}

// Node 双向链表节点
type Node struct {
	key   int   // key
	value int   // value
	prev  *Node // 节点前节点
	next  *Node // 节点后节点
}

func makeNode(key int, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

// 连接前结点
func (n *Node) linkPrev(prev *Node) {
	if n == nil {
		return
	}
	n.prev = prev
}

// 连接后节点
func (n *Node) linkNext(next *Node) {
	if n == nil {
		return
	}
	n.next = next
}

// 取消节点的连接
func (n *Node) unlink() {
	n.prev = nil
	n.next = nil
}

// Linked 双向链表结构
type Linked struct {
	capacity int   // 双向链表容量
	size     int   // 双向链表长度
	head     *Node // 双向链表头结点
	tail     *Node // 双向链表尾节点
}

// initLinked 初始化双向链表
func initLinked(capacity int) *Linked {
	// 初始化容量，长度，头结点，尾节点
	// 头结点，尾节点不计入长度
	linked := &Linked{
		capacity: capacity,
		size:     0,
		head:     nil,
		tail:     nil,
	}
	head := makeNode(0, 0)
	tail := makeNode(0, 0)
	// 初始化头结点next连接到尾节点,尾节点prev连接到头结点
	head.linkNext(tail)
	tail.linkPrev(head)
	linked.head = head
	linked.tail = tail
	return linked
}

// 向链表添加节点
func (l *Linked) add(newFirst *Node) {
	oldFirst := l.head.next
	// 建立新节点的连接
	newFirst.linkPrev(l.head)
	newFirst.linkNext(oldFirst)
	// 更新头结点的连接
	l.head.linkNext(newFirst)
	// 更新旧节点的连接
	oldFirst.linkPrev(newFirst)
	l.size++
}

// 删除链表中的某个节点
func (l *Linked) delete(delNode *Node) {
	if delNode.prev == nil && delNode.next == nil {
		return
	}
	prev := delNode.prev
	next := delNode.next
	// 要删除节点的上一个节点和下一个节点建立连接
	prev.linkNext(next)
	next.linkPrev(prev)
	delNode.unlink()
	l.size--
}

// 删除最后一个节点
func (l *Linked) deleteFinal() (finalNode *Node) {
	finalNode = l.tail.prev
	l.delete(finalNode)
	return
}

func (l *Linked) cap() int {
	return l.capacity
}

func (l *Linked) len() int {
	return l.size
}

// 删除节点，然后把节点放到头部
func (l *Linked) deleteAndAdd(node *Node) {
	// 在其他位置删除该节点
	l.delete(node)
	// 将该节点放到头部
	l.add(node)
}

// 整理
func (this *LRUCache) tidy() {
	if this.linked.len() > this.linked.cap() {
		f := this.linked.deleteFinal()
		this.hashMap.delete(f.key)
	}
	return
}

type HashMap map[int]*Node

func (h HashMap) delete(key int) {
	delete(h, key)
}

func (h HashMap) add(key int, node *Node) {
	h[key] = node
}

func Constructor(capacity int) LRUCache {
	linked := initLinked(capacity)
	hashMap := HashMap{}
	lru := LRUCache{
		linked:  linked,
		hashMap: hashMap,
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	this.linked.deleteAndAdd(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hashMap[key]
	if !ok {
		// 不存在，初始化该节点
		node = makeNode(key, value)
	}
	this.linked.deleteAndAdd(node)
	this.hashMap[key] = node
	this.tidy()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
