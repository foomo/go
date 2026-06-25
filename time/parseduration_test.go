package time_test

import (
	"fmt"
	"testing"
	"time"

	gotime "github.com/foomo/go/time"
)

func ExampleParseDuration() {
	for _, s := range []string{"2w", "5d", "1w2d3h"} {
		d, _ := gotime.ParseDuration(s)
		fmt.Printf("%s = %s\n", s, d)
	}

	// Output:
	// 2w = 336h0m0s
	// 5d = 120h0m0s
	// 1w2d3h = 219h0m0s
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		in   string
		want time.Duration
	}{
		{"5d", 120 * time.Hour},
		{"2w", 336 * time.Hour},
		{"1.5d", 36 * time.Hour},
		{"-2w", -336 * time.Hour},
		{"1w2d3h4m5s", 168*time.Hour + 48*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second},
		{"2h45m", 2*time.Hour + 45*time.Minute},
		{"300ms", 300 * time.Millisecond},
		{"0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := gotime.ParseDuration(tt.in)
			if err != nil {
				t.Fatalf("ParseDuration(%q) unexpected error: %v", tt.in, err)
			}

			if got != tt.want {
				t.Errorf("ParseDuration(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseDurationErrors(t *testing.T) {
	for _, in := range []string{"", "10", "1x", ".s", "-"} {
		t.Run(in, func(t *testing.T) {
			if _, err := gotime.ParseDuration(in); err == nil {
				t.Errorf("ParseDuration(%q) expected error, got nil", in)
			}
		})
	}
}
