package lang

import (
	//"fmt"
	//logLib "log"
	//"os"
	"testing"
	"time"
	//"reflect"

	gc "gopkg.in/check.v1" //gocheck testing framework
)

//var log = logLib.New(os.Stderr, "-->: ", logLib.LstdFlags|logLib.Lshortfile)

// Hook up gc (gocheck)  into the "go test" runner.
func Test(t *testing.T) { gc.TestingT(t) }

type Suite struct{}

var _ = gc.Suite(&Suite{})

//---- Test ----

//Test for correct inputs

func (s *Suite) TestLengthZero(c *gc.C) {
	q := NewQueue()
	c.Assert(q.Len(), gc.Equals, 0)
}

//--Test Push

func (s *Suite) TestSimplePoll(c *gc.C) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	c.Assert(q.Len(), gc.Equals, 1)
	val := q.Peek()
	c.Assert(val, gc.Equals, expected)
	c.Assert(q.Len(), gc.Equals, 1)
}

func (s *Suite) TestPushNils(c *gc.C) {
	q := NewQueue()
	for i := 0; i < 15; i++ {
		q.Push(nil)
		c.Assert(q.Len(), gc.Equals, i+1)
	}
	c.Assert(q.Len(), gc.Equals, 15)
	for i := 15; i > 0; i-- {
		c.Assert(q.Len(), gc.Equals, i)
		val := q.Poll()
		c.Assert(val, gc.Equals, nil)
		c.Assert(q.Len(), gc.Equals, i-1)
	}
}

func (s *Suite) TestMixedPushes(c *gc.C) {
	q := NewQueue()
	c.Assert(q.Len(), gc.Equals, 0)

	//Pushes
	q.Push(nil)
	c.Assert(q.Len(), gc.Equals, 1)

	q.Push(10)
	c.Assert(q.Len(), gc.Equals, 2)

	q.Push("foo")
	c.Assert(q.Len(), gc.Equals, 3)

	q.Push([]int{1, 2, 3})
	c.Assert(q.Len(), gc.Equals, 4)

	q.Push(nil)
	c.Assert(q.Len(), gc.Equals, 5)

	q.Push("bar")
	c.Assert(q.Len(), gc.Equals, 6)

	q.Push(nil)
	c.Assert(q.Len(), gc.Equals, 7)

	q.Push(nil)
	c.Assert(q.Len(), gc.Equals, 8)

	//Polls
	var val interface{}
	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 7)
	c.Assert(val, gc.IsNil)

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 6)
	c.Assert(val, gc.Equals, 10)

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 5)
	c.Assert(val, gc.Equals, "foo")

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 4)
	c.Assert(val, gc.DeepEquals, []int{1, 2, 3})

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 3)
	c.Assert(val, gc.IsNil)

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 2)
	c.Assert(val, gc.Equals, "bar")

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 1)
	c.Assert(val, gc.IsNil)

	val = q.Poll()
	c.Assert(q.Len(), gc.Equals, 0)
	c.Assert(val, gc.IsNil)
}

//--Test Poll

func (s *Suite) TestPullNothing(c *gc.C) {
	q := NewQueue()
	val := q.Poll()
	c.Assert(val, gc.IsNil)
	c.Assert(q.Len(), gc.Equals, 0)
}

func (s *Suite) TestMultiplePolls(c *gc.C) {
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
		c.Assert(q.Peek(), gc.Equals, 0)
		c.Assert(q.Len(), gc.Equals, i+1)
	}
	i := 0
	j := 10
	for i < 10 {
		length := j
		c.Assert(q.Len(), gc.Equals, length)
		val := q.Poll()
		c.Assert(val, gc.Equals, i)
		c.Assert(q.Len(), gc.Equals, length-1)
		i++
		j--
	}
	c.Assert(q.Len(), gc.Equals, 0)

	v := q.Poll()
	c.Assert(v, gc.IsNil)
	c.Assert(q.Len(), gc.Equals, 0)

	v = q.Poll()
	c.Assert(v, gc.IsNil)
	c.Assert(q.Len(), gc.Equals, 0)

	for i := 0; i < 20; i++ {
		v = q.Poll()
	}
	c.Assert(v, gc.IsNil)
	c.Assert(q.Len(), gc.Equals, 0)
}

//--Peeks tests

func (s *Suite) TestReadStuf(c *gc.C) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	c.Assert(q.Len(), gc.Equals, 1)
	val := q.Peek()
	c.Assert(val, gc.Equals, expected)
	c.Assert(q.Len(), gc.Equals, 1)
}

//--Concurrent tests

func (s *Suite) TestConcurrent(c *gc.C) {
	q := NewQueue()
	sleepTime := 100
	numberGoRoutines := 50
	numberOfPushes := 10000
	ch := make(chan int, numberGoRoutines)
	for i := 0; i < numberGoRoutines; i++ {
		go inceremtQueue(q, ch, numberOfPushes)
	}
	Wait(sleepTime)
	for {
		if len(ch) == numberGoRoutines {
			c.Assert(len(ch), gc.Equals, numberGoRoutines)
			c.Assert(q.Len(), gc.Equals, numberGoRoutines*numberOfPushes)
			return
		}
	}
}

func inceremtQueue(q *Queue, ch chan<- int, numberOfPushes int) {
	for j := 0; j < numberOfPushes; j++ {
		q.Push(j)
	}
	ch <- 1
}

func Wait(duration int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
}
