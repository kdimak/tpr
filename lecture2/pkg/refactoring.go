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

func Statement(invoice Invoice, plays map[string]Play) (string, error) {
	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.customer)

	for _, perf := range invoice.performances {
		play, ok := plays[perf.playID]
		if !ok {
			return "", errors.New("play not found")
		}

		thisAmount := 0

		switch play._type {
		case "tragedy":
			thisAmount = 40000
			if perf.audience > 30 {
				thisAmount += 1000 * (perf.audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.audience > 20 {
				thisAmount += 10000 + 500*(perf.audience-20)
			}
			thisAmount += 300 * perf.audience
		default:
			return "", errors.New(fmt.Sprintf("unknown type: %s", play._type))
		}

		volumeCredits += int(math.Max(float64(perf.audience-30), 0))

		if "comedy" == play._type {
			volumeCredits += int(math.Floor(float64(perf.audience) / 5))
		}

		result += fmt.Sprintf(" %s: %.2f (%d seats)\n", play.name, float64(thisAmount)/100, perf.audience)
		totalAmount += thisAmount
	}

	result += fmt.Sprintf("Amount owed is %.2f\n", float64(totalAmount)/100)
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)

	return result, nil
}
