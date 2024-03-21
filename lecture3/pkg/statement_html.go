package pkg

import "fmt"

type HTMLReport string

// HTMLStatementReport returns an HTML report for the given invoice and plays.
func HTMLStatementReport(invoice Invoice, plays Plays) (HTMLReport, error) {
	sd, err := buildStatementData(invoice, plays)
	if err != nil {
		return "", fmt.Errorf("create statement: %w", err)
	}

	return createHTMLReport(invoice, plays, sd), nil
}

func createHTMLReport(invoice Invoice, plays Plays, sd *statementData) HTMLReport {
	panic("not implemented")
}
