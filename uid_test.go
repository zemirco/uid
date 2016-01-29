package uid

import (
	"testing"
)

func TestNew(t *testing.T) {
	id := New(10)
	if len(id) != 10 {
		t.Error("Result length error")
	}
}

func TestNewBytes(t *testing.T) {
	id := NewBytes(10, "a")
	if len(id) != 10 {
		t.Error("Result length error")
	}
	if id != "aaaaaaaaaa" {
		t.Errorf("expected aaaaaaaaaa but got %s", id)
	}
}

func benchmarkNew(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(i)
	}
}

func BenchmarkNew1(b *testing.B)  { benchmarkNew(1, b) }
func BenchmarkNew5(b *testing.B)  { benchmarkNew(5, b) }
func BenchmarkNew10(b *testing.B) { benchmarkNew(10, b) }
func BenchmarkNew20(b *testing.B) { benchmarkNew(20, b) }
func BenchmarkNew50(b *testing.B) { benchmarkNew(50, b) }
