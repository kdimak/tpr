package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatement(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		invoice := Invoice{
			customer: "BigCo",
			performances: []Performance{
				{playID: "hamlet", audience: 55},
				{playID: "as-like", audience: 35},
				{playID: "othello", audience: 40},
			},
		}

		plays := map[string]Play{
			"hamlet":  {"Hamlet", "tragedy"},
			"as-like": {"As You Like It", "comedy"},
			"othello": {"Othello", "tragedy"},
		}

		got, err := Statement(invoice, plays)
		if err != nil {
			t.Errorf("Statement() returned an error: %v", err)
		}

		want := `Statement for BigCo
 Hamlet: 650.00 (55 seats)
 As You Like It: 580.00 (35 seats)
 Othello: 500.00 (40 seats)
Amount owed is 1730.00
You earned 47 credits
`

		assert.Equal(t, want, got)

		if got != want {
			t.Errorf("Statement() = %q, want %q", got, want)
		}
	})

	t.Run("Error: Play not found", func(t *testing.T) {
		invoice := Invoice{
			customer: "BigCo",
			performances: []Performance{
				{playID: "hamlet", audience: 55},
			},
		}

		plays := map[string]Play{
			"othello": {"Othello", "tragedy"},
		}

		got, err := Statement(invoice, plays)
		if err == nil {
			t.Errorf("Statement() did not return an error: %v", err)
		}

		if got != "" {
			t.Errorf("Statement() = %q, want %q", got, "")
		}
	})

	t.Run("Error: Unknown play type", func(t *testing.T) {
		invoice := Invoice{
			customer: "BigCo",
			performances: []Performance{
				{playID: "hamlet", audience: 55},
			},
		}

		plays := map[string]Play{
			"hamlet": {"Hamlet", "science-fiction"},
		}

		got, err := Statement(invoice, plays)
		if err == nil {
			t.Errorf("Statement() did not return an error: %v", err)
		}

		if got != "" {
			t.Errorf("Statement() = %q, want %q", got, "")
		}
	})
}
