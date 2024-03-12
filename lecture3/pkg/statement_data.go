package pkg

import (
	"errors"
	"fmt"
	"math"
)

// SOLID principles:
// - Single Responsibility Principle: The function Statement has only one reason to change, which is to calculate the statement for a given invoice.
// - Open/Closed Principle: The function Statement is open for extension and closed for modification. The function can be extended by adding new play types without modifying the existing code.
// - Liskov Substitution Principle: The function Statement does not have any subclasses.
// - Interface Segregation Principle: The function Statement does not have any interfaces.
// - Dependency Inversion Principle: The function Statement does not depend on any concrete implementations.

type statementData struct {
	amountByPerformance map[Performance]int
	totalAmount         int
	totalVolumeCredits  int
}

func newStatementData(invoice Invoice, plays Plays) (*statementData, error) {
	if err := validate(invoice, plays); err != nil {
		return nil, fmt.Errorf("create statement data: %w", err)
	}

	totalAmount := 0
	performanceAmount := make(map[Performance]int)

	for _, perf := range invoice.performances {
		performanceAmount[perf] = amountFor(plays.PlayFor(perf), perf)
		totalAmount += performanceAmount[perf]
	}

	return &statementData{
		amountByPerformance: performanceAmount,
		totalAmount:         totalAmount,
		totalVolumeCredits:  totalVolumeCredits(invoice, plays),
	}, nil
}

func amountFor(play Play, perf Performance) int {
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
		panic(fmt.Sprintf("unknown type: %s", play._type))
	}

	return result
}

func totalVolumeCredits(invoice Invoice, plays Plays) int {
	result := 0

	for _, perf := range invoice.performances {
		result += volumeCreditsFor(perf, plays)
	}

	return result
}

func volumeCreditsFor(perf Performance, plays Plays) int {
	result := int(math.Max(float64(perf.audience-30), 0))

	if "comedy" == plays.PlayFor(perf)._type {
		result += int(math.Floor(float64(perf.audience) / 5))
	}

	return result
}

func validate(invoice Invoice, plays Plays) error {
	for _, perf := range invoice.performances {
		if plays.HasNoPlay(perf) {
			return errors.New(fmt.Sprintf("play not found: %s", perf.playID))
		}
	}

	for _, play := range plays {
		switch play._type {
		case "tragedy", "comedy":
			break
		default:
			return errors.New(fmt.Sprintf("unknown type: %s", play._type))
		}
	}

	return nil
}
