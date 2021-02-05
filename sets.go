// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package cds

import (
	"sync"
)

// SetOfInt -- a set of elements of the type int.
type SetOfInt struct {
	set map[int]struct{}
	mux sync.RWMutex
}

// NewSetOfInt -- returns a new empty set.
func NewSetOfInt() *SetOfInt {
	return &SetOfInt{set: map[int]struct{}{}}
}

// Card -- returns the cardinality of `s`.
func (s *SetOfInt) Card() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.set)
}

// Erase -- removes all the elements from `s`.
func (s *SetOfInt) Erase() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.set = map[int]struct{}{}
}

// Has -- returns true iff `s` contains `x`.
func (s *SetOfInt) Has(x int) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	_, ok := s.set[x]
	return ok
}

// Extend -- extends `s` with `x`.
func (s *SetOfInt) Extend(x int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.set[x] = struct{}{}
}

// Remove -- removes `x` from `s`.
func (s *SetOfInt) Remove(x int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.set, x)
}

// List -- returns all the elements of `s`.
func (s *SetOfInt) List() []int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	a := make([]int, len(s.set))
	i := 0
	for k, _ := range s.set {
		a[i] = k
		i++
	}
	return a
}

// SetOfStr -- a set of elements of the type string.
type SetOfStr struct {
	set map[string]struct{}
	mux sync.RWMutex
}

// NewSetOfStr -- returns a new empty set.
func NewSetOfStr() *SetOfStr {
	return &SetOfStr{set: map[string]struct{}{}}
}

// Card -- returns the cardinality of `s`.
func (s *SetOfStr) Card() int {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return len(s.set)
}

// Erase -- removes all the elements from `s`.
func (s *SetOfStr) Erase() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.set = map[string]struct{}{}
}

// Has -- returns true iff `s` contains `x`.
func (s *SetOfStr) Has(x string) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	_, ok := s.set[x]
	return ok
}

// Extend -- extends `s` with `x`.
func (s *SetOfStr) Extend(x string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.set[x] = struct{}{}
}

// Remove -- removes `x` from `s`.
func (s *SetOfStr) Remove(x string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.set, x)
}

// List -- returns all the elements of `s`.
func (s *SetOfStr) List() []string {
	s.mux.RLock()
	defer s.mux.RUnlock()
	a := make([]string, len(s.set))
	i := 0
	for k, _ := range s.set {
		a[i] = k
		i++
	}
	return a
}
