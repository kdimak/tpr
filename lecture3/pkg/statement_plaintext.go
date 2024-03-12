package pkg

import (
	"fmt"
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

type StringReport string

// PlainTextStatementReport returns a plain text report for the given invoice and plays.
func PlainTextStatementReport(invoice Invoice, plays Plays) (StringReport, error) {
	sd, err := newStatementData(invoice, plays)
	if err != nil {
		return "", fmt.Errorf("create statement: %w", err)
	}

	return createStringReport(invoice, plays, sd), nil
}

func createStringReport(invoice Invoice, plays Plays, sd *statementData) StringReport {
	result := fmt.Sprintf("PlainTextStatementReport for %s\n", invoice.customer)

	for perf, amount := range sd.amountByPerformance {
		result += fmt.Sprintf(" %s: %.2f (%d seats)\n",
			plays.PlayFor(perf).name,
			float64(amount)/100,
			perf.audience)
	}

	result += fmt.Sprintf("Amount owed is %.2f\n", float64(sd.totalAmount)/100)
	result += fmt.Sprintf("You earned %d credits\n", sd.totalVolumeCredits)

	return StringReport(result)
}
