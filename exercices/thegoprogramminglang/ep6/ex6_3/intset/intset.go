package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

// ex6.3
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	max_len, min_len := max(len(s.words), len(t.words)), min(len(s.words), len(t.words))
	the_result_words := make([]uint64, max_len)
	for index := range the_result_words {
		if index < min_len {
			the_result_words[index] = s.words[index] & t.words[index]
		} else {
			the_result_words[index] = 0
		}
	}
	return &IntSet{words: the_result_words}
}
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	max_len, min_len := max(len(s.words), len(t.words)), min(len(s.words), len(t.words))
	the_result_words := make([]uint64, max_len)
	for index := range the_result_words {
		if index < min_len {
			the_result_words[index] = s.words[index] &^ t.words[index]
		} else {
			if max_len == len(s.words) {
				the_result_words[index] = s.words[index]
			} else {
				the_result_words[index] = 0
			}
		}
	}
	return &IntSet{words: the_result_words}
}
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	max_len, min_len := max(len(s.words), len(t.words)), min(len(s.words), len(t.words))
	the_result_words := make([]uint64, max_len)
	for index := range the_result_words {
		if index < min_len {
			the_result_words[index] = s.words[index] ^ t.words[index]
		} else {
			if max_len == len(s.words) {
				the_result_words[index] = s.words[index]
			} else {
				the_result_words[index] = t.words[index]
			}
		}
	}
	return &IntSet{words: the_result_words}
}
func (s *IntSet) Elems() []int64 {
	result := make([]int64, 0)
	for index := range s.words {
		for i := 0; i < 64; i++ {
			if (1<<i)&s.words[index] != 0 {
				result = append(result, int64(64*index+i))
			}
		}
	}
	return result
}
