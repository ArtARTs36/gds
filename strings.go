package gds

import (
	"slices"
	"strings"
)

type Strings struct {
	items []string
}

func NewStrings(items ...string) *Strings {
	return &Strings{
		items: items,
	}
}

func (s *Strings) Add(str string) {
	s.items = append(s.items, str)
}

func (s *Strings) Join(sep string) *String {
	return &String{
		Value: strings.Join(s.items, sep),
	}
}

func (s *Strings) Contains(str string) bool {
	return slices.Contains(s.items, str)
}

func (s *Strings) Len() int {
	return len(s.items)
}

func (s *Strings) First() string {
	if s.IsEmpty() {
		return ""
	}
	return s.items[0]
}

func (s *Strings) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Strings) IsNotEmpty() bool {
	return len(s.items) > 0
}

func (s *Strings) List() []string {
	return s.items
}
