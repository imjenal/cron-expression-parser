package fields

import (
	"cron_expr_parser/internal/utils"
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
			if f != nil {
				if (f.StartRange != tt.wantStart) || (f.EndRange != tt.wantEnd) {
					t.Errorf("NewMinuteField() StartRange = %v, EndRange = %v, want StartRange = %v, want EndRange = %v", f.StartRange, f.EndRange, tt.wantStart, tt.wantEnd)
				}
			} else {
				if tt.wantErr == false {
					t.Errorf("NewMinuteField() error = %v, wantErr %v", f, tt.wantErr)
				}
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
