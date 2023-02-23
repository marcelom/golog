package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_needsQuote(t *testing.T) {
	tests := []struct {
		s         string
		assertion assert.BoolAssertionFunc
	}{
		{"DontNeedQuotes", assert.False},
		{"Also_Dont_Need_Quotes", assert.False},
		{"Also-Also-Dont-Need-Quotes", assert.False},
		{"Need Quotes", assert.True},
		{"AlsoNeed.Quotes", assert.True},
		{"AlsoAlsoNeed$Quotes", assert.True},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := needsQuote(tt.s)
			tt.assertion(t, got)
		})
	}
}
