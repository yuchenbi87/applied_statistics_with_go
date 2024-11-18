package main

import "testing"

// Unit tests
func TestGenerateSample(t *testing.T) {
	n := 10
	mean := 100.0
	sd := 15.0
	sample := generateSample(n, mean, sd)
	if len(sample) != n {
		t.Errorf("Expected sample length %d, got %d", n, len(sample))
	}
}

func TestBootstrapSample(t *testing.T) {
	sample := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	n := 5
	bootstrap := bootstrapSample(sample, n)
	if len(bootstrap) != n {
		t.Errorf("Expected bootstrap length %d, got %d", n, len(bootstrap))
	}
}

func TestCalculateMean(t *testing.T) {
	data := []map[string]float64{
		{"bootMean": 1.0},
		{"bootMean": 2.0},
		{"bootMean": 3.0},
	}
	mean := calculateMean(data, "bootMean")
	if mean != 2.0 {
		t.Errorf("Expected mean 2.0, got %.2f", mean)
	}
}
