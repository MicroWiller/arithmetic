package queue

import (
	"testing"
)

func TestNewCircularQueueByUsed(t *testing.T) {
	circularQueue := NewCircularQueueByUsed(3)
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)
	status := circularQueue.EnQueue(4)
	if status != false {
		t.Errorf("Expect is fasle, but get status is:%v", status)
	}
	if circularQueue.Rear() != 3 {
		t.Errorf("Expect is:%v, but get status is:%v", 3, circularQueue.Rear())
	}
	if circularQueue.IsFull() != true {
		t.Errorf("Expect is true, but get status is:%v", circularQueue.IsFull())
	}
	circularQueue.Dequeue()
	circularQueue.EnQueue(4)
	if circularQueue.Rear() != 4 {
		t.Errorf("Expect is:%v, but get status is:%v", 4, circularQueue.Rear())
	}
}

func TestNewCircularQueueByPlus(t *testing.T) {
	circularQueue := NewCircularQueueByPlus(3)
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)
	status := circularQueue.EnQueue(4)
	if status != false {
		t.Errorf("Expect is fasle, but get status is:%v", status)
	}
	if circularQueue.Rear() != 3 {
		t.Errorf("Expect is:%v, but get status is:%v", 3, circularQueue.Rear())
	}
	if circularQueue.IsFull() != true {
		t.Errorf("Expect is true, but get status is:%v", circularQueue.IsFull())
	}
	circularQueue.DeQueue()
	circularQueue.EnQueue(4)
	if circularQueue.Rear() != 4 {
		t.Errorf("Expect is:%v, but get status is:%v", 4, circularQueue.Rear())
	}
}
