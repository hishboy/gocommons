package lang

import (
	//"fmt"
//	"testing"
//	"time"

	gc "gopkg.in/check.v1" //gocheck testing framework
)

// Hook up gc (gocheck)  into the "go test" runner.
//func Test(t *testing.T) { gc.TestingT(t) }

type LRUCacheTest struct{}

var _ = gc.Suite(&LRUCacheTest{})


func (s *LRUCacheTest) TestZeroCapacityCache(c *gc.C) {
	cache := NewLRUCache(0)
    cache.Put("key", "value")
	c.Assert(cache.Len(), gc.Equals, 0)
}

func (s *LRUCacheTest) TestOneCapacityCache(c *gc.C) {
	cache := NewLRUCache(1)
    cache.Put("key", "value")
    cache.Put("key2", "value2")
	c.Assert(cache.Len(), gc.Equals, 1)
}

func (s *LRUCacheTest) TestGetAndPut(c *gc.C) {
	cache := NewLRUCache(5)
    
    cache.Put("key0", "value0")
    cache.Put("key1", "value1")
    cache.Put("key2", "value2")
    cache.Put("key3", "value3")
    cache.Put("key4", "value4")
    
    c.Assert(cache.Get("key0"), gc.Equals, "value0")
    c.Assert(cache.Get("key1"), gc.Equals, "value1")
    c.Assert(cache.Get("key2"), gc.Equals, "value2")
    c.Assert(cache.Get("key3"), gc.Equals, "value3")
    c.Assert(cache.Get("key4"), gc.Equals, "value4")
	c.Assert(cache.Len(), gc.Equals, 5)
}

func (s *LRUCacheTest) TestRemove(c *gc.C) {
	cache := NewLRUCache(5)
    
    cache.Put("key0", "value0")
    cache.Put("key1", "value1")
    cache.Put("key2", "value2")
    cache.Put("key3", "value3")
    cache.Put("key4", "value4")
    
    cache.Remove("key1")
    cache.Remove("key2")
    cache.Remove("key3")
    
    c.Assert(cache.Get("key0"), gc.Equals, "value0")
    c.Assert(cache.Get("key1"), gc.Equals, nil)
    c.Assert(cache.Get("key2"), gc.Equals, nil)
    c.Assert(cache.Get("key3"), gc.Equals, nil)
    c.Assert(cache.Get("key4"), gc.Equals, "value4")
	c.Assert(cache.Len(), gc.Equals, 2)
}

func (s *LRUCacheTest) TestDumpingOfLeastUsedItems(c *gc.C) {
	cache := NewLRUCache(3)
    
    cache.Put("key0", "value0")
    cache.Put("key1", "value1")
    cache.Put("key2", "value2")
    cache.Put("key3", "value3")
    cache.Put("key4", "value4")
    
    c.Assert(cache.Get("key0"), gc.Equals, nil)
    c.Assert(cache.Get("key1"), gc.Equals, nil)
    c.Assert(cache.Get("key2"), gc.Equals, "value2")
    c.Assert(cache.Get("key3"), gc.Equals, "value3")
    c.Assert(cache.Get("key4"), gc.Equals, "value4")
	c.Assert(cache.Len(), gc.Equals, 3)
}

func (s *LRUCacheTest) TestDumpingOfLeastUsedItemsAfterGets(c *gc.C) {
	cache := NewLRUCache(3)
    
    cache.Put("key0", "value0")
    cache.Put("key1", "value1")
    cache.Put("key2", "value2")
    cache.Get("key0")
    cache.Put("key3", "value3")
    
    c.Assert(cache.Get("key0"), gc.Equals, "value0")
    c.Assert(cache.Get("key1"), gc.Equals, nil)
    c.Assert(cache.Get("key2"), gc.Equals, "value2")
    c.Assert(cache.Get("key3"), gc.Equals, "value3")
    c.Assert(cache.Get("key4"), gc.Equals, nil)
	c.Assert(cache.Len(), gc.Equals, 3)
}

func (s *LRUCacheTest) TestDumpingOfLeastUsedItemsAfterRemove(c *gc.C) {
	cache := NewLRUCache(3)
    
    cache.Put("key0", "value0")
    cache.Put("key1", "value1")
    cache.Put("key2", "value2")
    cache.Get("key0")
    cache.Put("key3", "value3")
    cache.Remove("key3")
    cache.Put("key4", "value4")
    
    c.Assert(cache.Get("key0"), gc.Equals, "value0")
    c.Assert(cache.Get("key1"), gc.Equals, nil)
    c.Assert(cache.Get("key2"), gc.Equals, "value2")
    c.Assert(cache.Get("key3"), gc.Equals, nil)
    c.Assert(cache.Get("key4"), gc.Equals, "value4")
	c.Assert(cache.Len(), gc.Equals, 3)
}
