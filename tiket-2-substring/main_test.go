package tiket_2_substring

import "testing"

func TestMaxSubstring(t *testing.T) {
	t.Log(MaxSubstring("5257523"))
	t.Log(MaxSubstring("776899834453999877"))
	t.Log("===========================")
	t.Log(MaxSub("5257523"))
	t.Log(MaxSub("776899834453999877"))
	t.Log("===========================")
	t.Log(MaxSubstringWithoutConvert("5257523"))
	t.Log(MaxSubstringWithoutConvert("776899834453999877"))
}
