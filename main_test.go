//go:build !integration

package main

import (
	"bytes"
	"math/big"
	"testing"
)

func TestTotalMonths(t *testing.T) {
	tests := []struct {
		name string
		d    Duration
		want int
	}{
		{"1y3m", Duration{1, 3}, 15},
		{"0y6m", Duration{0, 6}, 6},
		{"2y0m", Duration{2, 0}, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.TotalMonths(); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestToYearsDecimal(t *testing.T) {
	tests := []struct {
		name string
		d    Duration
		want *big.Float
	}{
		{"1y6m", Duration{1, 6}, big.NewFloat(float64(1.5))},
		{"2y6m", Duration{2, 6}, big.NewFloat(float64(2.5))},
		{"3y6m", Duration{3, 6}, big.NewFloat(float64(3.5))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ToYearsDecimal(); got.Cmp(tt.want) != 0 {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestSumDurations(t *testing.T) {
	tests := []struct {
		name      string
		durations []Duration
		want      Duration
	}{
		{"1y+1y=2y", []Duration{{1, 0}, {1, 0}}, Duration{2, 0}},
		{"0y6m+0y6m=1y", []Duration{{0, 6}, {0, 6}}, Duration{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumDurations(tt.durations); tt.want.Years != got.Years || tt.want.Months != got.Months {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		name        string
		test_string string
		want        Duration
	}{
		{"1y2m", "1y2m", Duration{1, 2}},
		{"3y2m", "3y2m", Duration{3, 2}},
		{"3years2m", "3years2m", Duration{3, 2}},
		{"4years 5m", "4years 5m", Duration{4, 5}},
		{"5 years 3 months", "5 years 3 months", Duration{5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ParseDuration(tt.test_string); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	var buf bytes.Buffer
	err := run([]string{"testdata/test.txt"}, &buf)
	if err != nil {
		t.Fatal(err)
	}
	if got := buf.String(); got != "18y0m\n" {
		t.Errorf("got %q, want %q", got, "18y0m\n")

	}
}
