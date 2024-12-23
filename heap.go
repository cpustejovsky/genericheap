package genericheap

import "iter"

type Heap[T any] struct {
	array        []T
	heapProperty HeapProperty[T]
}
type HeapProperty[T any] func(T, T) bool

type EmptyHeapError struct{}

func (e *EmptyHeapError) Error() string {
	return "empty heap"
}

func New[T any](arr []T, fn HeapProperty[T]) *Heap[T] {
	gh := Heap[T]{
		array:        []T{},
		heapProperty: fn,
	}
	for _, v := range arr {
		gh.Push(v)
	}
	return &gh
}

func (h *Heap[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for h.Len() > 0 {
			v, err := h.Pop()
			if !yield(v) || err != nil {
				return
			}
		}
	}
}

func (h *Heap[T]) Len() int {
	return len(h.array)
}

func (h *Heap[T]) Push(val T) {
	h.array = append(h.array, val)
	h.up()
}

func (h *Heap[T]) up() {
	for cur := h.Len() - 1; cur > 0; {
		parent := (cur - 1) / 2
		if !h.heapProperty(h.array[parent], h.array[cur]) {
			h.swap(parent, cur)
		}
		cur = parent
	}
}

func (h *Heap[T]) Peak() (T, error) {
	var r T
	if h.Len() == 0 {
		return r, &EmptyHeapError{}
	}
	r = h.array[0]
	return r, nil
}

func (h *Heap[T]) Pop() (T, error) {
	r, err := h.Peak()
	if err != nil {
		return r, err
	}
	h.array[0] = h.array[h.Len()-1]
	h.array = h.array[:h.Len()-1]
	h.down()
	return r, nil
}

func (h *Heap[T]) down() {
	for cur, target := 0, 0; ; cur = target {
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
	}
}

func (h *Heap[T]) PushPop(val T) T {
	if h.Len() == 0 || h.heapProperty(val, h.array[0]) {
		return val
	}
	val, h.array[0] = h.array[0], val
	h.down()
	return val
}

func (h *Heap[T]) swap(x, y int) {
	h.array[x], h.array[y] = h.array[y], h.array[x]
}
