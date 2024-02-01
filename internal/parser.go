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

func ParseCronExpression(cronExpression string) ([]string, [][]int, []string, error) {
	emptyCommand := make([]string, 0)
	fields := strings.Fields(cronExpression)
	fmt.Println(fields)
	num_fields := len(fields)
	fmt.Println(num_fields)
	fieldsWithoutCommandLen := len(config.CronFieldRanges)
	fmt.Println(fieldsWithoutCommandLen)
	//expectedFields := len(config.CronFieldRanges) + 1
	/*	if fieldsWithoutCommandLen+1 != expectedFields {
			return nil, nil, emptyCommand, fmt.Errorf("Invalid cron expression. It should have exactly %d fields", expectedFields)
		}
	*/
	var fieldValues [][]int
	fieldNames := []string{
		config.MinuteField,
		config.HourField,
		config.DayOfMonthField,
		config.MonthField,
		config.DayOfWeekField,
	}

	for i, fieldValue := range fields[:fieldsWithoutCommandLen] {
		cronField, err := NewCronField(fieldNames[i], fieldValue)
		if err != nil {
			return nil, nil, emptyCommand, err
		}

		if err := cronField.Validate(); err != nil {
			return nil, nil, emptyCommand, err
		}

		expandedValues, err := cronField.Expand()
		if err != nil {
			return nil, nil, emptyCommand, err
		}

		fieldValues = append(fieldValues, expandedValues)
	}

	command := fields[fieldsWithoutCommandLen:num_fields]
	fieldValues = append(fieldValues, []int{}) // Add an empty slice for the command

	return fieldNames, fieldValues, command, nil
}
