package fields

import (
	"fmt"
	"testing"
)

func TestNewHourField(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{"Valid - every hour", "*", 0, 23, false},
		{"Valid - every 2 hours", "*/2", 0, 23, false},
		{"Valid - specific hours", "0,6,12", 0, 23, false},
		{"Invalid value", "24", 0, 23, true},
		{"Invalid range", "0-24", 0, 23, true},
		{"Invalid step", "*/25", 0, 23, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewHourField(tt.value)
			// Type assert to the concrete type
			field, ok := f.(*HourField)
			if !ok {
				t.Fatalf("NewHourField() did not return a HourField")
			}

			if (field.StartRange != tt.wantStart) || (field.EndRange != tt.wantEnd) {
				t.Errorf("NewHourField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", field.StartRange, field.EndRange, tt.wantStart, tt.wantEnd)
			}
		})
	}
}

func TestHourField_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"Valid - every hour", "*", false},
		{"Valid - every 2 hours", "*/2", false},
		{"Valid - specific hours", "0,6,12", false},
		{"Invalid value- out of bounds", "24", true},
		{"Invalid range", "0-24", true},
		{"Invalid step", "*/25", true},
		{"Invalid value - letters", "abc", true},
		{"Invalid value - special characters", "$%&", true},
		{"Invalid value - mixed characters", "12a", true},
		{"Out of range - negative", "-1", true},
		{"Out of range - zero", "0", false},
		{"Out of range - above max", "24", true},
		{"Invalid - special characters", "!@#$", true},
		{"Invalid - mixed valid and invalid", "12a*", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewHourField(tt.value)
			if err := f.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("HourField.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHourField_Expand(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    []int
		wantErr bool
	}{
		{"Valid - every hour", "*", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, false},
		{"Valid - every 2 hours", "*/2", []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22}, false},
		{"Valid - specific hours", "0,6,12", []int{0, 6, 12}, false},
		{"Valid - single hour", "12", []int{12}, false},
		{"Valid - single hour at start", "0", []int{0}, false},
		{"Valid - single hour at end", "23", []int{23}, false},
		{"Valid - multiple ranges", "1-3,5-7", []int{1, 2, 3, 5, 6, 7}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewHourField(tt.value)
			got, err := f.Expand()
			if (err != nil) != tt.wantErr {
				t.Errorf("HourField.Expand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("HourField.Expand() = %v, want %v", got, tt.want)
			}
		})
	}
}
