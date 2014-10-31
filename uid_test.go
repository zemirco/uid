package uid

import (
	"testing"
)

func TestGen(t *testing.T) {
	id, err := Gen(10)
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 10 {
		t.Error("Result length error")
	}
}
