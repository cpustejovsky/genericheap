package genericheap

import (
	"cmp"
	"errors"
)

type GenericHeap[T cmp.Ordered] struct {
	array        []T
	heapProperty HeapProperty[T]
}
type HeapProperty[T cmp.Ordered] func(T, T) bool

func New[T cmp.Ordered](arr []T, fn HeapProperty[T]) *GenericHeap[T] {
	gh := GenericHeap[T]{
		array:        arr,
		heapProperty: fn,
	}
	if gh.Len() > 0 {
		v, _ := gh.Pop()
		gh.Push(v)
	}
	return &gh
}

func (h *GenericHeap[T]) Len() int {
	return len(h.array)
}

func (h *GenericHeap[T]) Push(val T) {
	h.array = append(h.array, val)
	h.up()
}

func (h *GenericHeap[T]) up() {
	for cur := h.Len() - 1; cur > 0; {
		parent := (cur - 1) / 2
		if !h.heapProperty(h.array[parent], h.array[cur]) {
			h.swap(parent, cur)
		}
		cur = parent
	}
}

func (h *GenericHeap[T]) Peak() (T, error) {
	var r T
	if h.Len() == 0 {
		return r, errors.New("empty heap")
	}
	r = h.array[0]
	return r, nil
}

func (h *GenericHeap[T]) Pop() (T, error) {
	r, err := h.Peak()
	if err != nil {
		return r, err
	}
	h.array[0] = h.array[h.Len()-1]
	h.array = h.array[:h.Len()-1]
	h.down()
	return r, nil
}

func (h *GenericHeap[T]) down() {
	var cur, target int
	for {
		target = cur
		if left := cur*2 + 1; left < h.Len() && !h.heapProperty(h.array[target], h.array[left]) {
			target = left
		}
		if right := cur*2 + 2; right < h.Len() && !h.heapProperty(h.array[target], h.array[right]) {
			target = right
		}
		if cur == target {
			break
		}
		h.swap(cur, target)
		cur = target

	}
}

func (h *GenericHeap[T]) swap(x, y int) {
	h.array[x], h.array[y] = h.array[y], h.array[x]
}
