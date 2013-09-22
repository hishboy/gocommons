package lang


import "sync"

type stacknode struct {
	data interface{}
	next *stacknode
}


type Stack struct {
	head *stacknode
	count int
	lock *sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

func (s *Stack) Count() int {
	return s.count
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	n := &stacknode { data: item }
	
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	
	s.count++
	s.lock.Unlock()
}
 
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	
	var n *stacknode
	if s.head != nil {
		n = s.head
		s.head = n.next
		s.count--
	}
	
	s.lock.Unlock()
	
	if n == nil {
		return nil
	}
	
	return n.data
	
}