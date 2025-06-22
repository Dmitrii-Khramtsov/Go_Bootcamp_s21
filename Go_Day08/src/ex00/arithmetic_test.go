package arithmetic

import (
	"testing"
)

func TestGetElement(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		idx       int
		want      int
		expectErr bool
	}{
		{"valid index 0", []int{10, 20, 30}, 0, 10, false},
		{"valid index middle", []int{10, 20, 30}, 1, 20, false},
		{"valid index last", []int{10, 20, 30}, 2, 30, false},
		{"empty slice", []int{}, 0, 0, true},
		{"negative index", []int{1, 2, 3}, -1, 0, true},
		{"index out of range", []int{1, 2, 3}, 3, 0, true},
		{"index way out of range", []int{1, 2, 3}, 100, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getElement(tt.arr, tt.idx)
			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want && !tt.expectErr {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestGetElementUnsafePtr(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		idx       int
		want      int
		expectErr bool
	}{
		{"valid index 0", []int{10, 20, 30}, 0, 10, false},
		{"valid index middle", []int{10, 20, 30}, 1, 20, false},
		{"valid index last", []int{10, 20, 30}, 2, 30, false},
		{"empty slice", []int{}, 0, 0, true},
		{"negative index", []int{1, 2, 3}, -1, 0, true},
		{"index out of range", []int{1, 2, 3}, 3, 0, true},
		{"index way out of range", []int{1, 2, 3}, 100, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getElementUnsafePtr(tt.arr, tt.idx)
			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want && !tt.expectErr {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
