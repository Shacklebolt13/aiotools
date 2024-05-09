package tools_test

import (
	shortener "aiotools/server/internal/tools"
	"testing"
)

func TestShortenId(t *testing.T) {
	const (
		input = 123
		want  = "1Z"
	)
	result := shortener.Shorten(123)
	if result != want {
		t.Errorf("Shorten(%d) = %s; want %s", input, result, want)
	}
}

func TestExpandId_WhenProperInput_ThenProperOutput(t *testing.T) {
	const (
		input = "1Z"
		want  = 123
	)
	result, e := shortener.Expand("1Z")
	if e != nil {
		t.Errorf("Expand(%s) = %d; want %d", input, result, want)
	}
	if result != want {
		t.Errorf("Expand(%s) = %d; want %d", input, result, want)
	}
}

func TestExpandId_WhenInvalidInput_ThenError(t *testing.T) {
	const (
		input = "1Z!"
	)
	_, e := shortener.Expand(input)
	if e == nil {
		t.Errorf("Expand(%s) = %v; want %v", input, e, "error")
	}
}

func TestLongRandom(t *testing.T) {
	for i := uint(0); i < 1000000; i++ {
		short := shortener.Shorten(i)
		result, e := shortener.Expand(short)
		if e != nil {
			t.Errorf("Expand(%s) = %d; want %d", short, result, i)
		}
		if result != i {
			t.Errorf("Expand(%s) = %d; want %d", short, result, i)
		}
	}
}
