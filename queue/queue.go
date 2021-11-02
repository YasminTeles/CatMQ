package queue

import (
	"errors"
)

var errEmptyQueue = errors.New("empty queue")

type Queue struct {
	nodes []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Push(node interface{}) {
	queue.nodes = append(queue.nodes, node)
}

func (queue *Queue) Pop() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errEmptyQueue
	}

	node := queue.nodes[0]
	queue.nodes = queue.nodes[1:]

	return node, nil
}

func (queue *Queue) IsEmpty() bool {
	return queue.len() == 0
}

func (queue *Queue) len() int {
	return len(queue.nodes)
}
