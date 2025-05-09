package test_utils

import (
	"fmt"
	"reflect"
)

type BasicTestCase[T, S any] struct {
	Value    T
	Expected S
}

type ChainedActionTestCase[T, S any] struct {
	Name     string
	Action   func(*T)
	Expected S
}

type ChainedActionReturnTestCase[T, S any] struct {
	Name     string
	Action   func(*T) S
	Expected S
}

type ActionReturnTestCase[T, S any] struct {
	Name     string
	Action   func() S
	Expected S
}

type testResultsObject struct {
	PASSED string
	FAILED string
}

var TestResults = testResultsObject{
	PASSED: "\033[32mpassed\033[0m",
	FAILED: "\033[31mfail\033[0m",
}

func ToEqual[T comparable](value, expected T) string {
	if value != expected {
		return fmt.Sprint(TestResults.FAILED)
	} else {
		return fmt.Sprint(TestResults.PASSED)
	}
}

func SlicesToEqual[T comparable](value, expected []T) string {
	if reflect.DeepEqual(value, expected) {
		return fmt.Sprint(TestResults.PASSED)
	} else {
		return fmt.Sprint(TestResults.FAILED)
	}
}
