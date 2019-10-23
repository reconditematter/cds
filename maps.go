package cds

import (
	"sync"
)

// MapStrStr -- a functional relation that maps strings to strings.
type MapStrStr struct {
	rel map[string]string
	mux sync.RWMutex
}

// NewMapStrStr -- returns a new empty map.
func NewMapStrStr() *MapStrStr {
	return &MapStrStr{rel: map[string]string{}}
}

// Card -- returns the cardinality of `m`.
func (m *MapStrStr) Card() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.rel)
}

// Erase -- removes all the pairs (k,v) from `m`.
func (m *MapStrStr) Erase() {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.rel = map[string]string{}
}

// Get -- if a pair (k,v) is in `m`, then returns (v,true);
// otherwise, returns ("",false).
func (m *MapStrStr) Get(k string) (v string, ok bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	v, ok = m.rel[k]
	return
}

// Take -- if a pair (k,v) is in `m`, then removes it and returns (v,true);
// otherwise, returns ("",false).
func (m *MapStrStr) Take(k string) (v string, ok bool) {
	m.mux.Lock()
	defer m.mux.Unlock()
	v, ok = m.rel[k]
	if ok {
		delete(m.rel, k)
	}
	return
}

// Put -- puts the pair (k,v) into `m`.
func (m *MapStrStr) Put(k, v string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.rel[k] = v
}

// Dom -- returns the domain of `m`
// (a set of `k` from the all the pairs (k,v) in `m`).
//
// It is true that m.Dom().Card() = m.Card().
func (m *MapStrStr) Dom() *SetOfStr {
	m.mux.RLock()
	defer m.mux.RUnlock()
	s := NewSetOfStr()
	for k := range m.rel {
		s.Extend(k)
	}
	return s
}

// Ran -- returns the range of `m`
// (a set of `v` from the all the pairs (k,v) in `m`).
//
// It is true that m.Ran().Card() ≤ m.Card().
func (m *MapStrStr) Ran() *SetOfStr {
	m.mux.RLock()
	defer m.mux.RUnlock()
	s := NewSetOfStr()
	for _, v := range m.rel {
		s.Extend(v)
	}
	return s
}

// List -- returns all the pairs (k,v) of `m`.
func (m *MapStrStr) List() []struct {
	K string
	V string
} {
	m.mux.RLock()
	defer m.mux.RUnlock()
	a := make([]struct {
		K string
		V string
	}, len(m.rel))
	i := 0
	for k, v := range m.rel {
		a[i] = struct {
			K string
			V string
		}{k, v}
		i++
	}
	return a
}

// MapStrInt -- a functional relation that maps strings to ints.
type MapStrInt struct {
	rel map[string]int
	mux sync.RWMutex
}

// NewMapStrInt -- returns a new empty map.
func NewMapStrInt() *MapStrInt {
	return &MapStrInt{rel: map[string]int{}}
}

// Card -- returns the cardinality of `m`.
func (m *MapStrInt) Card() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.rel)
}

// Erase -- removes all the pairs (k,v) from `m`.
func (m *MapStrInt) Erase() {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.rel = map[string]int{}
}

// Get -- if a pair (k,v) is in `m`, then returns (v,true);
// otherwise, returns (0,false).
func (m *MapStrInt) Get(k string) (v int, ok bool) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	v, ok = m.rel[k]
	return
}

// Take -- if a pair (k,v) is in `m`, then removes it and returns (v,true);
// otherwise, returns (0,false).
func (m *MapStrInt) Take(k string) (v int, ok bool) {
	m.mux.Lock()
	defer m.mux.Unlock()
	v, ok = m.rel[k]
	if ok {
		delete(m.rel, k)
	}
	return
}

// Put -- puts the pair (k,v) into `m`.
func (m *MapStrInt) Put(k string, v int) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.rel[k] = v
}

// Dom -- returns the domain of `m`
// (a set of `k` from all the pairs (k,v) in `m`).
//
// It is true that m.Dom().Card() = m.Card().
func (m *MapStrInt) Dom() *SetOfStr {
	m.mux.RLock()
	defer m.mux.RUnlock()
	s := NewSetOfStr()
	for k := range m.rel {
		s.Extend(k)
	}
	return s
}

// Ran -- returns the range of `m`
// (a set of `v` from the all pairs (k,v) in `m`).
//
// It is true that m.Ran().Card() ≤ m.Card().
func (m *MapStrInt) Ran() *SetOfInt {
	m.mux.RLock()
	defer m.mux.RUnlock()
	s := NewSetOfInt()
	for _, v := range m.rel {
		s.Extend(v)
	}
	return s
}

// List -- returns all the pairs (k,v) of `m`.
func (m *MapStrInt) List() []struct {
	K string
	V int
} {
	m.mux.RLock()
	defer m.mux.RUnlock()
	a := make([]struct {
		K string
		V int
	}, len(m.rel))
	i := 0
	for k, v := range m.rel {
		a[i] = struct {
			K string
			V int
		}{k, v}
		i++
	}
	return a
}
