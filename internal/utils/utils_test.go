package utils

import (
	"reflect"
	"testing"
)

func TestDetermineExpressionType(t *testing.T) {
	tests := []struct {
		name     string
		part     string
		expected ExpressionType
	}{
		{"SingleValue", "5", SingleValue},
		{"Range", "1-5", Range},
		{"Step", "*/15", Step},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DetermineExpressionType(tt.part)
			if result != tt.expected {
				t.Errorf("DetermineExpressionType(%s) = %v, want %v", tt.part, result, tt.expected)
			}
		})
	}
}

func TestGenerateSequence(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		step     int
		expected []int
	}{
		{"Sequence by 1", 1, 5, 1, []int{1, 2, 3, 4, 5}},
		{"Sequence by 2", 0, 6, 2, []int{0, 2, 4, 6}},
		{"Sequence by 3", 1, 10, 3, []int{1, 4, 7, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateSequence(tt.start, tt.end, tt.step)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateSequence(%d, %d, %d) = %v, want %v", tt.start, tt.end, tt.step, result, tt.expected)
			}
		})
	}
}

func TestValidateRange(t *testing.T) {
	tests := []struct {
		name      string
		rangeExpr string
		min       int
		max       int
		wantErr   bool
	}{
		{"Valid range", "1-5", 1, 10, false},
		{"Invalid range - start greater than end", "5-1", 1, 10, true},
		{"Invalid range - start out of bounds", "0-5", 1, 10, true},
		{"Invalid range - end out of bounds", "1-11", 1, 10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRange(tt.rangeExpr, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateRange(%s) error = %v, wantErr %v", tt.rangeExpr, err, tt.wantErr)
			}
		})
	}
}

func TestExpandRange(t *testing.T) {
	tests := []struct {
		name      string
		rangeExpr string
		expected  []int
		wantErr   bool
	}{
		{"Valid range", "1-5", []int{1, 2, 3, 4, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExpandRange(tt.rangeExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandRange(%s) error = %v, wantErr %v", tt.rangeExpr, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExpandRange(%s) = %v, want %v", tt.rangeExpr, result, tt.expected)
			}
		})
	}
}

func TestValidateStep(t *testing.T) {
	tests := []struct {
		name     string
		stepExpr string
		min      int
		max      int
		wantErr  bool
	}{
		{"Valid step", "*/15", 0, 59, false},
		{"Invalid step - out of bounds", "*/60", 0, 59, true},
		{"Invalid step - negative", "*/-5", 0, 59, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStep(tt.stepExpr, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateStep(%s) error = %v, wantErr %v", tt.stepExpr, err, tt.wantErr)
			}
		})
	}
}

func TestExpandStep(t *testing.T) {
	tests := []struct {
		name     string
		stepExpr string
		min      int
		max      int
		expected []int
		wantErr  bool
	}{
		{"Valid step", "*/15", 0, 59, []int{0, 15, 30, 45}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExpandStep(tt.stepExpr, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandStep(%s) error = %v, wantErr %v", tt.stepExpr, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExpandStep(%s) = %v, want %v", tt.stepExpr, result, tt.expected)
			}
		})
	}
}

func TestValidateSingleValue(t *testing.T) {
	tests := []struct {
		name      string
		valueExpr string
		min       int
		max       int
		wantErr   bool
	}{
		{"Valid single value", "5", 0, 59, false},
		{"Valid wildcard", "*", 0, 59, false},
		{"Invalid single value", "60", 0, 59, true},
		{"Invalid single value - non-numeric", "abc", 0, 59, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSingleValue(tt.valueExpr, tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSingleValue(%s) error = %v, wantErr %v", tt.valueExpr, err, tt.wantErr)
			}
		})
	}
}

func TestConvertSingleValue(t *testing.T) {
	tests := []struct {
		name      string
		valueExpr string
		expected  int
		wantErr   bool
	}{
		{"Valid single value", "5", 5, false},
		{"Invalid single value - non-numeric", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertSingleValue(tt.valueExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertSingleValue(%s) error = %v, wantErr %v", tt.valueExpr, err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("ConvertSingleValue(%s) = %v, want %v", tt.valueExpr, result, tt.expected)
			}
		})
	}
}
