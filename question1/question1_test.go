package que1

import (
	"testing"
)

func TestF(t *testing.T) {
	if f(9527) == "9,527" {
		t.Log("success")
	} else {
		t.Error("fail")
	}
	if f(3345678) == "3,345,678" {
		t.Log("success")
	} else {
		t.Error("fail")
	}
	if f(-1234.45) == "-1,234.45" {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
