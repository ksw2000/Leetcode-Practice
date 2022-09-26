// https://leetcode.com/problems/design-circular-queue/
// 7 ms	6.8 MB

package main

type MyCircularQueue struct {
	list  []int
	front int
	rear  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		list:  make([]int, k),
		front: 0,
		rear:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if !this.IsFull() {
		this.list[this.rear%len(this.list)] = value
		this.rear++
		return true
	}
	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if !this.IsEmpty() {
		this.front++
		return true
	}
	return false
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[this.front%len(this.list)]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[(this.rear-1)%len(this.list)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	return this.rear-this.front == len(this.list)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
