package utils

import (
	"math/cmplx"
	"time"
)

// The ComparingResult of the comparison is determined by the difference between the two values.
// Input a and b should return result:
//
//		Less  		if a < b
//		Equal      	if a = b
//		MoreThan 	if a > b
//	    Different 	if a.type != b.type
type ComparingResult int8

const (
	Less ComparingResult = iota
	Equal
	MoreThan
	Different
)

func Compare(a, b interface{}) ComparingResult {
	if a == nil || b == nil {
		return Different
	}
	typ, bTyp := getValueType(a), getValueType(b)
	if typ != bTyp || typ == NotCompare || bTyp == NotCompare {
		return Different
	}

	switch typ {
	case NotCompare:
		return Different
	case Bool:
		aA, bB := a.(bool), b.(bool)
		switch {
		case !aA && bB:
			return Less
		case aA && !bB:
			return MoreThan
		default:
			return Equal
		}
	case Int, Int8, Int16, Int32, Int64:
		var aA, bB int64
		switch typ {
		case Int:
			aA, bB = int64(a.(int)), int64(b.(int))
		case Int8:
			aA, bB = int64(a.(int8)), int64(b.(int8))
		case Int16:
			aA, bB = int64(a.(int16)), int64(b.(int16))
		case Int32:
			aA, bB = int64(a.(int32)), int64(b.(int32))
		case Int64:
			aA, bB = a.(int64), b.(int64)
		}
		switch {
		case aA > bB:
			return MoreThan
		case aA < bB:
			return Less
		default:
			return Equal
		}
	case Uint, Uint8, Uint16, Uint32, Uint64:
		var aA, bB uint64
		switch typ {
		case Uint:
			aA, bB = uint64(a.(uint)), uint64(b.(uint))
		case Uint8:
			aA, bB = uint64(a.(uint8)), uint64(b.(uint8))
		case Uint16:
			aA, bB = uint64(a.(uint16)), uint64(b.(uint16))
		case Uint32:
			aA, bB = uint64(a.(uint32)), uint64(b.(uint32))
		case Uint64:
			aA, bB = a.(uint64), b.(uint64)
		}
		switch {
		case aA > bB:
			return MoreThan
		case aA < bB:
			return Less
		default:
			return Equal
		}
	case Float32, Float64:
		var aA, bB float64
		switch typ {
		case Float32:
			aA, bB = float64(a.(float32)), float64(b.(float32))
		case Float64:
			aA, bB = a.(float64), b.(float64)
		}
		switch {
		case aA > bB:
			return MoreThan
		case aA < bB:
			return Less
		default:
			return Equal
		}
	case Complex64, Complex128:
		var aA, bB complex128
		switch typ {
		case Complex64:
			aA, bB = complex128(a.(complex64)), complex128(b.(complex64))
		case Complex128:
			aA, bB = a.(complex128), b.(complex128)
		}
		absA, absB := cmplx.Abs(aA), cmplx.Abs(bB)
		switch {
		case absA > absB:
			return MoreThan
		case absA < absB:
			return Less
		default:
			return Equal
		}
	case String, Bytes:
		var aA, bB []byte
		switch typ {
		case String:
			aA, bB = []byte(a.(string)), []byte(b.(string))
		case Bytes:
			aA, bB = a.([]byte), b.([]byte)
		}
		minLen := len(bB)
		if aALen := len(aA); aALen < minLen {
			minLen = aALen
		}
		diff := 0
		for i := 0; i < minLen && diff == 0; i++ {
			diff = int(aA[i]) - int(bB[i])
		}

		if diff == 0 {
			diff = len(aA) - len(bB)
		}
		switch {
		case diff > 0:
			return MoreThan
		case diff < 0:
			return Less
		default:
			return Equal
		}
	case Time:
		aA, bB := a.(time.Time), b.(time.Time)
		switch {
		case aA.After(bB):
			return MoreThan
		case aA.Before(bB):
			return Less
		default:
			return Equal
		}
	default:
		return Different
	}
}

type compareType = int8

const (
	NotCompare compareType = iota - 1
	Bool
	String
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Complex64
	Complex128
	Bytes
	Time
)

func getValueType(val interface{}) compareType {
	switch val.(type) {
	case bool:
		return Bool
	case int8:
		return Int8
	case int16:
		return Int16
	case int32:
		return Int32
	case int64:
		return Int64
	case int:
		return Int
	case uint8:
		return Uint8
	case uint16:
		return Uint16
	case uint32:
		return Uint32
	case uint64:
		return Uint64
	case uint:
		return Uint
	case float32:
		return Float32
	case float64:
		return Float64
	case complex64:
		return Complex64
	case complex128:
		return Complex128
	case string:
		return String
	case []byte:
		return Bytes
	case time.Time:
		return Time
	}
	return NotCompare
}
