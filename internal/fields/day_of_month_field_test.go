package fields

import (
	"reflect"
	"testing"
)

func TestNewDayOfMonthField(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{"Valid - every day", "*", 1, 31, false},
		{"Valid - specific days", "1,15,31", 1, 31, false},
		{"Invalid value- out of bounds", "33", 1, 31, true},
		{"Invalid range", "1-32", 1, 31, true},
		{"Invalid step", "*/33", 1, 31, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewDayOfMonthField(tt.value)
			// Type assert to concrete type
			field, ok := f.(*DayOfMonthField)
			if !ok {
				t.Fatalf("NewDayOfMonthField() did not return a DayOfMonthField")
			}

			if (field.StartRange != tt.wantStart) || (field.EndRange != tt.wantEnd) {
				t.Errorf("NewDayOfMonthField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", field.StartRange, field.EndRange, tt.wantStart, tt.wantEnd)
			}
		})
	}
}

func TestDayOfMonthField_Expand(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		wantExpanded []int
		wantErr      bool
	}{
		{"Valid - every day", "*", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}, false},
		{"Valid - specific days", "1,15,31", []int{1, 15, 31}, false},
		{"Valid - range", "5-10", []int{5, 6, 7, 8, 9, 10}, false},
		{"Valid - every 2 days", "*/2", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}, false},
		{"Valid - every 3 days", "*/3", []int{1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31}, false},
		{"Valid - single day", "15", []int{15}, false},
		{"Valid - multiple ranges", "1-5,10-15", []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14, 15}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewDayOfMonthField(tt.value)
			expanded, err := field.Expand()
			if (err != nil) != tt.wantErr {
				t.Errorf("DayOfMonthField.Expand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(expanded, tt.wantExpanded) {
				t.Errorf("DayOfMonthField.Expand() = %v, want %v", expanded, tt.wantExpanded)
			}
		})
	}
}

func TestDayOfMonthField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid - every day", "*", false},
		{"Valid - specific days", "1,15,31", false},
		{"Invalid value", "32", true},
		{"Invalid range", "0-31", true},
		{"Invalid step", "*/32", true},
		{"Invalid value - letters", "abc", true},
		{"Invalid value - special characters", "$%&", true},
		{"Invalid value - mixed characters", "12a", true},
		{"Out of range - negative", "-5", true},
		{"Out of range - zero", "0", true},
		{"Out of range - above max", "32", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewDayOfMonthField(tt.value)
			if err := field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("DayOfMonthField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
