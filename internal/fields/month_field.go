package fields

import (
	"cron_expr_parser/internal/config"
	"cron_expr_parser/internal/utils"
	"fmt"
	"strings"
)

type MonthField struct {
	Value      string
	StartRange int
	EndRange   int
}

func NewMonthField(value string) *MonthField {
	value = config.ConvertNamesToNumbers(value, config.DayOfMonthNames)
	startRange, endRange, _ := config.GetRange(config.MonthField)
	return &MonthField{
		Value:      value,
		StartRange: startRange,
		EndRange:   endRange,
	}
}

func (f *MonthField) Validate() error {
	if f.Value == "*" {
		return nil
	}

	parts := strings.Split(f.Value, ",")
	for _, part := range parts {
		switch utils.DetermineExpressionType(part) {
		case utils.Range:
			if err := utils.ValidateRange(part, f.StartRange, f.EndRange); err != nil {
				return err
			}
		case utils.Step:
			if err := utils.ValidateStep(part, f.StartRange, f.EndRange); err != nil {
				return err
			}
		case utils.SingleValue:
			if err := utils.ValidateSingleValue(part, f.StartRange, f.EndRange); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *MonthField) Expand() ([]int, error) {
	if f.Value == "*" {
		return utils.GenerateSequence(f.StartRange, f.EndRange, 1), nil
	}

	var result []int
	parts := strings.Split(f.Value, ",")
	for _, part := range parts {
		switch utils.DetermineExpressionType(part) {
		case utils.Range:
			expanded, err := utils.ExpandRange(part)
			if err != nil {
				return nil, fmt.Errorf("expanding range: %w", err)
			}
			result = append(result, expanded...)
		case utils.Step:
			expanded, err := utils.ExpandStep(part, f.StartRange, f.EndRange)
			if err != nil {
				return nil, fmt.Errorf("expanding step: %w", err)
			}
			result = append(result, expanded...)
		case utils.SingleValue:
			value, err := utils.ConvertSingleValue(part)
			if err != nil {
				return nil, fmt.Errorf("converting single value: %w", err)
			}
			result = append(result, value)
		}
	}
	return result, nil
}
