//
//  hashset.go
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

type HashSet struct {
	lock *sync.Mutex
	items map[interface{}]interface{}
}


func NewHashSet() *HashSet {
	instance := &HashSet{}
	
	instance.lock = &sync.Mutex{}
	instance.items = make(map[interface{}]interface{})
	
	return instance
}

func (self *HashSet) Len() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return len(self.items)
}

func (self *HashSet) ToSlice() []interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	var result []interface{}
	for k, _ := range self.items {
		result = append(result, k)
	}
	
	return result
}

func (self *HashSet) IsEmpty() bool {
	return self.Len() == 0
}

func (self *HashSet) Add(objects ...interface{}) {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	for _, o := range objects {
		self.items[o] = true
	}
}

func (self *HashSet) Get(k interface{}) interface{} {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return self.items[k]
}

func (self *HashSet) Contains(k interface{}) bool {
	return self.Get(k) != nil
}

func (self *HashSet) Remove(k interface{}) bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	found := self.items[k]
	if found != nil {
		delete(self.items, k)
		return true
	} else {
		return false
	}
}

func (self *HashSet) Clear() {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.items = make(map[interface{}]interface{})
}

func (self *HashSet) String() string {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	var buffer bytes.Buffer
	
	i := 0
	for k, _ := range self.items {
		stringify := fmt.Sprintf("%s", k)
		buffer.WriteString(stringify)
		if i != len(self.items)-1 {
			buffer.WriteString(", ")
		}
		i++
	}
	return fmt.Sprintf("{ %s }", buffer.String())
}