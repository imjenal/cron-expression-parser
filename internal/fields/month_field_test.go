package fields

import (
	"reflect"
	"testing"
)

func TestNewMonthField(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{"Valid - every month", "*", 1, 12, false},
		{"Valid - specific months", "1,6,12", 1, 12, false},
		{"Invalid value", "13", 1, 12, true},
		{"Invalid range", "0-12", 1, 12, true},
		{"Invalid step", "*/13", 1, 12, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewMonthField(tt.value)
			// Type assert to the concrete type
			field, ok := f.(*MonthField)
			if !ok {
				t.Fatalf("NewMonthField() did not return a MonthField")
			}

			if (field.StartRange != tt.wantStart) || (field.EndRange != tt.wantEnd) {
				t.Errorf("NewMonthField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", field.StartRange, field.EndRange, tt.wantStart, tt.wantEnd)
			}
		})
	}
}

func TestMonthField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid - every month", "*", false},
		{"Valid - specific months", "1,6,12", false},
		{"Invalid value", "13", true},
		{"Invalid range", "0-12", true},
		{"Invalid step", "*/13", true},
		{"Invalid value - letters", "abc", true},
		{"Invalid value - special characters", "$%&", true},
		{"Invalid value - mixed characters", "1a", true},
		{"Out of range - negative", "-1", true},
		{"Out of range - zero", "0", true},
		{"Out of range - above max", "13", true},
		{"Invalid - special characters", "!@#$", true},
		{"Invalid - mixed valid and invalid", "1a*", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewMonthField(tt.value)
			if err := field.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("MonthField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMonthField_Expand(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		wantExpanded []int
		wantErr      bool
	}{
		{"Valid - every month", "*", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, false},
		{"Valid - specific months", "1,6,12", []int{1, 6, 12}, false},
		{"Valid - range", "2-5", []int{2, 3, 4, 5}, false},
		{"Valid - step", "*/2", []int{1, 3, 5, 7, 9, 11}, false},
		{"Valid - single month", "5", []int{5}, false},
		{"Valid - single month at start", "1", []int{1}, false},
		{"Valid - single month at end", "12", []int{12}, false},
		{"Valid - multiple ranges", "1-3,6-8", []int{1, 2, 3, 6, 7, 8}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewMonthField(tt.value)
			expanded, err := field.Expand()
			if (err != nil) != tt.wantErr {
				t.Errorf("MonthField.Expand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(expanded, tt.wantExpanded) {
				t.Errorf("MonthField.Expand() = %v, want %v", expanded, tt.wantExpanded)
			}
		})
	}
}
