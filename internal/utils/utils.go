package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type ExpressionType int

const (
	SingleValue ExpressionType = iota
	Range
	Step
)

// DetermineExpressionType takes a cron expression part and determines its type.
func DetermineExpressionType(part string) ExpressionType {
	if strings.Contains(part, "-") {
		return Range
	} else if strings.Contains(part, "/") {
		return Step
	}
	return SingleValue
}

// GenerateSequence generates a sequence of numbers within a range with a given step.
func GenerateSequence(start, end, step int) []int {
	var sequence []int
	for i := start; i <= end; i += step {
		sequence = append(sequence, i)
	}
	return sequence
}

// ValidateRange checks if the range is valid.
func ValidateRange(rangeExpr string, min, max int) error {
	rangeParts := strings.Split(rangeExpr, "-")
	start, err := strconv.Atoi(rangeParts[0])
	if err != nil {
		return fmt.Errorf("invalid range start: %s", rangeParts[0])
	}
	end, err := strconv.Atoi(rangeParts[1])
	if err != nil {
		return fmt.Errorf("invalid range end: %s", rangeParts[1])
	}
	if start < min || start > max || end < min || end > max || start > end {
		return fmt.Errorf("range out of bounds: %s", rangeExpr)
	}
	return nil
}

// ExpandRange expands a range expression.
func ExpandRange(rangeExpr string) ([]int, error) {
	rangeParts := strings.Split(rangeExpr, "-")
	start, _ := strconv.Atoi(rangeParts[0]) // already validated
	end, _ := strconv.Atoi(rangeParts[1])   // already validated
	return GenerateSequence(start, end, 1), nil
}

// ValidateStep checks if the step value is valid.
func ValidateStep(stepExpr string, min, max int) error {
	stepParts := strings.Split(stepExpr, "/")
	step, err := strconv.Atoi(stepParts[1])
	if err != nil {
		return fmt.Errorf("invalid step value: %s", stepParts[1])
	}
	if step <= 0 || step > max-min {
		return fmt.Errorf("step value out of bounds: %s", stepParts[1])
	}
	return nil
}

// ExpandStep expands a step expression.
func ExpandStep(stepExpr string, min, max int) ([]int, error) {
	step, _ := strconv.Atoi(strings.Split(stepExpr, "/")[1]) // already validated
	return GenerateSequence(min, max, step), nil
}

// ValidateSingleValue checks if a single value is valid.
func ValidateSingleValue(valueExpr string, min, max int) error {
	if valueExpr == "*" {
		return nil // "*" is a valid wildcard
	}
	value, err := strconv.Atoi(valueExpr)
	if err != nil {
		return fmt.Errorf("invalid value: %s", valueExpr)
	}
	if value < min || value > max {
		return fmt.Errorf("value out of bounds: %s", valueExpr)
	}
	return nil
}

// ConvertSingleValue converts a single value expression.
func ConvertSingleValue(valueExpr string) (int, error) {
	return strconv.Atoi(valueExpr) // assumed to be already validated
}
