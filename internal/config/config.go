package config

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

// GetRange returns the range for a given field type.
func GetRange(fieldType string) (int, int, bool) {
	rangeVals, ok := CronFieldRanges[fieldType]
	return rangeVals[0], rangeVals[1], ok
}
