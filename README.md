# generic heap

Binary Heaps using Generics

## Rationale

I realized heaps using generics would be faster than heaps using the standard library because those heaps use interfaces and reflection.

## How to use

To create a heap with this package, you need a backing array of type that satisfies `cmp.Ordered` and a `HeapProperty` function.

The `HeapProperty` function takes two elements of the same type as your back array and returns a boolean.

You should think of it as returning the relationship between parent and child that you want this heap to maintain.

For a min-heap, the function would look like this:
```go
func(parent, child int) bool {
    return parent < child
}
```
For a max-heap, the function would look like this:
```go
func(parent, child int) bool {
    return parent < child
}
```
