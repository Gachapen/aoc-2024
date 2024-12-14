package pq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	pq := MakePriorityQueue[string, int]()

	pq.PushItem("second", 2)
	pq.PushItem("first", 1)
	pq.PushItem("third", 3)

	first := pq.PopItem()
	second := pq.PopItem()
	third := pq.PopItem()

	assert.Equal(t, "first", first.Value)
	assert.Equal(t, "second", second.Value)
	assert.Equal(t, "third", third.Value)
}
