package goheap_test

import (
	"testing"

	"github.com/prantoran/goheap"
)

func TestMinHeap(t *testing.T) {
	t.Logf("en testminheap")
	h := goheap.NewMinHeap(11)

	h.Insert(1, 11)
	h.Insert(2, 3)
	h.Insert(3, 2)
	h.Insert(4, 1)
	h.Insert(9, 1)
	h.Insert(5, 15)
	h.Insert(6, 5)
	h.Insert(7, 4)
	h.Insert(8, 45)

	id, w, err := h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	h.Insert(8, 1)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	err = h.Insert(3, 2)
	t.Log("err for pushing extracted id:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

	id, w, err = h.ExtractMin()
	t.Log("id:", id, "\tw:", w, "\terr:", err)

}
