package gds

import (
	"maps"
	"slices"
)

type Set[T comparable] struct {
	set  map[T]bool
	list []T
}

func NewSet[T comparable](values ...T) *Set[T] {
	set := &Set[T]{
		set:  map[T]bool{},
		list: make([]T, 0, len(values)),
	}

	for _, value := range values {
		set.Add(value)
	}

	return set
}

func (s *Set[T]) Add(val T) {
	_, exists := s.set[val]
	if exists {
		return
	}

	s.list = append(s.list, val)
	s.set[val] = true
}

func (s *Set[T]) List() []T {
	return s.list
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Set[T]) IsNotEmpty() bool {
	return len(s.list) > 0
}

func (s *Set[T]) Has(value T) bool {
	return s.set[value]
}

func (s *Set[T]) Len() int {
	return len(s.list)
}

func (s *Set[T]) Merge(that *Set[T]) *Set[T] {
	newSet := s.Clone()

	for _, item := range that.list {
		newSet.Add(item)
	}

	return newSet
}

func (s *Set[T]) Clone() *Set[T] {
	return &Set[T]{
		set:  maps.Clone(s.set),
		list: slices.Clone(s.list),
	}
}

func (s *Set[T]) Walk(callback func(item T) bool) {
	for _, item := range s.list {
		continueWalk := callback(item)
		if !continueWalk {
			return
		}
	}
}

func (s *Set[T]) Equal(that *Set[T]) bool {
	if s.Len() != that.Len() {
		return false
	}

	for item := range s.set {
		if !that.Has(item) {
			return false
		}
	}

	return true
}
