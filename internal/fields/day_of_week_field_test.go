package fields

import (
	"reflect"
	"testing"
)

func TestNewDayOfWeekField(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{"Valid - every day of the week", "*", 0, 6, false},
		{"Valid - specific days of the week", "0,3,6", 0, 6, false},
		{"Invalid value", "7", 0, 6, true},
		{"Invalid range", "0-7", 0, 6, true},
		{"Invalid step", "*/7", 0, 6, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewDayOfWeekField(tt.value)
			// Type assert to the concrete type
			field, ok := f.(*DayOfWeekField)
			if !ok {
				t.Fatalf("NewDayOfWeekField() did not return a DayOfWeekField")
			}

			if (field.StartRange != tt.wantStart) || (field.EndRange != tt.wantEnd) {
				t.Errorf("NewDayOfWeekField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", field.StartRange, field.EndRange, tt.wantStart, tt.wantEnd)
			}
		})
	}
}

func TestDayOfWeekField_Expand(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		wantExpanded []int
		wantErr      bool
	}{
		{"Valid - every day of the week", "*", []int{0, 1, 2, 3, 4, 5, 6}, false},
		{"Valid - specific days of the week", "0,3,6", []int{0, 3, 6}, false},
		{"Valid - range", "0-4", []int{0, 1, 2, 3, 4}, false},
		{"Valid - every 2 days", "*/2", []int{0, 2, 4, 6}, false},
		{"Valid - every 3 days", "*/3", []int{0, 3, 6}, false},
		{"Valid - single day of the week", "3", []int{3}, false},
		{"Valid - multiple ranges", "1-2,4-5", []int{1, 2, 4, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewDayOfWeekField(tt.value)
			expanded, err := field.Expand()
			if (err != nil) != tt.wantErr {
				t.Errorf("DayOfWeekField.Expand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(expanded, tt.wantExpanded) {
				t.Errorf("DayOfWeekField.Expand() = %v, want %v", expanded, tt.wantExpanded)
			}
		})
	}
}

func TestDayOfWeekField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid - every day of the week", "*", false},
		{"Valid - specific days of the week", "0,3,6", false},
		{"Invalid value", "7", true},
		{"Invalid range", "0-7", true},
		{"Invalid step", "*/8", true},
		{"Invalid value - letters", "abc", true},
		{"Invalid value - special characters", "$%&", true},
		{"Invalid value - mixed characters", "1a2", true},
		{"Out of range - negative", "-1", true},
		{"Out of range - zero", "0", false},
		{"Out of range - above max", "8", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewDayOfWeekField(tt.value)
			if err := field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("DayOfWeekField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
