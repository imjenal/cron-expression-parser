package fields

import (
	"cron-expression-parser/internal/utils"
	"reflect"
	"testing"
)

func TestNewMinuteField(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{"Valid - every 15 minutes", "*/15", 0, 59, false},
		{"Valid - every specific minutes", "5,15,30", 0, 59, false},
		{"Invalid value", "60", 0, 59, true},
		{"Invalid range", "0-60", 0, 59, true},
		{"Invalid step", "*/61", 0, 59, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewMinuteField(tt.value)
			// Type assert to the concrete type
			field, ok := f.(*MinuteField)
			if !ok {
				t.Fatalf("NewMinuteField() did not return a MinuteField")
			}

			if (field.StartRange != tt.wantStart) || (field.EndRange != tt.wantEnd) {
				t.Errorf("NewMinuteField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", field.StartRange, field.EndRange, tt.wantStart, tt.wantEnd)
			}
		})
	}
}

func TestMinuteField_Expand(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected []int
		wantErr  bool
	}{
		{"Every minute", "*", utils.GenerateSequence(0, 59, 1), false},
		{"Every 15 minutes", "*/15", []int{0, 15, 30, 45}, false},
		{"Specific minutes", "5,15,30", []int{5, 15, 30}, false},
		{"Single minute", "15", []int{15}, false},
		{"Single minute at start", "0", []int{0}, false},
		{"Single minute at end", "59", []int{59}, false},
		{"Multiple ranges", "0-15,30-45", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewMinuteField(tt.value)
			got, err := f.Expand()
			if (err != nil) != tt.wantErr {
				t.Errorf("MinuteField.Expand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MinuteField.Expand() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMinuteField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid - every minute", "*", false},
		{"Valid - every 15 minutes", "*/15", false},
		{"Valid - specific minutes", "5,15,30", false},
		{"Invalid value", "60", true},
		{"Invalid range", "0-60", true},
		{"Invalid step", "*/61", true},
		{"Invalid value - letters", "abc", true},
		{"Invalid value - special characters", "$%&", true},
		{"Invalid value - mixed characters", "15a", true},
		{"Out of range - negative", "-1", true},
		{"Out of range - zero", "0", false},
		{"Out of range - above max", "60", true},
		{"Invalid - special characters", "!@#$", true},
		{"Invalid - mixed valid and invalid", "15a*", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewMinuteField(tt.value)
			if err := f.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("MinuteField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
