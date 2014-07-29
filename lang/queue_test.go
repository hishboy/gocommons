package lang

import (
	logLib "log"
	"os"
	"testing"
	//"reflect"

	gc "gopkg.in/check.v1" //gocheck testing framework
)

var log = logLib.New(os.Stderr, "-->: ", logLib.LstdFlags|logLib.Lshortfile)

// Hook up gc (gocheck)  into the "go test" runner.
func Test(t *testing.T) { gc.TestingT(t) }

type Suite struct{}

var _ = gc.Suite(&Suite{})

// ---- Test ----

//Test for correct inputs

func (s *Suite) TestPulltLengthZero(c *gc.C) {
	q := NewQueue()
	c.Assert(q.Len(), gc.Equals, 0)
}

func (s *Suite) TestPullNothing(c *gc.C) {
	q := NewQueue()
	val := q.Poll()
	c.Assert(val, gc.IsNil)
	c.Assert(q.Len(), gc.Equals, 0)
}

func (s *Suite) TestPutStuffGetStuffBack(c *gc.C) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	c.Assert(q.Len(), gc.Equals, 1)
	val := q.Poll()
	c.Assert(val, gc.Equals, expected)
	c.Assert(q.Len(), gc.Equals, 0)
}

func (s *Suite) TestReadStuf(c *gc.C) {
	q := NewQueue()
	expected := 20
	q.Push(expected)
	c.Assert(q.Len(), gc.Equals, 1)
	val := q.Peek()
	c.Assert(val, gc.Equals, expected)
	c.Assert(q.Len(), gc.Equals, 1)
}
