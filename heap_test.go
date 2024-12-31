package genericheap_test

import (
	"container/heap"
	"errors"
	"math/rand"
	"testing"

	"github.com/cpustejovksy/genericheap"
)

func RandomNumbers(size int) []int {
	nums := make([]int, size)
	for i := range size {
		nums[i] = rand.Intn(4096)
	}
	return nums
}

func TestHeapNew(t *testing.T) {
	t.Run("MaxHeap", func(t *testing.T) {
		maxHeapProperty := func(parent, child int) bool {
			return parent > child
		}
		nums := RandomNumbers(100)
		mh := genericheap.New(nums, maxHeapProperty)
		largest := 5000
		for mh.Len() > 0 {
			v, _ := mh.Pop()
			if v > largest {
				t.Fatalf("failed")
			}
			largest = v

		}
	})
	t.Run("MinHeap", func(t *testing.T) {
		minHeapProperty := func(parent, child int) bool {
			return parent < child
		}
		nums := RandomNumbers(100)
		mh := genericheap.New(nums, minHeapProperty)
		smallest := -5000
		for mh.Len() > 0 {
			v, _ := mh.Pop()
			if v < smallest {
				t.Fatalf("failed")
			}
			smallest = v
		}
	})
}

func TestMaxHeap(t *testing.T) {
	maxHeapProperty := func(parent, child int) bool {
		return parent > child
	}
	mh := genericheap.New([]int{}, maxHeapProperty)
	nums := RandomNumbers(100)
	t.Run("Push", func(t *testing.T) {
		for _, num := range nums {
			mh.Push(num)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		largest := 5000
		for mh.Len() > 0 {
			v, _ := mh.Pop()
			if v > largest {
				t.Fatalf("failed")
			}
			largest = v

		}
		_, err := mh.Pop()
		var check *genericheap.EmptyHeapError
		if !errors.As(err, &check) {
			t.Fatalf("error of type %T should of type %T", err, &check)
		}
	})
}
func TestMinHeap(t *testing.T) {
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	mh := genericheap.New([]int{}, minHeapProperty)
	nums := RandomNumbers(100)
	t.Run("Push", func(t *testing.T) {
		for _, num := range nums {
			mh.Push(num)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		smallest := -5000
		for mh.Len() > 0 {
			v, _ := mh.Pop()
			if v < smallest {
				t.Fatalf("failed")
			}
			smallest = v
		}
		_, err := mh.Pop()
		var check *genericheap.EmptyHeapError
		if !errors.As(err, &check) {
			t.Fatalf("error of type %T should of type %T", err, &check)
		}
	})
}

func TestHeapAll(t *testing.T) {
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	nums := RandomNumbers(10)
	mh := genericheap.New(nums, minHeapProperty)
	smallest := -5000
	for v := range mh.All() {
		if v < smallest {
			t.Fatalf("failed")
		}
		smallest = v
	}
}

func TestHeapPushPop(t *testing.T) {
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	nums := RandomNumbers(100)
	mh := genericheap.New(nums, minHeapProperty)

	biggerNum := 6000
	smallest := -5000
	for range nums {
		v := mh.PushPop(biggerNum)
		if v < smallest {
			t.Fatal("failed")
		}
		smallest = v
		biggerNum--
	}
	for v := range mh.All() {
		if v < smallest {
			t.Fatal("failed")
		}
		smallest = v
	}
}

var result int

func BenchmarkPushThenPop(b *testing.B) {
	nums1 := RandomNumbers(1000)
	nums2 := RandomNumbers(1000)
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	h := genericheap.New(nums1, minHeapProperty)
	for b.Loop() {
		for _, num := range nums2 {
			h.Push(num)
			h.Pop()
		}
	}
}
func BenchmarkPushPop(b *testing.B) {
	nums1 := RandomNumbers(1000)
	nums2 := RandomNumbers(1000)
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	h := genericheap.New(nums1, minHeapProperty)
	for b.Loop() {
		for _, num := range nums2 {
			h.PushPop(num)
		}
	}
}

func BenchmarkGenericMinHeap(b *testing.B) {
	nums := RandomNumbers(1000)
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	h := genericheap.New([]int{}, minHeapProperty)
	for b.Loop() {
		for _, num := range nums {
			h.Push(num)
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkContainersMinHeap(b *testing.B) {
	nums := RandomNumbers(1000)
	h := &IntHeap{}
	heap.Init(h)
	for b.Loop() {
		for _, num := range nums {
			heap.Push(h, num)
		}
		for h.Len() > 0 {
			heap.Pop(h)
		}
	}
}
