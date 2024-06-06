package queue

import "errors"

type Queue struct {
	head     int
	tail     int
	size     int
	elements []int
}

func NewQueue(size int) *Queue {
	return &Queue{
		head:     0,
		tail:     0,
		size:     size,
		elements: make([]int, size),
	}
}

func (q *Queue) Enqueue(elem int) error {
	if (q.head == q.tail+1) || (q.head == 0 && q.tail == q.size-1) {
		return errors.New("queue overflow")
	}
	q.elements[q.tail] = elem
	if q.tail == q.size-1 {
		q.tail = 0
	} else {
		q.tail += 1
	}
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.head == q.tail {
		return -1, errors.New("queue underflow")
	}
	x := q.elements[q.head]
	if q.head == q.size-1 {
		q.head = 0
	} else {
		q.head += 1
	}
	return x, nil
}
