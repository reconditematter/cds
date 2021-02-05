// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package cds

import (
	"sync"
)

// FIFO -- a first-in, first-out queue.
type FIFO struct {
	mux   sync.Mutex
	rel   map[uint64]interface{}
	peak  int
	first uint64
	limit uint64
}

// NewFIFO -- returns a new empty queue.
func NewFIFO() *FIFO {
	return &FIFO{rel: make(map[uint64]interface{})}
}

// Len -- returns the number of elements currently in `q`.
func (q *FIFO) Len() int {
	q.mux.Lock()
	defer q.mux.Unlock()
	return len(q.rel)
}

// Max -- returns the maximum number of elements that have been in `q`.
func (q *FIFO) Max() int {
	q.mux.Lock()
	defer q.mux.Unlock()
	return q.peak
}

// Enq -- places `e` at the tail of `q`.
func (q *FIFO) Enq(e interface{}) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.rel[q.limit] = e
	q.limit++
	if len(q.rel) > q.peak {
		q.peak = len(q.rel)
	}
}

// Deq -- if `q` is not empty, then removes and returns the element at
// the head of `q` as (e,true); otherwise, returns (nil,false).
func (q *FIFO) Deq() (e interface{}, ok bool) {
	q.mux.Lock()
	defer q.mux.Unlock()
	e, ok = q.rel[q.first]
	if ok {
		delete(q.rel, q.first)
		q.first++
	}
	return
}
