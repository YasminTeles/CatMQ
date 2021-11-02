package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()

	queue.Push("Fernando")

	assert.NotEmpty(t, queue)
}

func TestPopQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()

	queue.Push("Alejandro")
	queue.Push("Diego")

	node, err := queue.Pop()

	assert.Equal(t, "Alejandro", node)
	assert.NoError(t, err)
}

func TestPopEmptyQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()

	node, err := queue.Pop()

	assert.Nil(t, node)
	assert.Error(t, err)
}

func TestIsEmptyQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()

	isEmpty := queue.IsEmpty()

	assert.True(t, isEmpty)
}

func TestIsNotEmptyQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()
	queue.Push("Diego")

	isEmpty := queue.IsEmpty()

	assert.False(t, isEmpty)
}

func TestLenQueue(t *testing.T) {
	t.Parallel()

	queue := NewQueue()
	queue.Push("Alejandro")

	length := queue.len()

	assert.Equal(t, 1, length)
}
