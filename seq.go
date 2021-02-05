// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package cds

import (
	"sync"
)

// Seq -- an extensible sequence of elements.
// Elements can be added or removed at either end of a sequence.
// Elements can also be accesses or updated at specified indexes.
type Seq struct {
	lock sync.RWMutex
	elem []interface{}
	orig int
	size int
}

// NewSeq -- returns a new empty sequence.
func NewSeq() *Seq {
	return &Seq{elem: make([]interface{}, 16)}
}

func (s *Seq) grow() {
	if len(s.elem) == 0 {
		// initialize `s`
		s.elem = make([]interface{}, 16)
		return
	}
	//
	n := len(s.elem)
	m := n - s.orig
	// grow factor = the golden ratio
	enew := make([]interface{}, int(float64(n)*1.61803398874989484820458683436563811772030917980576286213544862))
	copy(enew[0:], s.elem[s.orig:])
	copy(enew[m:], s.elem[:s.orig])
	s.orig = 0
	s.elem = enew
}

// Size -- returns the number of elements in `s`.
func (s *Seq) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.size
}

// Addhi -- extends `s` with the element `e` at the high-index end.
func (s *Seq) Addhi(e interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.size == len(s.elem) {
		s.grow()
	}
	//
	i := s.orig + s.size
	if i >= len(s.elem) {
		i -= len(s.elem)
	}
	s.elem[i] = e
	s.size++
}

// Addlo -- extends `s` with the element `e` at the low-index end.
func (s *Seq) Addlo(e interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.size == len(s.elem) {
		s.grow()
	}
	//
	i := s.orig
	if i == 0 {
		i = len(s.elem) - 1
	} else {
		i--
	}
	s.elem[i] = e
	s.size++
	s.orig = i
}

// Pophi -- removes and returns the element of `s` at the high-index end.
// This method causes a runtime panic when `s` is empty.
func (s *Seq) Pophi() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !(s.size > 0) {
		panic("cds.Seq.Pophi: empty")
	}
	i := s.orig + s.size - 1
	if i >= len(s.elem) {
		i -= len(s.elem)
	}
	res := s.elem[i]
	s.elem[i] = nil
	s.size--
	return res
}

// Poplo -- removes and returns the element of `s` at the low-index end.
// This method causes a runtime panic when `s` is empty.
func (s *Seq) Poplo() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !(s.size > 0) {
		panic("cds.Seq.Poplo: empty")
	}
	res := s.elem[s.orig]
	s.elem[s.orig] = nil
	s.size--
	s.orig++
	if s.orig == len(s.elem) {
		s.orig = 0
	}
	return res
}

// Puthi -- replaces the element of `s` at the high-index end with `e`.
// This method causes a runtime panic when `s` is empty.
func (s *Seq) Puthi(e interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !(s.size > 0) {
		panic("cds.Seq.Puthi: empty")
	}
	i := s.orig + s.size - 1
	if i >= len(s.elem) {
		i -= len(s.elem)
	}
	s.elem[i] = e
}

// Putlo -- replaces the element of `s` at the low-index end with `e`.
// This method causes a runtime panic when `s` is empty.
func (s *Seq) Putlo(e interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !(s.size > 0) {
		panic("cds.Seq.Putlo: empty")
	}
	s.elem[s.orig] = e
}

// Gethi -- returns the element of `s` at the high-index end.
// This method causes a runtime panic when `s` is empty.
func (s *Seq) Gethi() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if !(s.size > 0) {
		panic("cds.Seq.Gethi: empty")
	}
	i := s.orig + s.size - 1
	if i >= len(s.elem) {
		i -= len(s.elem)
	}
	return s.elem[i]
}

// Getlo -- returns the element of `s` at the low-index end.
// This method causes a runtime panic when `s` is empty.
func (s *Seq) Getlo() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if !(s.size > 0) {
		panic("cds.Seq.Getlo: empty")
	}
	return s.elem[s.orig]
}

// Put -- replaces the element of `s` at the index `i` with `e`.
// This method causes a runtime panic when i∉[0,s.Size()-1].
func (s *Seq) Put(i int, e interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !(0 <= i && i < s.size) {
		panic("cds.Seq.Put: index")
	}
	j := s.orig + i
	if j >= len(s.elem) {
		j -= len(s.elem)
	}
	s.elem[j] = e
}

// Get -- returns the element of `s` at the index `i`.
// This method causes a runtime panic when i∉[0,s.Size()-1].
func (s *Seq) Get(i int) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if !(0 <= i && i < s.size) {
		panic("cds.Seq.Get: index")
	}
	j := s.orig + i
	if j >= len(s.elem) {
		j -= len(s.elem)
	}
	return s.elem[j]
}

// Swap -- exchanges the elements of `s` at the indexes `i` and `j`.
// This method causes a runtime panic when i,j∉[0,s.Size()-1].
func (s *Seq) Swap(i, j int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !(0 <= i && i < s.size && 0 <= j && j < s.size) {
		panic("cds.Seq.Swap: index")
	}
	ii := s.orig + i
	if ii >= len(s.elem) {
		ii -= len(s.elem)
	}
	jj := s.orig + j
	if jj >= len(s.elem) {
		jj -= len(s.elem)
	}
	s.elem[ii], s.elem[jj] = s.elem[jj], s.elem[ii]
}

// Export -- returns all the elements of `s` as a slice.
func (s *Seq) Export() []interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	t := make([]interface{}, s.size)
	for i := range t {
		j := s.orig + i
		if j >= len(s.elem) {
			j -= len(s.elem)
		}
		t[i] = s.elem[j]
	}
	return t
}
