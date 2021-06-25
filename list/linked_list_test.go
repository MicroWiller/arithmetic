package list

import "testing"

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func TestMyLinkedList(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	firstNode := obj.dummy.Next
	secendNode := firstNode.Next
	thirdNode := secendNode.Next
	if firstNode.Value != 1 || secendNode.Value != 2 || thirdNode.Value != 3 {
		t.Errorf("expect: 1 2 3, but get %v %v %v", firstNode.Value, secendNode.Value, thirdNode.Value)
	}
	get1Value := obj.Get(1)
	if get1Value != 2 {
		t.Errorf("expect: 2, but get %v", get1Value)
	}
	obj.DeleteAtIndex(1)
	get2Value := obj.Get(1)
	if get2Value != 3 {
		t.Errorf("expect: 3, but get %v", get2Value)
	}
}
