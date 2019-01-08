package goheap

// MinHeap defines the MinHeap interface
type MinHeap interface {
	Insert(id int, weight float64) error
	ExtractMin() (id int, weight float64, err error)
}
