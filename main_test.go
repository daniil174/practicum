package main

import "testing"

func Test1(t *testing.T) {

}

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		in   float64
		want float64
	}{
		{
			"test zero",
			-0,
			0,
		},
		{
			"test 1",
			-1,
			1,
		},
		{
			"test same int",
			3,
			3,
		},
		{
			"test same float",
			3.0000001,
			3.0000001,
		},
		{
			"test float",
			-3.0000001,
			3.0000001,
		},
		{
			"test extra small",
			-0.0000000000000000001,
			0.0000000000000000001,
		},
		{
			"test same extra small",
			0.0000000000000000001,
			0.0000000000000000001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.in); got != tt.want {
				t.Errorf("Abs(%f) = %f, want %v", got, got, tt.want)
			}
		})
	}
}
