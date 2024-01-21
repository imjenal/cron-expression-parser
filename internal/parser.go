package internal

import (
	"cron_expr_parser/internal/fields"
	"fmt"
	"strings"
)

func NewCronField(fieldType, value string) (CronField, error) {
	switch fieldType {
	case "minute":
		return fields.NewMinuteField(value), nil
	case "hour":
		return fields.NewHourField(value), nil
	case "day of month":
		return fields.NewDayOfMonthField(value), nil
	case "month":
		return fields.NewMonthField(value), nil
	case "day of week":
		return fields.NewDayOfWeekField(value), nil
	default:
		return nil, fmt.Errorf("unknown field type: %s", fieldType)
	}
}

func ParseCronExpression(cronExpression string) ([]string, [][]int, string, error) {
	fields := strings.Fields(cronExpression)
	if len(fields) != 6 {
		return nil, nil, "", fmt.Errorf("Invalid cron expression. It should have exactly 6 fields.")
	}

	fieldNames := []string{"minute", "hour", "day of month", "month", "day of week"}
	var fieldValues [][]int

	for i, fieldValue := range fields[:5] {
		cronField, err := NewCronField(fieldNames[i], fieldValue)
		if err != nil {
			return nil, nil, "", err
		}

		if err := cronField.Validate(); err != nil {
			return nil, nil, "", err
		}

		expandedValues, err := cronField.Expand()
		if err != nil {
			return nil, nil, "", err
		}

		fieldValues = append(fieldValues, expandedValues)
	}

	command := fields[5]
	fieldValues = append(fieldValues, []int{}) // Add an empty slice for the command

	return fieldNames, fieldValues, command, nil
}
