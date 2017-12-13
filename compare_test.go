package compare_test

import (
	"github.com/kjljon/compare"
	"reflect"
	"testing"
)

/**
 * Mocks input
 */
var inputMock = map[string]interface{}{
	"profile": map[string]interface{}{
		"first": "Jon",
		"last":  "Doe",
	},
	"age": 20,
}

/**
 * Mocks compares
 */
var compareMock = []compare.GroupCompare{
	{
		Name: "Jon-child",
		Criteria: []compare.Criteria{
			{"/profile/first", compare.EQUAL, "Jon"},
			{"/age", "<", 18},
		},
	},
	{
		Name: "Jon-adult",
		Criteria: []compare.Criteria{
			{"/profile/first", compare.EQUAL, "Jon"},
			{"/age", ">", 18},
		},
	},
	{
		Name: "first-Jon",
		Criteria: []compare.Criteria{
			{"/profile/first", compare.EQUAL, "Jon"},
		},
	},
}

var matchTest = []struct {
	input  map[string]interface{}
	expect []string
	errors bool
}{
	{
		input: map[string]interface{}{
			"profile": map[string]interface{}{"first": "Jon"},
			"age":     20,
		},
		expect: []string{"Jon-adult", "first-Jon"},
		errors: false,
	}, {
		input: map[string]interface{}{
			"profile": map[string]interface{}{"first": "Joe"},
			"age":     20,
		},
		expect: []string{},
		errors: false,
	}, {
		input:  map[string]interface{}{},
		expect: []string{},
		errors: true,
	},
}

/**
 * test compare.MultipleGroups
 */
func TestMultipleGroups(t *testing.T) {
	for index, test := range matchTest {
		matches, err := compare.MultipleGroups(compareMock, test.input)
		if err != nil {
			if test.errors == false {
				t.Fatal(err)
			}
		} else if test.errors {
			t.Errorf("Expected to error but didn't [index: %v]: %v", index, err)
		} else if reflect.DeepEqual(matches, test.expect) == false {
			t.Errorf("Doesn't match [index: %v]: expected %v, actual %v", index, test.expect, matches)
		}
	}
}

/**
 * benchmarks compare.MultipleGroups
 */
func BenchmarkMultipleGroups(b *testing.B) {
	var matches []string
	var err error

	for n := 0; n < b.N; n++ {
		matches, err = compare.MultipleGroups(compareMock, inputMock)
		if err != nil {
			b.Fatal(err)
		}
	}

	matches = matches
}

//todo add test
//Group
//Eval
