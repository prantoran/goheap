# goheap

- Min Heap implementation in Go
- Support
    - Insert
    - ExtractMin
    - Reseting key weight with lower value
        - Insert the previously inserted key with lower weight


### How to use `MinHeap`:

* First create a new instance:
```
	h := goheap.NewMinHeap(maxCap)
```
Here, `maxCap` is the maximum number of items to be stored in the heap

* Insert an item index and the weight of the item:
```
	h.Insert(itemIndexID, itemWeight)
```
Here, `itemIndexID` is the index of the item and `itemWeight` is the weight of the item. The lower the weight, the closer the itemID will be to the root of the min heap. 

If an itemIndexID already exists in the heap and inserted again, then the latest weight is updated in the min heap and the position of `itemIndexID` in the heap is adjusted.

* Extract Min itemID with the lowest weight in the min heap:
```
	id, w, err := h.ExtractMin()
```
Here, `id` and `w` is the index of the item and weight respectively at the top of the min heap.