package head

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	arr := []int{1, 2, 1, 1, 3, 3, 2, 3}
	k := 2
	expect1 := []int{1, 3}
	expect2 := []int{3, 1}

	result := TopKFrequent(arr, k)
	if !reflect.DeepEqual(result, expect1) {
		if !reflect.DeepEqual(result, expect2) {
			t.Errorf("expect %v or %v, but get %v", expect1, expect2, result)
		}
	}
}
