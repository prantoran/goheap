package goheap

type MinHeap interface {
	Insert(id int, weight float64) error
	ExtractMin() (id int, weight float64, err error)
}
