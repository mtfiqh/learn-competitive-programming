package tiket_2_array_sum_is_n

import "testing"

func TestFind(t *testing.T) {
	expected := true
	arr := []int{
		4, 3, 6, 12, 3, -7, 4, 7,
	}
	found := Find(arr, 0)
	if found == expected {
		t.Log("passed")
	} else {
		t.Error("not passed")
	}
}
