package que2

import "testing"

func TestPipe(t *testing.T) {
	if pipe(5, increment) == 6 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
	if pipe(5, increment, increment, increment) == 8 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
