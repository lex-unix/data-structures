package list

type CircularList struct {
	sentinel *Node
}

func NewCircularList() *CircularList {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &CircularList{sentinel: sentinel}
}

func (l *CircularList) Prepend(key int) {
	x := &Node{key: key}
	x.next = l.sentinel.next
	x.prev = l.sentinel
	l.sentinel.next.prev = x
	l.sentinel.next = x
}

func (l *CircularList) Append(key int) {
	x := &Node{key: key}
	x.next = l.sentinel
	x.prev = l.sentinel.prev
	l.sentinel.prev.next = x
	l.sentinel.prev = x
}

func (l *CircularList) Search(key int) *Node {
	l.sentinel.key = key
	x := l.sentinel.next
	for x.key != key {
		x = x.next
	}

	if x == l.sentinel {
		return nil
	}
	return x
}

func (l *CircularList) DeleteAt(i int) {
	if i < 0 {
		return
	}
	current := l.sentinel.next
	for pos := 0; current != l.sentinel && pos != i; pos++ {
		current = current.next
	}
	if current == l.sentinel {
		return
	}
	l.Delete(current)
}

func (l *CircularList) DeleteKey(key int) {
	node := l.Search(key)
	if node == nil {
		return
	}
	l.Delete(node)
}

func (l *CircularList) Delete(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
