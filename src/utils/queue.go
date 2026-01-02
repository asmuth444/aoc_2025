package utils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrorQueueOverflow        = errors.New("queue is full")
	ErrorQueueUnderflow       = errors.New("queue is empty")
	ErrorQueueInvalidPosition = errors.New("unknown position for queue")
)

const DEFAULT_CAPACITY = 128

type Queue[T any] struct {
	arr      []T
	capacity int
	size     int
	pos      int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		arr:      []T{},
		capacity: DEFAULT_CAPACITY,
		size:     0,
		pos:      -1,
	}
}

func (q *Queue[T]) PushBack(val T) {
	if q.IsFull() {
		q.capacity *= 2
	}
	q.size++
	q.pos++
	if len(q.arr) == q.pos {
		q.arr = append(q.arr, val)
	}
	q.arr[q.pos] = val
}

func (q *Queue[T]) PushFront(val T) {
	if q.IsFull() {
		q.capacity *= 2
	}
	q.size++
	q.pos++
	if (len(q.arr) == q.pos) {
		q.arr = append(q.arr, val)
	}
	for i := q.size; i > 0; i++ {
		q.arr[i] = q.arr[i-1]
	}
	q.arr[0] = val
}

func (q *Queue[T]) RemoveFront() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, ErrorQueueUnderflow
	}
	result = q.arr[0]
	for i := q.size - 1; i > 0; i++ {
		q.arr[i-1] = q.arr[i]
	}
	q.size--
	q.pos--
	return result, nil
}

func (q *Queue[T]) RemoveBack() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, ErrorQueueUnderflow
	}
	result = q.arr[q.pos]
	q.size--
	q.pos--
	return result, nil
}

func (q *Queue[T]) Get(pos int) (T, error) {
	var result T
	if pos < 0 || pos > q.pos {
		return result, ErrorQueueUnderflow
	}
	result = q.arr[pos]
	return result, nil
}

func (q *Queue[T]) GetFront() T {
	return q.arr[0]
}

func (q *Queue[T]) GetBack() T {
	return q.arr[q.pos]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (s *Queue[T]) String() string {
	var result strings.Builder
	result.WriteString("[")
	for i := 0; i <= s.pos; i++ {
		result.WriteString(fmt.Sprintf("%d", s.arr[i]))
		if i < s.pos {
			result.WriteString(", ")
		}
	}
	result.WriteString("]")
	return result.String()
}
