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
			if f != nil {
				if (f.StartRange != tt.wantStart) || (f.EndRange != tt.wantEnd) {
					t.Errorf("NewDayOfWeekField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", f.StartRange, f.EndRange, tt.wantStart, tt.wantEnd)
				}
			} else {
				if tt.wantErr == false {
					t.Errorf("NewDayOfWeekField() error = %v, wantErr %v", f, tt.wantErr)
				}
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
