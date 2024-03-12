package pkg

import (
	"fmt"
	"math"
)

type PlayType interface {
	AmountFor(Performance) int
	VolumeCreditsFor(Performance) int
}

func CreatePlayType(play Play) PlayType {
	switch play._type {
	case "tragedy":
		return Tragedy{}
	case "comedy":
		return Comedy{}
	default:
		panic(fmt.Sprintf("unknown type: %s", play._type))
	}
}

type Tragedy struct {
}

func (t Tragedy) AmountFor(perf Performance) int {
	result := 40000

	if perf.audience > 30 {
		result += 1000 * (perf.audience - 30)
	}

	return result
}

func (t Tragedy) VolumeCreditsFor(perf Performance) int {
	return int(math.Max(float64(perf.audience-30), 0))
}

type Comedy struct {
}

func (c Comedy) AmountFor(perf Performance) int {
	result := 30000

	if perf.audience > 20 {
		result += 10000 + 500*(perf.audience-20)
	}

	result += 300 * perf.audience

	return result
}

func (c Comedy) VolumeCreditsFor(perf Performance) int {
	return int(math.Max(float64(perf.audience-30), 0)) + int(math.Floor(float64(perf.audience)/5))
}
