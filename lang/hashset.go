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

func (self *HashSet) Count() int {
	self.lock.Lock()
	defer self.lock.Unlock()
	
	return len(self.items)
}

func (self *HashSet) IsEmpty() bool {
	return self.Count() == 0
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