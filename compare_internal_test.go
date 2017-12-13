package compare

import (
	"testing"
)

/**
 * Mock test for TestInArray
 */
var inArrayTest = []struct {
	list   []string
	a      string
	expect bool
}{
	{[]string{"foo", "bar"}, "foo", true},
	{[]string{"foo", "bar"}, "bar", true},
	{[]string{"foo", "bar"}, "boo", false},
}

/**
 * Mock test for TestCompareString
 */
var compareStringTest = []struct {
	str      string
	operator Operator
	compare  string
	expect   bool
}{
	{"foo", EQUAL, "foo", true},
	{"foo", EQUAL, "bar", false},
	{"foo", NOT_EQUAL, "foo", false},
	{"foo", NOT_EQUAL, "bar", true},
	{"foo", IN, "foo,bar", true},
	{"boo", IN, "foo,bar", false},
	{"foo", NOT_IN, "foo,bar", false},
	{"boo", NOT_IN, "foo,bar", true},
}

/**
 * Mock test for TestCompareNumber
 */
var compareNumberTest = []struct {
	nbr      float64
	operator Operator
	compare  float64
	expect   bool
}{
	//test operators with integers
	{1, EQUAL, 1, true},
	{1, EQUAL, 2, false},
	{1, NOT_EQUAL, 1, false},
	{1, NOT_EQUAL, 2, true},
	{1, GREATER_THAN, 0, true},
	{1, GREATER_THAN, 1, false},
	{1, GREATER_THAN, 2, false},
	{1, LESS_THAN, 0, false},
	{1, LESS_THAN, 1, false},
	{1, GREATER_THAN_EQUAL, 0, true},
	{1, GREATER_THAN_EQUAL, 1, true},
	{1, GREATER_THAN_EQUAL, 2, false},
	{1, LESS_THAN_EQUAL, 0, false},
	{1, LESS_THAN_EQUAL, 1, true},
	{1, LESS_THAN_EQUAL, 2, true},
	//test floats
	{3.1415, NOT_EQUAL, 3.14159, true},
	{3.1415, LESS_THAN, 3.14159, true},
	{3.1415, GREATER_THAN, 3.14159, false},
}

/**
 * Tests inArray functionality
 */
func TestInArray(t *testing.T) {
	for _, test := range inArrayTest {
		actual := inArray(test.a, test.list)
		if actual != test.expect {
			t.Errorf("inArray(%v, %v): expected %v, actual %v", test.a, test.list, test.expect, actual)
		}
	}
}

/**
 * Tests compareString functionality
 */
func TestCompareString(t *testing.T) {
	for _, test := range compareStringTest {
		actual, err := compareString(test.str, test.operator, test.compare)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if actual != test.expect {
			t.Errorf("compareString(%v, %v, %v): expected %v, actual %v", test.str, test.operator, test.compare, test.expect, actual)
		}
	}
}

/**
 * Tests compareNumber functionality
 */
func TestCompareNumber(t *testing.T) {
	for _, test := range compareNumberTest {
		actual, err := compareNumber(test.nbr, test.operator, test.compare)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if actual != test.expect {
			t.Errorf("compareNumber(%v, %v, %v): expected %v, actual %v", test.nbr, test.operator, test.compare, test.expect, actual)
		}
	}
}

//todo add tests for the following:
//compareMultipleCriteria
//normalizeStringValue
//normalizeNumberValue
