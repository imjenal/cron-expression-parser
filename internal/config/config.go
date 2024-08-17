package config

const (
	MinuteField     = "minute"
	HourField       = "hour"
	DayOfMonthField = "day of month"
	MonthField      = "month"
	DayOfWeekField  = "day of week"
)

// CronFieldType represents a type of cron field with its min and max ranges.
type CronFieldType struct {
	Min int
	Max int
}

// Define the different cron field types and their corresponding ranges.
var (
	MinuteFieldType     = CronFieldType{Min: 0, Max: 59}
	HourFieldType       = CronFieldType{Min: 0, Max: 23}
	DayOfMonthFieldType = CronFieldType{Min: 1, Max: 31}
	MonthFieldType      = CronFieldType{Min: 1, Max: 12}
	DayOfWeekFieldType  = CronFieldType{Min: 0, Max: 6}
)

// CronFieldTypes is a map of field names to their corresponding CronFieldType.
var CronFieldTypes = map[string]CronFieldType{
	MinuteField:     MinuteFieldType,
	HourField:       HourFieldType,
	DayOfMonthField: DayOfMonthFieldType,
	MonthField:      MonthFieldType,
	DayOfWeekField:  DayOfWeekFieldType,
}
