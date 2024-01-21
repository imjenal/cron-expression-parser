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
			if f != nil {
				if (f.StartRange != tt.wantStart) || (f.EndRange != tt.wantEnd) {
					t.Errorf("NewHourField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", f.StartRange, f.EndRange, tt.wantStart, tt.wantEnd)
				}
			} else {
				if tt.wantErr == false {
					t.Errorf("NewHourField() error = %v, wantErr %v", f, tt.wantErr)
				}
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
		{"Invalid value", "24", true},
		{"Invalid range", "0-24", true},
		{"Invalid step", "*/25", true},
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
