package helpers

import (
	"testing"
	"time"
)

func TestFormatParameter_String(t *testing.T) {
	result := FormatParameter("test-value", "form")
	if result != "test-value" {
		t.Errorf("Expected 'test-value', got '%s'", result)
	}
}

func TestFormatParameter_Int32(t *testing.T) {
	result := FormatParameter(int32(42), "form")
	if result != "42" {
		t.Errorf("Expected '42', got '%s'", result)
	}
}

func TestFormatParameter_Bool(t *testing.T) {
	if FormatParameter(true, "form") != "true" {
		t.Error("Expected 'true'")
	}
	if FormatParameter(false, "form") != "false" {
		t.Error("Expected 'false'")
	}
}

func TestFormatParameter_StringArray(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		style    string
		expected string
	}{
		{"form style", []string{"a", "b", "c"}, "form", "a,b,c"},
		{"space delimited", []string{"a", "b", "c"}, "spaceDelimited", "a b c"},
		{"pipe delimited", []string{"a", "b", "c"}, "pipeDelimited", "a|b|c"},
		{"empty array", []string{}, "form", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatParameter(tt.values, tt.style)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestFormatParameter_IntArray(t *testing.T) {
	result := FormatParameter([]int{1, 2, 3}, "form")
	if result != "1,2,3" {
		t.Errorf("Expected '1,2,3', got '%s'", result)
	}
}

func TestFormatParameter_Int32Array(t *testing.T) {
	result := FormatParameter([]int32{10, 20, 30}, "form")
	if result != "10,20,30" {
		t.Errorf("Expected '10,20,30', got '%s'", result)
	}
}

func TestOptionalParameter_Nil(t *testing.T) {
	var nilPtr *string
	result := OptionalParameter(nilPtr, "form")
	if result != "" {
		t.Errorf("Expected empty string for nil pointer, got '%s'", result)
	}
}

func TestOptionalParameter_Value(t *testing.T) {
	value := "test"
	result := OptionalParameter(&value, "form")
	if result != "test" {
		t.Errorf("Expected 'test', got '%s'", result)
	}
}

func TestFormatTime(t *testing.T) {
	testTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	result := FormatTime(testTime)
	expected := "2024-01-01T12:00:00Z"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFormatTimePtr_Nil(t *testing.T) {
	var nilTime *time.Time
	result := FormatTimePtr(nilTime)
	if result != "" {
		t.Errorf("Expected empty string for nil time, got '%s'", result)
	}
}

func TestFormatTimePtr_Value(t *testing.T) {
	testTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	result := FormatTimePtr(&testTime)
	expected := "2024-01-01T12:00:00Z"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
