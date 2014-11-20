package main

import (
    "fmt"
    "math"
)

// 堆
type Heap []int

// 最大堆调整
func (h Heap) siftDown(index int) {
    for {
        child := index*2 + 1
        if child >= len(h) {
            break
        }
        if child+1 < len(h) && h[child] < h[child+1] {
            child++
        }
        if h[index] > h[child] {
            return
        }
        h[index], h[child] = h[child], h[index]
        index = child
    }
}

// 构建大顶堆
func (h Heap) buildMaxHeap() {
    for i := (len(h) - 2) / 2; i >= 0; i-- {
        h.siftDown(i)
    }
}

// 堆排序
func (h Heap) HeapSort() {
    heap := []int{}
    length := len(h)
    var x int
    for i := 0; i < length; i++ {
        x, h = h[0], h[1:]
        heap = append(heap, x)
        h.buildMaxHeap()
    }
    h = heap
}

func (h Heap) Show() {
    deep := int(math.Ceil(math.Log2(float64(len(h) + 1))))
    start := 0
    fmt.Println("----------")
    for i := 0; i < deep; i++ {
        length := int(math.Pow(2, float64(i)))
        end := start + length
        if end > len(h) {
            end = len(h)
        }
        fmt.Println(h[start:end])
        start += length
    }
    fmt.Println("----------")
}

func main() {
    //heap := Heap{0, 1, 2, 3, 4, 5, 6}
    heap := Heap{6, 3, 1, 5, 9, 2, 7, 8, 0, 4}
    heap.Show()
    heap.buildMaxHeap()
    heap.Show()
    heap.HeapSort()
    heap.Show()
}
