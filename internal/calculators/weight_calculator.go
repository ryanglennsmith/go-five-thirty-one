package calculators

import (
	"errors"
	"math"
)
type WeightCalculator interface {
	CalculateWeight(modifier LiftModifier) (float64, error)
	CalculatePlates(weight float64) []float64
}

type weightCalculator struct {
	Lift string
	OneRM float64
	BaseWeight float64
}
type LiftModifier struct {
	Set     string
	Modifier float64
}

const baseModifier = 0.9

var WeightModifiers = []LiftModifier{
	{"W1S1", 0.65},
	{"W1S2", 0.75},
	{"W1S3", 0.85},
	{"W2S1", 0.7},
	{"W2S2", 0.8},
	{"W2S3", 0.9},
	{"W3S1", 0.75},
	{"W3S2", 0.85},
	{"W3S3", 0.95},
	{"W4S1", 0.4},
	{"W4S2", 0.5},
	{"W4S3", 0.6},
}

//constructor
func NewWeightCalculator(lift string, oneRM float64) WeightCalculator {
	return &weightCalculator{
		Lift: lift,
		OneRM: oneRM,
	}
}

func (wc *weightCalculator) CalculateBase() {
	wc.BaseWeight = wc.OneRM * baseModifier
}

func (wc *weightCalculator) CalculateWeight(modifier LiftModifier) (float64, error) {
	wc.CalculateBase()
	mod := modifier.Modifier
	if mod == 0 || !isValidModifierSet(modifier.Set) {
		return 0, errors.New("invalid modifier")
	}
	
	switch wc.Lift {
	case "DL":
		return math.Ceil(wc.BaseWeight * mod / 5) * 5, nil
	case "SQ":
		return math.Ceil(wc.BaseWeight * mod / 2.5) * 2.5, nil
	case "BP", "OHP":
		return math.Ceil(wc.BaseWeight * mod), nil
	default:
		return 0, errors.New("invalid lift")
	}
}

func (wc *weightCalculator) CalculatePlates(weightToLift float64) []float64 {
	// Mocked platesAvailable array
	platesAvailable := []float64{20, 15, 10, 5, 2.5, 2, 1.25, 1, 0.5}
	bar := 20.0
	var platesToUse []float64
	weightRemaining := (weightToLift - bar) / 2

	for i := 0; i < len(platesAvailable); i++ {
		if weightRemaining >= platesAvailable[i] {
			platesToUse = append(platesToUse, platesAvailable[i])
			weightRemaining -= platesAvailable[i]
			if weightRemaining >= platesAvailable[i] {
				i--
			}
		}
	}
	return platesToUse
}

func isValidModifierSet(set string) bool {
	for _, modifier := range WeightModifiers {
		if modifier.Set == set {
			return true
		}
	}
	return false
}