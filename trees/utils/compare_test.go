package utils

import (
	"math"
	"testing"
	"time"
)

func TestCompare(t *testing.T) {
	baseTime := time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
	tests := []struct {
		name string
		a    interface{}
		b    interface{}
		want ComparingResult
	}{
		{name: "bool1", a: true, b: true, want: Equal},
		{name: "bool2", a: false, b: false, want: Equal},
		{name: "bool3", a: true, b: false, want: MoreThan},
		{name: "bool4", a: false, b: true, want: Less},

		{name: "number1", a: 1, b: 1, want: Equal},
		{name: "number2", a: 1, b: 2, want: Less},
		{name: "number3", a: 2, b: 1, want: MoreThan},
		{name: "number4", a: 102, b: 201, want: Less},
		{name: "number5", a: 0, b: 0, want: Equal},
		{name: "number6", a: 1, b: 0, want: MoreThan},
		{name: "number7", a: 0, b: 1, want: Less},

		{name: "uint_number1", a: uint(1), b: uint(1), want: Equal},
		{name: "uint_number2", a: uint(1), b: uint(2), want: Less},
		{name: "uint_number3", a: uint(2), b: uint(1), want: MoreThan},
		{name: "uint_number4", a: uint(102), b: uint(201), want: Less},
		{name: "uint_number5", a: uint(0), b: uint(0), want: Equal},
		{name: "uint_number6", a: uint(1), b: uint(0), want: MoreThan},
		{name: "uint_number7", a: uint(0), b: uint(1), want: Less},

		{name: "int32_number1", a: int32(1), b: int32(1), want: Equal},
		{name: "int32_number2", a: int32(1), b: int32(2), want: Less},
		{name: "int32_number3", a: int32(2), b: int32(1), want: MoreThan},
		{name: "int32_number4", a: int32(102), b: int32(201), want: Less},
		{name: "int32_number5", a: int32(0), b: int32(0), want: Equal},
		{name: "int32_number6", a: int32(1), b: int32(0), want: MoreThan},
		{name: "int32_number7", a: int32(0), b: int32(1), want: Less},

		{name: "float_number1", a: float64(1), b: float64(1), want: Equal},
		{name: "float_number2", a: float64(1), b: float64(2), want: Less},
		{name: "float_number3", a: float64(2), b: float64(1), want: MoreThan},
		{name: "float_number4", a: float64(102), b: float64(201), want: Less},
		{name: "float_number5", a: float64(0), b: float64(0), want: Equal},
		{name: "float_number6", a: float64(1), b: float64(0), want: MoreThan},
		{name: "float_number7", a: float64(0), b: float64(1), want: Less},

		{name: "complex128_number1", a: complex128(1), b: complex128(1), want: Equal},
		{name: "complex128_number2", a: complex128(1), b: complex128(2), want: Less},
		{name: "complex128_number3", a: complex128(2), b: complex128(1), want: MoreThan},
		{name: "complex128_number4", a: complex128(102), b: complex128(201), want: Less},
		{name: "complex128_number5", a: complex128(0), b: complex128(0), want: Equal},
		{name: "complex128_number6", a: complex128(1), b: complex128(0), want: MoreThan},
		{name: "complex128_number7", a: complex128(0), b: complex128(1), want: Less},

		{name: "string1", a: "a", b: "a", want: Equal},
		{name: "string2", a: "a", b: "b", want: Less},
		{name: "string3", a: "b", b: "a", want: MoreThan},
		{name: "string4", a: "aa", b: "aba", want: Less},
		{name: "string5", a: "", b: "", want: Equal},
		{name: "string6", a: "a", b: "", want: MoreThan},
		{name: "string7", a: "", b: "a", want: Less},
		{name: "string8", a: "", b: "a123", want: Less},

		{name: "bytes1", a: []byte{1}, b: []byte{1}, want: Equal},
		{name: "bytes2", a: []byte{1}, b: []byte{2}, want: Less},
		{name: "bytes3", a: []byte{2}, b: []byte{1}, want: MoreThan},
		{name: "bytes4", a: []byte{1, 1}, b: []byte{1, 2, 1}, want: Less},
		{name: "bytes5", a: []byte{1, 1}, b: []byte{1, 1}, want: Equal},
		{name: "bytes6", a: []byte{1, 1, 1}, b: []byte{}, want: MoreThan},
		{name: "bytes7", a: []byte{}, b: []byte{1, 1, 1}, want: Less},
		{name: "bytes8", a: []byte{}, b: []byte{1}, want: Less},

		{name: "time1", a: baseTime, b: baseTime, want: Equal},
		{name: "time2", a: baseTime.Add(time.Second), b: baseTime, want: MoreThan},
		{name: "time3", a: baseTime, b: baseTime.Add(time.Second), want: Less},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.a, tt.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompare(b *testing.B) {

	tests := []struct {
		name string
		a    interface{}
		b    interface{}
	}{
		{name: "bool", a: true, b: true},
		{name: "string1", a: "1", b: "2"},
		{name: "string2_many", a: "123456789123456789123456789", b: "123456789123456789123456789123456789"},
		{name: "int8_number", a: int8(1), b: int8(1)},
		{name: "int_number", a: 1000001, b: 1000001},
		{name: "float64_number", a: 1000001.01, b: 1000001.01},
		{name: "time", a: time.Now(), b: time.Now()},
		{name: "big_int64_number", a: math.MaxInt64, b: math.MaxInt64},
		{name: "big_float64_number", a: math.MaxFloat64, b: math.MaxFloat64},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Compare(tt.a, tt.b)
			}
		})
	}
}

func BenchmarkGetValueType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getValueType(1)
	}
}
