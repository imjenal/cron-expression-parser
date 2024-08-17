package fields

import (
	"cron-expression-parser/internal/common"
	"cron-expression-parser/internal/config"
	"cron-expression-parser/internal/utils"
	"fmt"
	"strings"
)

type HourField struct {
	Value      string
	StartRange int
	EndRange   int
}

var _ common.CronField = &HourField{}

func NewHourField(value string) common.CronField {
	return &HourField{
		Value:      value,
		StartRange: config.HourFieldType.Min,
		EndRange:   config.HourFieldType.Max,
	}
}

func (f *HourField) Validate() error {
	// Handle the wildcard case
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

func (f *HourField) Expand() ([]int, error) {
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
