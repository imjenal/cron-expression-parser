package internal

import (
	"cron_expr_parser/internal/config"
	"cron_expr_parser/internal/fields"
	"fmt"
	"strings"
)

func NewCronField(fieldType, value string) (CronField, error) {
	switch fieldType {
	case config.MinuteField:
		return fields.NewMinuteField(value), nil
	case config.HourField:
		return fields.NewHourField(value), nil
	case config.DayOfMonthField:
		return fields.NewDayOfMonthField(value), nil
	case config.MonthField:
		return fields.NewMonthField(value), nil
	case config.DayOfWeekField:
		return fields.NewDayOfWeekField(value), nil
	default:
		return nil, fmt.Errorf("unknown field type: %s", fieldType)
	}
}

func ParseCronExpression(cronExpression string) ([]string, [][]int, string, error) {
	fields := strings.Fields(cronExpression)
	num_fields := len(fields)
	expectedFields := len(config.CronFieldRanges) + 1
	if num_fields != expectedFields {
		return nil, nil, "", fmt.Errorf("Invalid cron expression. It should have exactly %d fields", expectedFields)
	}

	var fieldValues [][]int
	fieldNames := []string{
		config.MinuteField,
		config.HourField,
		config.DayOfMonthField,
		config.MonthField,
		config.DayOfWeekField,
	}

	fmt.Println(fieldNames)

	for i, fieldValue := range fields[:num_fields-1] {
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

	command := fields[num_fields-1]
	fieldValues = append(fieldValues, []int{}) // Add an empty slice for the command

	return fieldNames, fieldValues, command, nil
}
