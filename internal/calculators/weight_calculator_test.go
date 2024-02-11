package calculators

import (
	"errors"
	"testing"
)

func TestCalculateWeight_deadlift(t *testing.T) {
	// Initialize a weightCalculator for testing
	wc := NewWeightCalculator("DL", 200.0)

	// Test cases for CalculateWeight
	tests := []struct {
		modifier LiftModifier
		expected float64
	}{
		{LiftModifier{"W1S1", 0.65}, 120},
		{LiftModifier{"W1S2", 0.75}, 135},
		{LiftModifier{"W2S3", 0.9}, 165},
	}

	for _, tt := range tests {
		result, err := wc.CalculateWeight(tt.modifier)
		if err != nil {
			t.Errorf("Error calculating weight: %s", err)
		}
		if result != tt.expected {
			t.Errorf("Expected %f for modifier %s, but got %f", tt.expected, tt.modifier.Set, result)
		}
	}
}

func TestCalculateWeight_squat(t *testing.T) {
	// Initialize a weightCalculator for testing
	wc := NewWeightCalculator("SQ", 150.0)

	// Test cases for CalculateWeight
	tests := []struct {
		modifier LiftModifier
		expected float64
	}{
		{LiftModifier{"W1S1", 0.65}, 90},
		{LiftModifier{"W1S2", 0.75}, 102.5},
		{LiftModifier{"W2S3", 0.9}, 122.5},
	}

	for _, tt := range tests {
		result, err := wc.CalculateWeight(tt.modifier)
		if err != nil {
			t.Errorf("Error calculating weight: %s", err)
		}
		if result != tt.expected {
			t.Errorf("Expected %f for modifier %s, but got %f", tt.expected, tt.modifier.Set, result)
		}
	}
}

func TestCalculateWeight_ohp(t *testing.T) {
	// Initialize a weightCalculator for testing
	wc := NewWeightCalculator("OHP", 75.0)

	// Test cases for CalculateWeight
	tests := []struct {
		modifier LiftModifier
		expected float64
	}{
		{LiftModifier{"W1S1", 0.65}, 44},
		{LiftModifier{"W1S2", 0.75}, 51},
		{LiftModifier{"W2S3", 0.9}, 61},
	}

	for _, tt := range tests {
		result, err := wc.CalculateWeight(tt.modifier)
		if err != nil {
			t.Errorf("Error calculating weight: %s", err)
		}
		if result != tt.expected {
			t.Errorf("Expected %f for modifier %s, but got %f", tt.expected, tt.modifier.Set, result)
		}
	}
}

func TestCalculateWeight_bad_lift(t *testing.T) {
	// Initialize a weightCalculator for testing
	wc := NewWeightCalculator("RDL", 0.0)
tests := []struct {
		modifier LiftModifier
		expected error
	}{
		{LiftModifier{"W1S1", 0.65}, errors.New("invalid lift")},
	}

	for _, tt := range tests {
		_, err := wc.CalculateWeight(tt.modifier)
		if err == nil {
			t.Errorf("Expected error for invalid lift, but got nil")
		}
		if err.Error() != tt.expected.Error() {
			t.Errorf("Expected error %s, but got %s", tt.expected, err)
		}
	}
}

func TestCalculateWeight_bad_modifier(t *testing.T) {
	// Initialize a weightCalculator for testing
	wc := NewWeightCalculator("DL", 200.0)

	// Test cases for CalculateWeight
	tests := []struct {
		modifier LiftModifier
		expected error
	}{
		{LiftModifier{"", 0.0}, errors.New("invalid modifier")},
		{LiftModifier{"W5S1", 50.0}, errors.New("invalid modifier")},
		{LiftModifier{"anything", 100.0}, errors.New("invalid modifier")},
		{LiftModifier{"W1S9", 100.0}, errors.New("invalid modifier")},
	}

	for _, tt := range tests {
		_, err := wc.CalculateWeight(tt.modifier)
		if err == nil {
			t.Errorf("Expected error for invalid modifier, but got nil")
		}
		if err.Error() != tt.expected.Error() {
			t.Errorf("Expected error %s, but got %s", tt.expected, err)
		}
	}
}

func TestCalculatePlates(t *testing.T) {
	// Initialize a weightCalculator for testing
	wc := NewWeightCalculator("DL", 200.0)

	// Test cases for CalculatePlates
	tests := []struct {
		weight float64
		expected []float64
	}{
		{0, []float64{}}, // no weight, no plates
		{44, []float64{10, 2}},
		{62.5, []float64{20, 1.25}},
		{90, []float64{20, 15}},
		{120, []float64{20, 20, 10}},
		{135, []float64{20, 20, 15, 2.5}},
		{165, []float64{20, 20, 20, 10, 2.5 }},
	}

	for _, tt := range tests {
		result := wc.CalculatePlates(tt.weight)
		if len(result) != len(tt.expected) {
			t.Errorf("Expected %d plates, but got %d", len(tt.expected), len(result))
		}
		for i, plate := range result {
			if plate != tt.expected[i] {
				t.Errorf("Expected %f for plate %d, but got %f", tt.expected[i], i, plate)
			}
		}
	}
}
