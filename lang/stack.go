package lang


import "sync"

type StackNode struct {
	data interface{}
	next *StackNode
}


type Stack struct {
	head *StackNode
	count int
	lock *sync.Mutex
}

func NewStack() *Stack {
 	q := &Stack{}
  q.lock = &sync.Mutex{}
	return q
}

func (q *Stack) Push(item interface{}) {
	q.lock.Lock()
	n := &StackNode { data: item }
	
	if q.head == nil {
		q.head = n
	} else {
		n.next = q.head
		q.head = n
	}
	
	q.lock.Unlock()
}
 
func (q *Stack) Pop() interface{} {
	q.lock.Lock()
	
	var n *StackNode
	if q.head != nil {
		n = q.head
		q.head = n.next
	}
	
	q.lock.Unlock()
	if n == nil {
		return nil
	}
	return n.data
	
}