//
//  stack.go
//
//  Created by Hicham Bouabdallah
//  Copyright (c) 2012 SimpleRocket LLC
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//

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

func (s *Stack) Len() int {
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

func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.head
	if n == nil || n.data == nil {
		return nil
	}
	
	return n.data
}