package lib

import (
	"fmt"
	"regexp"
	"time"
)

const (
	Daily   = "daily"
	Weekly  = "weekly"
	Monthly = "monthly"
)

func PrintDateInfo(period string) string {
	now := time.Now()
	switch period {
	case Daily:
		return now.Format("2006-01-02")
	case Weekly:
		_, week := now.ISOWeek()
		return fmt.Sprintf("%d年第%d周", now.Year(), week)
	case Monthly:
		return fmt.Sprintf("%d年%d月", now.Year(), int(now.Month()))
	default:
		return ""
	}
}

func PrintDateRankCycle(period string) string {
	switch period {
	case Daily:
		return "today"
	case Weekly:
		return "this week"
	case Monthly:
		return "this month"
	default:
		return ""
	}
}

var regItem *regexp.Regexp

func init() {
	regItem = regexp.MustCompile(`\s+`)
}

// NoWhitespace 去除空格换行
func NoWhitespace(str string) string {
	noWhitespaceStr := regItem.ReplaceAllString(str, "")
	return noWhitespaceStr
}
