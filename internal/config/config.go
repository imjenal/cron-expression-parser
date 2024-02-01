package config

import (
	"cron_expr_parser/internal/utils"
	"strconv"
	"strings"
)

const (
	MinuteField     = "minute"
	HourField       = "hour"
	DayOfMonthField = "day of month"
	MonthField      = "month"
	DayOfWeekField  = "day of week"
)

// CronFieldRanges defines the valid ranges for different cron fields.
var CronFieldRanges = map[string][2]int{
	MinuteField:     {0, 59},
	HourField:       {0, 23},
	DayOfMonthField: {1, 31},
	MonthField:      {1, 12},
	DayOfWeekField:  {0, 6},
}

var DayOfWeekNames = map[string]int{
	"sun": 0,
	"mon": 1,
	"tue": 2,
	"wed": 3,
	"thu": 4,
	"fri": 5,
	"sat": 6,
}

var DayOfMonthNames = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

// GetRange returns the range for a given field type.
func GetRange(fieldType string) (int, int, bool) {
	rangeVals, ok := CronFieldRanges[fieldType]
	return rangeVals[0], rangeVals[1], ok
}

func ConvertNamesToNumbers(value string, nameMap map[string]int) string {
	parts := strings.Split(value, ",")
	for index, part := range parts {
		switch utils.DetermineExpressionType(part) {
		case utils.Range:
			rangeParts := strings.Split(part, "-")
			for j, rangePart := range rangeParts {
				rangeParts[j] = convertName(rangePart, nameMap)
			}
			parts[index] = strings.Join(rangeParts, "-")
		case utils.SingleValue:
			parts[index] = convertName(part, nameMap)
		}
	}
	return strings.Join(parts, ",")
}
func convertName(name string, nameMap map[string]int) string {
	if num, ok := nameMap[strings.ToLower(name)]; ok {
		return strconv.Itoa(num)
	}
	return name

}
