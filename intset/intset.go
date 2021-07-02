package synclist

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type IntSet struct {
	head   *intNode
	length int64
}

type intNode struct {
	value int64
	next  *intNode
	isDel uint32
	mu    sync.Mutex
}

func NewIntSet() *IntSet {
	return &IntSet{head: newIntNode(0)}
}

func newIntNode(value int) *intNode {
	return &intNode{value: int64(value)}
}

func (s *IntSet) Contains(value int) bool {
	//找到节点，并判断
	x := (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&s.head.next))))
	for x != nil && int(atomic.LoadInt64(&x.value)) < value {
		x = (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&x.next))))
	}
	if x == nil {
		return false
	}
	return int(atomic.LoadInt64(&x.value)) == value
}

func (s *IntSet) Insert(value int) bool {
begin:
	//1. 找到节点
	a := (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&s.head))))
	b := (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&a.next))))
	for b != nil && int(atomic.LoadInt64(&b.value)) < value {
		a = b
		b = (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&b.next))))
	}
	if b != nil && int(atomic.LoadInt64(&b.value)) == value {
		return false
	}
	//2. 锁定节点a，检查
	a.mu.Lock()
	if (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&a.next)))) != b || atomic.LoadUint32(&a.isDel) != 0 {
		a.mu.Unlock()
		goto begin
	}
	//3. 检查无误，添加新节点
	x := newIntNode(value)
	x.next = b
	atomic.StorePointer((*unsafe.Pointer)((unsafe.Pointer)(&a.next)), unsafe.Pointer(x))
	atomic.AddInt64(&s.length, 1)
	//4. 解锁
	a.mu.Unlock()
	return true
}

func (s *IntSet) Delete(value int) bool {
begin:
	//1. 找到节点
	a := (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&s.head))))
	b := (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&a.next))))
	for b != nil && int(atomic.LoadInt64(&b.value)) < value {
		a = b
		b = (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&b.next))))
	}
	if b == nil || int(atomic.LoadInt64(&b.value)) != value {
		return false
	}
	//2. 锁定b,检查
	b.mu.Lock()
	if atomic.LoadUint32(&b.isDel) != 0 {
		b.mu.Unlock()
		goto begin
	}
	//3. 锁定a,检查
	a.mu.Lock()
	if (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&a.next)))) != b || atomic.LoadUint32(&a.isDel) != 0 {
		a.mu.Unlock()
		b.mu.Unlock()
		goto begin
	}
	//4. 操作并解锁
	atomic.StoreUint32(&b.isDel, 1)
	atomic.StorePointer((*unsafe.Pointer)((unsafe.Pointer)(&a.next)), unsafe.Pointer(b.next))
	atomic.AddInt64(&s.length, -1)
	a.mu.Unlock()
	b.mu.Unlock()
	return true
}

func (s *IntSet) Len() int {
	return int(atomic.LoadInt64(&s.length))
}

func (s *IntSet) Range(f func(value int) bool) {
	x := (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&s.head.next))))
	for x != nil {
		if !f(int(atomic.LoadInt64(&x.value))) {
			break
		}
		x = (*intNode)(atomic.LoadPointer((*unsafe.Pointer)((unsafe.Pointer)(&x.next))))
	}
}
