package main

import (
	"bytes"
	"fmt"
)

const size = 32 << (^uint(0) >> 63)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))
}

type IntSet struct {
	words []uint64
	count int
}

func (s *IntSet) Len() int {
	return s.count
}

func (s *IntSet) Elems() []int {
	if s.Len() == 0 {
		return nil
	}
	r := make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				r = append(r, size*i+j)
			}
		}
	}
	return r
}
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) AddAll(list ...int) {
	for _, i := range list {
		s.Add(i)
	}
}

func (s *IntSet) Add(x int) {
	word, bit := x/size, x%size
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	s.count++
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/size, uint(x%size)
		s.words[word] &^= 1 << bit
		s.count--
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith return intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	r := NewIntSet()
	for _, i := range t.Elems() {
		if s.Has(i) {
			r.Add(i)
		}
	}
	return r
}

// DifferenceWith return difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	r := NewIntSet()
	for _, i := range s.Elems() {
		if !t.Has(i) {
			r.Add(i)
		}
	}
	return r
}

// SymmetricDifference return symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	u := s.Copy()
	u.UnionWith(t)

	return u.DifferenceWith(s.IntersectWith(t))
}
func (s *IntSet) Copy() *IntSet {
	c := NewIntSet()
	c.words = make([]uint64, len(s.words))
	copy(c.words, s.words)
	c.count = s.count
	return c
}

func NewIntSet() *IntSet {
	t := IntSet{nil, 0}
	return &t
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
