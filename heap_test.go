package genericheap_test

import (
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
