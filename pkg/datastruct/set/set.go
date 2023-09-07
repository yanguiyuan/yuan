package set

import "fmt"

type Set[T comparable] struct {
	data map[T]struct{}
}

func (s *Set[T]) String() string {
	res := "{"
	for k, _ := range s.data {
		res += fmt.Sprint(k, ",")
	}
	res = res[0 : len(res)-1]
	res += "}"
	return res
}
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}
func (s *Set[T]) Add(v T) bool {
	if _, ok := s.data[v]; ok {
		return false
	} else {
		s.data[v] = struct{}{}
		return true
	}
}
func (s *Set[T]) Delete(v T) {
	delete(s.data, v)
}
func (s *Set[T]) Len() int {
	return len(s.data)
}
func (s *Set[T]) Contain(v T) bool {
	_, ok := s.data[v]
	return ok
}
func (s *Set[T]) ForEach(f func(v T)) {
	for k, _ := range s.data {
		f(k)
	}
}
func (s *Set[T]) Map(f func(v T) T) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		set.Add(f(k))
	}
	return set
}
func (s *Set[T]) Find(f func(v T) bool) []T {
	var list []T
	for k, _ := range s.data {
		if f(k) {
			list = append(list, k)
		}
	}
	return list
}
