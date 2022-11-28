package binarydiagnostic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	want, got := 3923414, Part1()
	assert.Equal(t, want, got)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Part1()
	}
}