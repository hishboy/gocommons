//
//  arraylist.go
//
//  Created by Hicham Bouabdallah
//  Copyright (c) 2014 SimpleRocket LLC
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
//import "fmt"

type lrunode struct {
	value interface{}
    key interface{}
	next *lrunode
    previous *lrunode
}

type LRUCache struct {
	lock *sync.Mutex
    head *lrunode
    tail *lrunode
    capacity int
	items map[interface{}]*lrunode
}

func NewLRUCache(capacity int) *LRUCache {
	instance := &LRUCache{}
	
	instance.lock = &sync.Mutex{}
	instance.items = make(map[interface{}]*lrunode)
    instance.capacity = capacity

	return instance
}

func (self *LRUCache) Len() int {
	self.lock.Lock()
	defer self.lock.Unlock()
    
    return len(self.items)
}

func (self *LRUCache) Put(key interface{}, value interface{}) {
    self.lock.Lock()
    defer self.lock.Unlock()

    node := self.items[key]

    if node == nil {
        node = &lrunode {}
        node.key = key 
        self.items[key] = node
    } else {
        self.remove_node_from_dll(node)
    }
    
    node.value = value
    self.add_item_to_head_of_dll(node)
    self.dump_objects_if_reached_capacity()
}

func (self *LRUCache) Remove(key interface{}) {
    self.lock.Lock()
    defer self.lock.Unlock()

    node := self.items[key]

    if node != nil {
        self.remove_node_from_dll(node)
        delete(self.items, key)
    }
}

func (self *LRUCache) Get(key interface{}) interface{} {
    self.lock.Lock()
    defer self.lock.Unlock()

    node := self.items[key]

    if node == nil {
        return nil
    } else {
        self.remove_node_from_dll(node)
        self.add_item_to_head_of_dll(node)
        return node.value
    }
}

func (self *LRUCache)remove_node_from_dll(item *lrunode) {
    previous_node := item.previous
    next_node := item.next

    if previous_node != nil && next_node != nil {
        previous_node.next = next_node
        next_node.previous = previous_node
    } else if next_node != nil {
        self.head = item.next
        self.head.previous = nil
    } else if previous_node != nil {
        self.tail = item.previous
        self.tail.next = nil
    }

    item.next = nil
    item.previous = nil
}

func (self *LRUCache)add_item_to_head_of_dll(item *lrunode) {
    if self.head == nil {
        self.head = item
        self.tail = item
        item.next = nil
        item.previous = nil
    } else {
        old_head := self.head
        old_head.previous = item
        self.head = item
        self.head.next = old_head
        self.head.previous = nil
    }
}

func (self *LRUCache)dump_objects_if_reached_capacity() {
        // fmt.Println(len(self.items))
        // fmt.Println("********************************")
    for ;len(self.items) > self.capacity; {
        least_used_node := self.tail
        delete(self.items, least_used_node.key)
        self.tail = self.tail.previous
    }
}
