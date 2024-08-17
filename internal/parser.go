package internal

import (
	"cron-expression-parser/internal/common"
	"cron-expression-parser/internal/config"
	"cron-expression-parser/internal/fields"
	"fmt"
	"strings"
)

func NewCronField(fieldType, value string) (common.CronField, error) {
	// Map from field type string to constructor functions
	fieldConstructors := map[string]func(string) common.CronField{
		config.MinuteField:     fields.NewMinuteField,
		config.HourField:       fields.NewHourField,
		config.DayOfMonthField: fields.NewDayOfMonthField,
		config.MonthField:      fields.NewMonthField,
		config.DayOfWeekField:  fields.NewDayOfWeekField,
	}

	constructor, exists := fieldConstructors[fieldType]
	if !exists {
		return nil, fmt.Errorf("unknown field type: %s", fieldType)
	}

	return constructor(value), nil
}

// ParseCronExpression parses a cron expression and returns the expanded values for each field.
func ParseCronExpression(cronExpression string) ([]string, [][]int, []string, error) {
	cronFields := strings.Fields(cronExpression)
	expectedFields := len(config.CronFieldTypes) // This should be 5 (minute, hour, day of month, month, day of week)

	if len(cronFields) != expectedFields+1 {
		return nil, nil, nil, fmt.Errorf("invalid cron expression. It should have %d fields ", expectedFields+1)
	}

	orderedFieldNames := []string{
		config.MinuteField,
		config.HourField,
		config.DayOfMonthField,
		config.MonthField,
		config.DayOfWeekField,
	}

	fieldValues := make([][]int, len(orderedFieldNames))

	for i, fieldName := range orderedFieldNames {
		cronField, err := NewCronField(fieldName, cronFields[i])
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to create field '%s': %w", fieldName, err)
		}

		if err := cronField.Validate(); err != nil {
			return nil, nil, nil, fmt.Errorf("validation failed for field '%s': %w", fieldName, err)
		}

		expandedValues, err := cronField.Expand()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("expansion failed for field '%s': %w", fieldName, err)
		}

		fieldValues[i] = expandedValues
	}

	command := cronFields[len(orderedFieldNames):]

	return orderedFieldNames, fieldValues, command, nil
}
