package dummy

import "testing"

func TestDummy(t *testing.T) {
	if !Dummy() {
		t.Error("Dummy() should return true")
	}
}
