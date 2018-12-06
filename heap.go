package goheap

import (
	"errors"
)

type heapItem struct {
	id int
	w  float64
}

type minHeap struct {
	items         []heapItem
	itemIndex     map[int]int
	vis           []bool
	cap, heapSize int
}

// p = parent
func (h *minHeap) p(i int) int {
	return (i - 1) / 2
}

// lc = left child
func (h *minHeap) lc(i int) int {
	return (2 * i) + 1
}

// rc = right child
func (h *minHeap) rc(i int) int {
	return (2 * i) + 2
}

func (h *minHeap) updateExisting(id int, weight float64) bool {
	if !h.vis[id] {
		return false
	}

	i := h.itemIndex[id]

	if h.items[i].w <= weight {
		return true
	}

	h.updateW(i, weight)

	return true
}

func (h *minHeap) updateW(i int, newWeight float64) {
	h.items[i].w = newWeight

	for i != 0 && h.items[h.p(i)].w > h.items[i].w {
		h.swap(i, h.p(i))
		i = h.p(i)
	}
}

func (h *minHeap) Insert(id int, weight float64) error {

	if h.updateExisting(id, weight) {
		return nil
	}

	h.vis[id] = true

	if h.cap == h.heapSize {
		return errors.New("heap full cap")
	}

	i := h.heapSize
	h.heapSize++

	h.items[i] = heapItem{id: id, w: weight}
	h.itemIndex[id] = i

	for i != 0 && h.items[h.p(i)].w > h.items[i].w {
		h.swap(i, h.p(i))
		i = h.p(i)
	}

	return nil
}

func (h *minHeap) deleteItemIndex(id int) {
	_, ok := h.itemIndex[id]
	if ok {
		delete(h.itemIndex, id)
	}
}

func (h *minHeap) ExtractMin() (id int, weight float64, err error) {
	if h.heapSize <= 0 {
		return 0, 0, errors.New("no item in heap")
	}

	id = h.items[0].id
	weight = h.items[0].w

	h.vis[id] = false

	if h.heapSize == 1 {
		h.heapSize--
		return id, weight, nil
	}

	root := h.items[0]

	h.swap(0, h.heapSize-1)

	h.heapSize--

	h.deleteItemIndex(id)

	h.minHeapify(0)

	return root.id, root.w, nil
}

// swap over heap array indexes and updates itemIndex id ref to heap array index
func (h *minHeap) swap(i, j int) {
	t := h.items[i]
	u := t.id
	v := h.items[j].id

	h.items[i] = h.items[j]
	h.items[j] = t

	tt := h.itemIndex[u]
	h.itemIndex[u] = h.itemIndex[v]
	h.itemIndex[v] = tt
}

func (h *minHeap) minHeapify(i int) {
	l := h.lc(i)
	r := h.rc(i)
	ret := i

	if l < h.heapSize && h.items[l].w < h.items[ret].w {
		ret = l
	}

	if r < h.heapSize && h.items[r].w < h.items[ret].w {
		ret = r
	}

	if ret != i {
		h.swap(i, ret)
		h.minHeapify(ret)
	}
}

func NewMinHeap(cap int) MinHeap {
	return &minHeap{
		cap:       cap,
		items:     make([]heapItem, cap),
		vis:       make([]bool, cap),
		itemIndex: map[int]int{},
	}
}
