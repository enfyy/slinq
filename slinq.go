// Package slinq is a collection of LINQ functions that is admittedly not nearly as versatile as its original.
// It is mostly only for slices (-> s(lices)linq) and unfortunately not chainable because generics usage on methods is quite restricted in go.
package slinq

import (
	"errors"
	"math"
)

// Aggregate applies an accumulator function to every element of the provided slice.
// The provided value is used as the initial value for the accumulator and the provided function is used to select the result value.
func Aggregate[T any](slice []T, initial T, accumulator func(T, T) T) T {
	result := initial
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}

// All returns true when all the elements in the provided slice satisfy the provided condition.
func All[T any](slice []T, condition func(T) bool) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if !condition(v) {
			return false
		}
	}
	return true
}

// Any returns true when at least one element in the provided slice satisfies the provided condition.
func Any[T any](slice []T, condition func(T) bool) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if condition(v) {
			return true
		}
	}
	return false
}

// Chunk returns a slice of slices of the provided size that contain the elements of the provided slice.
func Chunk[T any](slice []T, size int) ([][]T, error) {
	cap := math.Ceil(float64(len(slice)) / float64(size))
	result := make([][]T, int(cap))
	if size == 0 {
		return result, errors.New("size cannot be zero")
	}

	for i, v := range slice {
		chunkIndex := i / size
		result[chunkIndex] = append(result[chunkIndex], v)
	}

	return result, nil
}

// Count returns the count of elements in the provided slice that satisfy the provided condition.
func Count[T any](slice []T, condition func(T) bool) int {
	i := 0
	for _, v := range slice {
		if condition(v) {
			i++
		}
	}
	return i
}

// Distinct returns removes duplicate values from the provided slice. The order of the elements is not maintained/given.
func Distinct[T comparable](slice []T) []T {
	var result []T
	dict := make(map[T]int, len(slice))
	for i, v := range slice {
		dict[v] = i
	}
	for k, _ := range dict {
		result = append(result, k)
	}
	return result
}

// Except returns the elements of the first provided slice that don't appear in the second provided slice.
func Except[T comparable](first, second []T) []T {
	var result []T
	dict := make(map[T]int, len(second))
	for i, v := range second {
		dict[v] = i
	}
	for _, v := range first {
		if _, exists := dict[v]; !exists {
			result = append(result, v)
		}
	}
	return result
}

// Intersect returns the elements that appear in both of the provided slices.
func Intersect[T comparable](first, second []T) []T {
	var result []T
	dict := make(map[T]int, len(first))
	for i, v := range first {
		dict[v] = i
	}
	for _, v := range second {
		if _, exists := dict[v]; exists {
			result = append(result, v)
		}
	}
	return result
}

// First returns a pointer to the first element of the provided slice that satisfies the provided condition
func First[T any](slice []T, condition func(T) bool) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("slice is empty")
	}

	for _, v := range slice {
		if condition(v) {
			return v, nil
		}
	}
	return slice[0], errors.New("no element in the slice satisfies the condition")
}

// Repeat generates a slice that contains one repeated value the provided number of times.
func Repeat[T any](value T, count int) []T {
	var result []T
	for i := 0; i < count; i++ {
		result = append(result, value)
	}
	return result
}

// Reverse returns a slice with the elements of the provided slice in reversed order.
func Reverse[T any](slice []T) []T {
	var result []T
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}

// Select returns a slice of elements of the provided slice that have been modified by the provided selector.
func Select[TSource any, TResult any](slice []TSource, selector func(TSource) TResult) []TResult {
	var result []TResult
	for _, v := range slice {
		result = append(result, selector(v))
	}

	return result
}

// SelectMany returns a slice of elements of the provided slice that have been modified by the provided selector and flattened into a single slice.
func SelectMany[TSource any, TResult any](slice []TSource, selector func(TSource, int) []TResult) []TResult {
	var result []TResult
	for i, outer := range slice {
		for _, inner := range selector(outer, i) {
			result = append(result, inner)
		}
	}
	return result
}

// Single returns a single, specific element of the provided slice, returns an error if there are not exactly one element that satisfy the provided condition.
func Single[T any](slice []T, condition func(T) bool) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, errors.New("slice is empty")
	}
	var result T
	found := false

	for _, v := range slice {
		if condition(v) {
			if found {
				var zero T
				return zero, errors.New("more than one element satisfy the condition")
			} else {
				found = true
				result = v
			}
		}
	}

	return result, nil
}

// ToMap returns a map that was created by applying the provided key- and value-selector to the elements of the provided slice.
func ToMap[T any, TKey comparable, TValue any](slice []T, keySelector func(T) TKey, valueSelector func(T) TValue) map[TKey]TValue {
	dict := make(map[TKey]TValue, len(slice))

	for _, v := range slice {
		dict[keySelector(v)] = valueSelector(v)
	}

	return dict
}

// ToSlice returns a slice that was created by applying the provided selector to the key-value pairs of the provided map.
func ToSlice[TKey comparable, TValue any, TResult any](dict map[TKey]TValue, selector func(TKey, TValue) TResult) []TResult {
	var result []TResult
	for key, value := range dict {
		result = append(result, selector(key, value))
	}
	return result
}

// Where returns a slice that contains all the elements of the provided slice that satisfy the provided condition.
func Where[T any](slice []T, condition func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// Zip returns a slice that was created by applying the provided selector to each corresponding elements of both provided slices.
// Elements that don't have a corresponding element at their index are ignored.
func Zip[T1 any, T2 any, TResult any](first []T1, second []T2, selector func(T1, T2) TResult) []TResult {
	var result []TResult

	var shorterLength int
	if len(first) < len(second) {
		shorterLength = len(first)
	} else {
		shorterLength = len(second)
	}

	for i := 0; i < shorterLength; i++ {
		elementFirst := first[i]
		elementSecond := second[i]
		result = append(result, selector(elementFirst, elementSecond))
	}

	return result
}
