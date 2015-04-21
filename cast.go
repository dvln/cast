// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast is used to cast generic interface{} types
// into exact Go types.
package cast

import "time"

// ToBool attempts to cast an interface into what is expected to be a boolean,
// if you want to check for errors use ToBoolE() instead
func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToTime attempts to cast an interface into what should be a time.Time
// if you want to check for errors use ToTimeE() instead
func ToTime(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

// ToDuration attempts to cast an interface into what should be a time.Duration
// if you want to check for errors use ToDurationE() instead
func ToDuration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToFloat64 attempts to cast an interface into what should be a float64,
// if you want to check for errors use ToFloat64E() instead
func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

// ToInt attempts to cast an interface into what should be an int,
// if you want to check for errors use ToIntE() instead
func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

// ToString attempts to cast an interface into what should be a string,
// if you want to check for errors use ToStringE() instead
func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

// ToStringMapString attempts to cast an interface{} into what should be a map
// of string keys and string values, if you want to detect an error use
// ToStringMapStringE() directly
func ToStringMapString(i interface{}) map[string]string {
	v, _ := ToStringMapStringE(i)
	return v
}

// ToStringMapBool attempts to cast an interface{} into what should be a map
// of string keys and boolean values, if you want to detect an error use
// ToStringMapBoolE() directly
func ToStringMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

// ToStringMap attempts to cast an interface{} into what should be a map
// of string keys and interface{} values, if you want to detect an error use
// ToStringMapE() directly
func ToStringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

// ToSlice to cast an interface{} into what should be an array of
// interface{}, if you want to detect an error use ToSliceE() directly
func ToSlice(i interface{}) []interface{} {
	v, _ := ToSliceE(i)
	return v
}

// ToStringSlice to cast an interface{} into what should be an array of
// strings, if you want to detect an error use ToStringSliceE() directly
func ToStringSlice(i interface{}) []string {
	v, _ := ToStringSliceE(i)
	return v
}

// ToIntSlice to cast an interface{} into what should be an array of
// integers, if you want to detect an error use ToIntSliceE() directly
func ToIntSlice(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}
