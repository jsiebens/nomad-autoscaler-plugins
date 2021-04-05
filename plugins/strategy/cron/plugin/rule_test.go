package plugin

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRule_New(t *testing.T) {
	testCases := []struct {
		name          string
		inputValue    string
		expectedError bool
		expectedCount int64
	}{
		{
			name:          "default value",
			inputValue:    "* * 9-17 * * mon-fri *",
			expectedCount: 1,
		},
		{
			name:          "valid count",
			inputValue:    "* * 9-17 * * mon-fri *;5",
			expectedCount: 5,
		},
		{
			name:          "invalid expression",
			inputValue:    "invalid",
			expectedError: true,
		},
		{
			name:          "invalid count",
			inputValue:    "* * 9-17 * * mon-fri *;invalid",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		rule, err := parsePeriodRule(tc.inputValue, ";")
		if tc.expectedError {
			assert.NotNil(t, err)
			assert.Nil(t, rule)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedCount, rule.count)
		}
	}
}

func TestRule_Sort(t *testing.T) {
	rule1, _ := parsePeriodRule("* 1 * * *;1", ";")
	rule2, _ := parsePeriodRule("* 1 * * *;6", ";")
	rule3, _ := parsePeriodRule("* 1 * * *;5", ";")

	rules := []*Rule{
		rule1,
		rule2,
		rule3,
	}

	sort.Sort(RuleSorter(rules))

	assert.Equal(t, rule2, rules[0])
	assert.Equal(t, rule3, rules[1])
	assert.Equal(t, rule1, rules[2])
}
