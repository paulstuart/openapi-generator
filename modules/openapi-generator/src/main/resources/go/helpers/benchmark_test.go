package helpers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// This benchmarks the OLD reflection-based approach (from current generator)
func parameterValueToStringReflection(obj interface{}, key string) string {
	var v = reflect.ValueOf(obj)
	if v == reflect.ValueOf(nil) {
		return ""
	}

	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.String:
		return obj.(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprint(obj)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprint(obj)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprint(obj)
	case reflect.Bool:
		return strconv.FormatBool(obj.(bool))
	case reflect.Slice:
		// Simplified slice handling
		if strSlice, ok := obj.([]string); ok {
			return strings.Join(strSlice, ",")
		}
		return fmt.Sprint(obj)
	default:
		return fmt.Sprint(obj)
	}
}

// Benchmark: Current reflection-based string parameter
func BenchmarkReflection_String(b *testing.B) {
	value := "test-value"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parameterValueToStringReflection(value, "")
	}
}

// Benchmark: New generic string parameter
func BenchmarkGeneric_String(b *testing.B) {
	value := "test-value"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatParameter(value, "form")
	}
}

// Benchmark: Current reflection-based int32 parameter
func BenchmarkReflection_Int32(b *testing.B) {
	value := int32(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parameterValueToStringReflection(value, "")
	}
}

// Benchmark: New generic int32 parameter
func BenchmarkGeneric_Int32(b *testing.B) {
	value := int32(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatParameter(value, "form")
	}
}

// Benchmark: Current reflection-based string array (10 items)
func BenchmarkReflection_StringArray10(b *testing.B) {
	values := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parameterValueToStringReflection(values, "")
	}
}

// Benchmark: New generic string array (10 items) with strings.Builder
func BenchmarkGeneric_StringArray10(b *testing.B) {
	values := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatParameter(values, "form")
	}
}

// Benchmark: strings.Join vs strings.Builder (for reference)
func BenchmarkStringsJoin_10Items(b *testing.B) {
	values := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(values, ",")
	}
}

func BenchmarkStringsBuilder_10Items(b *testing.B) {
	values := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9", "val10"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.Grow(len(values) * 10)
		for i, v := range values {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(v)
		}
		_ = sb.String()
	}
}

// Benchmark: Large array (100 items)
func BenchmarkReflection_StringArray100(b *testing.B) {
	values := make([]string, 100)
	for i := range values {
		values[i] = fmt.Sprintf("value%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parameterValueToStringReflection(values, "")
	}
}

func BenchmarkGeneric_StringArray100(b *testing.B) {
	values := make([]string, 100)
	for i := range values {
		values[i] = fmt.Sprintf("value%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatParameter(values, "form")
	}
}

// Benchmark: Mixed parameter types (realistic API call scenario)
func BenchmarkReflection_MixedParameters(b *testing.B) {
	stringParam := "device-name"
	intParam := int32(10)
	arrayParam := []string{"tag1", "tag2", "tag3"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = parameterValueToStringReflection(stringParam, "name")
		_ = parameterValueToStringReflection(intParam, "limit")
		_ = parameterValueToStringReflection(arrayParam, "tags")
	}
}

func BenchmarkGeneric_MixedParameters(b *testing.B) {
	stringParam := "device-name"
	intParam := int32(10)
	arrayParam := []string{"tag1", "tag2", "tag3"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FormatParameter(stringParam, "form")
		_ = FormatParameter(intParam, "form")
		_ = FormatParameter(arrayParam, "form")
	}
}
