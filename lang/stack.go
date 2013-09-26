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
	s.lock.Lock()
	defer s.lock.Unlock()
	
	return s.count
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	n := &stacknode { data: item }
	
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	
	s.count++
}
 
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	var n *stacknode
	if s.head != nil {
		n = s.head
		s.head = n.next
		s.count--
	}
	
	if n == nil {
		return nil
	}
	
	return n.data
	
}