package plugin

import (
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/cronexpr"
)

func parsePeriodRule(value, separator string) (*Rule, error) {
	var count int64 = 1

	entries := strings.Split(value, separator)

	expr, err := cronexpr.Parse(strings.TrimSpace(entries[0]))
	if err != nil {
		return nil, err
	}

	if len(entries) > 1 {
		v, err := strconv.ParseInt(strings.TrimSpace(entries[1]), 10, 64)
		if err != nil {
			return nil, err
		}
		count = v
	}

	return &Rule{
		expression: expr,
		count:      count,
	}, nil
}

type Rule struct {
	expression *cronexpr.Expression
	count      int64
}

func (t *Rule) InPeriod(now time.Time) bool {
	nextIn := t.expression.Next(now)
	timeSince := now.Sub(nextIn)
	if -time.Second <= timeSince && timeSince <= time.Second {
		return true
	}

	return false
}

type RuleSorter []*Rule

func (r RuleSorter) Len() int           { return len(r) }
func (r RuleSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RuleSorter) Less(i, j int) bool { return r[j].count < r[i].count }
