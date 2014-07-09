//
//  arraylist.go
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
import "bytes"
import "fmt"
import "math/rand"
import "time"

type ArrayList struct {
	count int
	lock *sync.Mutex
	items []interface{}
}

func NewArrayList() *ArrayList {
	instance := &ArrayList{}
	
	instance.lock = &sync.Mutex{}
	instance.items = make([]interface{}, 10)
	instance.count = 0
	
	rand.Seed( time.Now().UTC().UnixNano())

	return instance
}

func (self *ArrayList) Len() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.count
}

func (self *ArrayList) IsEmpty() bool {
	return self.Len() == 0
}

func (self *ArrayList) Add(objects ...interface{}) {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	for _, o := range objects {
		self.add(o)
	}
}

func (self *ArrayList) add(o interface{}) {
	self.items[self.count] = o
	self.count++
	self.resize_slice_if_necessary()
}



func (self *ArrayList) resize_slice_if_necessary() {
	capacity := cap(self.items)
	
	if self.count >= (capacity-1) {
		newCapacity := (capacity+1)*2
		temp := make([]interface{}, newCapacity, newCapacity)
		copy(temp, self.items)
		self.items = temp
	}
}

func (self *ArrayList) ToSlice() []interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	result := make([]interface{}, self.count)
	copy(result, self.items)
	
	return result
}

func (self *ArrayList) Get(index int) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.items[index]
}

func (self *ArrayList) Sample() interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if (self.count == 0) {
		return nil
	}
	index := rand.Intn(self.count)
	return self.items[index]
}

func (self *ArrayList) IndexOf(o interface{}) int {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.index_of(o)
}

func (self *ArrayList) index_of(o interface{}) int {
	index := -1
	for i := 0; i < self.count; i++ {
		if self.items[i] == o {
			index = i
			break;
		}
	}
	return index
}

func (self *ArrayList) Contains(o interface{}) bool {
	return self.IndexOf(o) != -1
}

func (self *ArrayList) Remove(o interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	index := self.index_of(o)
	
	if index == -1 {
		return false
	}
	
	self.items[index] = nil
	
	for i := index; i < self.count-1; i++ {
		self.swap(i, i+1)
	}
	self.count--
	return true
}

func (self *ArrayList) Swap(x int, y int) {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	self.swap(x, y)
}

func (self *ArrayList) swap(x int, y int) {
	self.items[x], self.items[y] = self.items[y], self.items[x]
}

func (self *ArrayList) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	capacity := cap(self.items)
	length := len(self.items)
	self.items = make([]interface{}, length, capacity)
	self.count = 0
}

func (self *ArrayList) AddFromArrayList(arrayList *ArrayList) {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	if arrayList == nil {
		return;
	}
	
	for i := 0; i < arrayList.Len(); i++ {
		self.add(arrayList.Get(i))
	}
}

func (self *ArrayList) First() interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.items[0]	
}

func (self *ArrayList) Last() interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	return self.items[self.count-1]	
}

func (self *ArrayList) String() string {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	var buffer bytes.Buffer
	
	for i := 0; i < self.count; i++ {
		item := self.items[i]
		stringify := fmt.Sprintf("%s", item)
		buffer.WriteString(stringify)
		if i != (self.count-1) {
			buffer.WriteString(", ")
		}
	}
	return fmt.Sprintf("[ %s ]", buffer.String())
}
