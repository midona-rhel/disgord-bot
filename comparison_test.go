package disgordbot

import "testing"

func Test_compareString(t *testing.T) {
	if b := compareString(" a a ", "A A"); !b {
		t.Error("Expected a true")
	}
	if b := compareString(" a a ", "AA"); b {
		t.Error("Expected a false")
	}
}
