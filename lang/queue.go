
package lang


import "sync"

type QueueNode struct {
	data interface{}
	next *QueueNode
}


type Queue struct {
	head *QueueNode
	tail *QueueNode
	count int
	lock *sync.Mutex
}

func NewQueue() *Queue {
 	q := &Queue{}
  q.lock = &sync.Mutex{}
	return q
}

func (q *Queue) Count() int {
	return q.count
}

func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	n := &QueueNode { data: item }
	
	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
	
	q.lock.Unlock()
}
 
func (q *Queue) Poll() interface{} {
	q.lock.Lock()
	
	if q.head == nil {
		q.lock.Unlock()
		return nil
	}
	
	n := q.head
	q.head = n.next
	
	if q.head == nil {
		q.tail = nil
	}
	q.count--
	
	q.lock.Unlock()
	return n.data
	
}