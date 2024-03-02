package pkg

import (
	"errors"
	"fmt"
	"math"
)

type Performance struct {
	playID   string
	audience int
}

type Play struct {
	name  string
	_type string
}

type Plays map[string]Play

func (p Plays) PlayFor(perf Performance) Play {
	return p[perf.playID]
}

func (p Plays) HasNoPlay(perf Performance) bool {
	_, ok := p[perf.playID]

	return !ok
}

type Invoice struct {
	customer     string
	performances []Performance
}

// SOLID principles:
// - Single Responsibility Principle: The function Statement has only one reason to change, which is to calculate the statement for a given invoice.
// - Open/Closed Principle: The function Statement is open for extension and closed for modification. The function can be extended by adding new play types without modifying the existing code.
// - Liskov Substitution Principle: The function Statement does not have any subclasses.
// - Interface Segregation Principle: The function Statement does not have any interfaces.
// - Dependency Inversion Principle: The function Statement does not depend on any concrete implementations.

func Statement(invoice Invoice, plays Plays) (string, error) {
	totalAmount := 0
	totalVolumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.customer)

	for _, perf := range invoice.performances {
		if plays.HasNoPlay(perf) {
			return "", errors.New("play not found")
		}

		perfAmount, err := amountFor(plays.PlayFor(perf), perf)
		if err != nil {
			return "", err
		}
		totalAmount += perfAmount

		totalVolumeCredits += volumeCreditsFor(perf, plays)

		result += fmt.Sprintf(" %s: %.2f (%d seats)\n", plays.PlayFor(perf).name, float64(perfAmount)/100, perf.audience)
	}

	result += fmt.Sprintf("Amount owed is %.2f\n", float64(totalAmount)/100)
	result += fmt.Sprintf("You earned %d credits\n", totalVolumeCredits)

	return result, nil
}

func volumeCreditsFor(perf Performance, plays Plays) int {
	result := int(math.Max(float64(perf.audience-30), 0))

	if "comedy" == plays.PlayFor(perf)._type {
		result += int(math.Floor(float64(perf.audience) / 5))
	}

	return result
}

func amountFor(play Play, perf Performance) (int, error) {
	result := 0

	switch play._type {
	case "tragedy":
		result = 40000
		if perf.audience > 30 {
			result += 1000 * (perf.audience - 30)
		}
	case "comedy":
		result = 30000
		if perf.audience > 20 {
			result += 10000 + 500*(perf.audience-20)
		}
		result += 300 * perf.audience
	default:
		return 0, errors.New(fmt.Sprintf("unknown type: %s", play._type))
	}

	return result, nil
}
