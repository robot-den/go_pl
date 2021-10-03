package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	got := IntMin(2, -2)
	want := -2

	if got != want {
		t.Errorf("IntMin(2, -2) = %d, want %d", got, want)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{-2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("IntMin(%d, %d)", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			got := IntMin(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("IntMin(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
