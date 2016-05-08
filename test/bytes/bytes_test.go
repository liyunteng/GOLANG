package bytes_test

import (
	"testing"
	"bytes"
)



func TestCompareIdenticalSlice(t *testing.T) {
	var b = []byte("hello, lyt")
	if bytes.Compare(b, b) != 0 {
		t.Error("b != b")
	}
	if bytes.Compare(b, b[:1]) != 1 {
		t.Error("b > b[:1] failed")
	}
}

// go test -v -bench CompareBytesEqual bytes_test.go
func BenchmarkCompareBytesEqual(b *testing.B) {
	b1 := []byte("Hello, lyt")
	b2 := []byte("Hello, lyt")

	for i := 0; i < b.N; i++ {
		if bytes.Compare(b1, b2) != 0 {
			b.Fatal("b1 != b2")
		}
	}
}
