package genericheap_test

import (
	"container/heap"
	"math/rand"
	"sort"
	"testing"

	"github.com/cpustejovksy/genericheap"
)

func TestHeapNew(t *testing.T) {
	maxHeapProperty := func(parent, child int) bool {
		return parent > child
	}
	nums := []int{}
	for range 100 {
		nums = append(nums, rand.Intn(4096))
	}
	mh := genericheap.New(nums, maxHeapProperty)
	largest := -1
	for mh.Len() > 0 {
		v, err := mh.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if v < largest {
			t.Fatalf("expected %d to be largest element so far", v)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	var nums []int
	maxHeapProperty := func(parent, child int) bool {
		return parent > child
	}
	nums = []int{}
	mh := genericheap.New(nums, maxHeapProperty)
	for range 100 {
		nums = append(nums, rand.Intn(4096))
	}

	t.Run("Push", func(t *testing.T) {
		for _, num := range nums {
			mh.Push(num)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		sort.Ints(nums)
		for i := len(nums) - 1; i > 0; i-- {
			num := nums[i]
			v, err := mh.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if v != num {
				t.Errorf("got %d\texpected %d\n", v, num)
			}
		}
	})
}
func TestMinHeap(t *testing.T) {
	var nums []int
	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	nums = []int{}
	mh := genericheap.New(nums, minHeapProperty)
	for range 100 {
		nums = append(nums, rand.Intn(4096))
	}

	t.Run("Push", func(t *testing.T) {
		for _, num := range nums {
			mh.Push(num)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		sort.Ints(nums)
		for _, num := range nums {
			v, err := mh.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if v != num {
				t.Errorf("got %d\texpected %d\n", v, num)
			}
		}
	})
}

func BenchmarkGenericMinHeap(b *testing.B) {
	var nums []int
	for range 1000 {
		nums = append(nums, rand.Intn(4096))
	}

	minHeapProperty := func(parent, child int) bool {
		return parent < child
	}
	h := genericheap.New([]int{}, minHeapProperty)
	for range b.N {
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
	var nums []int
	for range 1000 {
		nums = append(nums, rand.Intn(4096))
	}
	h := &IntHeap{}
	heap.Init(h)
	for range b.N {
		for _, num := range nums {
			heap.Push(h, num)
		}
		for h.Len() > 0 {
			heap.Pop(h)
		}
	}
}
