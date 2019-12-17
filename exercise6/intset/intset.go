package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

const intLen = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/intLen, uint(x%intLen)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/intLen, uint(x%intLen)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
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

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intLen; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intLen*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// 要素数を返します
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for i := 0; i < intLen; i++ {
			if (word>>i)&1 == 1 {
				count++
			}
		}
	}
	return count
}

// セットからxを取り除きます
func (s *IntSet) Remove(x int) {
	word, bit := x/intLen, uint(x%intLen)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// セットからすべての要素を取り除きます
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

// セットのコピーを返します
func (s *IntSet) Copy() *IntSet {
	var ret IntSet
	ret.words = make([]uint, len(s.words), cap(s.words))
	copy(ret.words, s.words)

	return &ret
}
