package config

// CronFieldRanges defines the valid ranges for different cron fields.
var CronFieldRanges = map[string][2]int{
	"minute":       {0, 59},
	"hour":         {0, 23},
	"day_of_month": {1, 31},
	"month":        {1, 12},
	"day_of_week":  {0, 6},
}

// GetRange returns the range for a given field type.
func GetRange(fieldType string) (int, int, bool) {
	rangeVals, ok := CronFieldRanges[fieldType]
	return rangeVals[0], rangeVals[1], ok
}
