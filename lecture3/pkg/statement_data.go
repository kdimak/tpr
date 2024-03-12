package pkg

import (
	"errors"
	"fmt"
)

// SOLID principles:
// - Single Responsibility Principle: The function PlainTextStatementReport has only one reason to change, which is to calculate the statement for a given invoice.
// - Open/Closed Principle: The function PlainTextStatementReport is open for extension and closed for modification. The function can be extended by adding new play types without modifying the existing code.
// - Liskov Substitution Principle: The function PlainTextStatementReport does not have any subclasses.
// - Interface Segregation Principle: The function PlainTextStatementReport does not have any interfaces.
// - Dependency Inversion Principle: The function PlainTextStatementReport does not depend on any concrete implementations.

type statementData struct {
	amountByPerformance map[Performance]int
	totalAmount         int
	totalVolumeCredits  int
}

func (sd *statementData) AddPerformance(perf Performance, amount int) {
	sd.amountByPerformance[perf] = amount
	sd.totalAmount += amount
}

func newStatementData() *statementData {
	return &statementData{
		amountByPerformance: make(map[Performance]int),
	}
}

func buildStatementData(invoice Invoice, plays Plays) (*statementData, error) {
	if err := validate(invoice, plays); err != nil {
		return nil, fmt.Errorf("create statement data: %w", err)
	}

	sd := newStatementData()

	for _, perf := range invoice.performances {
		sd.AddPerformance(perf, amountFor(plays.PlayFor(perf), perf))
	}

	sd.totalVolumeCredits = totalVolumeCredits(invoice, plays)

	return sd, nil
}

func amountFor(play Play, perf Performance) int {
	return CreatePlayType(play).AmountFor(perf)
}

func totalVolumeCredits(invoice Invoice, plays Plays) int {
	result := 0

	for _, perf := range invoice.performances {
		result += volumeCreditsFor(perf, plays)
	}

	return result
}

func volumeCreditsFor(perf Performance, plays Plays) int {
	return CreatePlayType(plays.PlayFor(perf)).VolumeCreditsFor(perf)
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
