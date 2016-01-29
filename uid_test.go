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

func benchmarkGen(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(i)
	}
}

func BenchmarkGen1(b *testing.B)  { benchmarkGen(1, b) }
func BenchmarkGen5(b *testing.B)  { benchmarkGen(5, b) }
func BenchmarkGen10(b *testing.B) { benchmarkGen(10, b) }
func BenchmarkGen20(b *testing.B) { benchmarkGen(20, b) }
func BenchmarkGen50(b *testing.B) { benchmarkGen(50, b) }
