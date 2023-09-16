package vec

import (
	"fmt"
)

type Pair[A comparable, B comparable] struct {
	a A
	b B
}

func NewVec[T comparable]() *Vec[T] {
	return &Vec[T]{
		data: make([]T, 0, 3),
	}
}
func FromSlice[T comparable](s []T) *Vec[T] {
	v := make([]T, len(s), len(s))
	copy(v, s)
	return &Vec[T]{
		data: v,
	}
}

type Vec[T comparable] struct {
	data []T
}
type Iterator[T comparable] struct {
	data []T
}

func (i *Iterator[T]) Count() int {
	return len(i.data)
}
func (i *Iterator[T]) Filter(f func(v T) bool) *Iterator[T] {
	list := make([]T, 0, len(i.data))
	for _, d := range i.data {
		if f(d) {
			list = append(list, d)
		}
	}
	return &Iterator[T]{
		data: list,
	}
}
func (i *Iterator[T]) Map(f func(v T) T) *Iterator[T] {
	list := make([]T, 0, len(i.data))
	for _, d := range i.data {
		list = append(list, f(d))
	}
	return &Iterator[T]{
		data: list,
	}
}

func (i *Iterator[T]) ForEach(f func(v T)) {
	for _, d := range i.data {
		f(d)
	}
}
func (i *Iterator[T]) Collect() *Vec[T] {
	return &Vec[T]{
		data: i.data,
	}
}
func (i *Iterator[T]) Unfold(init T, f func(result *T, e T)) T {
	result := init
	for _, d := range i.data {
		f(&result, d)
	}
	return result
}
func (i *Iterator[T]) RemoveDuplicate() *Iterator[T] {
	set := make(map[T]struct{})
	for _, d := range i.data {
		set[d] = struct{}{}
	}
	newData := make([]T, 0, len(set))
	for t, _ := range set {
		newData = append(newData, t)
	}
	return &Iterator[T]{
		data: newData,
	}
}
func (v *Vec[T]) Get(i int) T {
	return v.data[i]
}
func (v *Vec[T]) Len() int {
	return len(v.data)
}
func (v *Vec[T]) Insert(i int, value T) {
	v.data = append(v.data[0:i], append([]T{value}, v.data[i:len(v.data)]...)...)
}
func (v *Vec[T]) Push(value T) {
	v.data = append(v.data, value)
}
func (v *Vec[T]) Index(value T) int {
	for i, value0 := range v.data {
		if value0 == value {
			return i
		}
	}
	return -1
}
func (v *Vec[T]) Delete(i int) {
	v.data = append(v.data[0:i], v.data[i+1:len(v.data)]...)
}
func (v *Vec[T]) Iter() *Iterator[T] {
	l := len(v.data)
	newSlice := make([]T, l, l)
	copy(newSlice, v.data)
	return &Iterator[T]{
		newSlice,
	}
}
func (i *Iterator[T]) String() string {
	return fmt.Sprint(i.data)
}
func (v *Vec[T]) String() string {
	return fmt.Sprint(v.data)
}
