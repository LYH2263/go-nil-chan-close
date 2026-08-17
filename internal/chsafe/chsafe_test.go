package chsafe

import "testing"

func TestCloseSafe(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic: %v", r)
		}
	}()
	Close(nil)
}
