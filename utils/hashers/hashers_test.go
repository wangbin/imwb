package hashers

import (
	"testing"
)

func TestMakePassword(t *testing.T) {
	p := MakePassword(RawPass)
	if p == RawPass {
		t.FailNow()
	}
}

func TestCheckPassword(t *testing.T) {
	p := MakePassword(RawPass)
	if !CheckPassword(RawPass, p) {
		t.FailNow()
	}
}
