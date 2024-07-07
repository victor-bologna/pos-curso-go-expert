package main

import (
	"fmt"
	"sync"
)

type span struct {
	size      int
	allocated bool
}

type mheap struct {
	spans []*span
	lock  sync.Mutex
}

type mcentral struct {
	sizeSpans []*span
	lock      sync.Mutex
}

type mcache struct {
	localSpan []*span
}

func NewHeap(size int) *mheap {
	h := &mheap{}
	for i := 0; i < size; i++ {
		h.spans = append(h.spans, &span{size: i + 1})
	}
	return h
}

func (h *mheap) getSpan(size int) *span {
	h.lock.Lock()
	defer h.lock.Unlock()
	for _, span := range h.spans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}
	return nil
}

func (mc *mcentral) getSpanFromCentral(size int) *span {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	for _, span := range mc.sizeSpans {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}
	return nil
}

func (mc *mcache) getSpanFromCache(size int) *span {
	for _, span := range mc.localSpan {
		if span.size == size && !span.allocated {
			span.allocated = true
			return span
		}
	}
	return nil
}

func main() {
	heap := NewHeap(4)
	mcentral := &mcentral{sizeSpans: heap.spans}
	mcache := &mcache{}

	spanPosition := 5
	requestSpan := mcache.getSpanFromCache(spanPosition)
	if requestSpan == nil {
		requestSpan = mcentral.getSpanFromCentral(spanPosition)
	}
	if requestSpan == nil {
		requestSpan = heap.getSpan(spanPosition)
	}
	if requestSpan == nil {
		panic("Out of memory")
	}
	fmt.Println("Span allocated")
}
