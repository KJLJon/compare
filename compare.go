package compare

import (
	"errors"
	"github.com/mattn/go-jsonpointer"
	"strconv"
	"strings"
)

type Operator string

const (
	EQUAL              Operator = "="
	LESS_THAN          Operator = "<"
	LESS_THAN_EQUAL    Operator = "<="
	GREATER_THAN       Operator = ">"
	GREATER_THAN_EQUAL Operator = ">="
	NOT_EQUAL          Operator = "!="
	IN                 Operator = "in"
	NOT_IN             Operator = "not-in"
)

type GroupCompare struct {
	Name     string
	Criteria []Criteria
}

type Criteria struct {
	Key      string
	Operator Operator
	Compare  interface{}
}

// Checks if the input matches any groupings of criteria
// returns the matched group names
//
// if Criteria is the same for most of the Groups, it could be optimized
// by comaparing the unique criteria then finding which ones implement all the criteria
func MultipleGroups(groups []GroupCompare, input interface{}) ([]string, error) {
	var matched bool
	var err error
	matches := []string{}

	for _, group := range groups {
		matched, err = Group(group, input)
		if err != nil {
			return nil, err
		} else if matched {
			matches = append(matches, group.Name)
		}
	}

	return matches, nil
}

// checks a GroupCompare to see if the input matches the criteria
func Group(group GroupCompare, input interface{}) (bool, error) {
	matched, err := compareMultipleCriteria(group.Criteria, input)
	if err != nil {
		return false, err
	}

	return matched, nil
}

// Logical evaulator.  Takes two items and compares them based on the Operator passed in
func Eval(a interface{}, operator Operator, b interface{}) (bool, error) {
	var err error

	switch a.(type) {
	case float64, float32, int:
		var nbrA float64
		var nbrB float64

		nbrA, err = normalizeNumberValue(a)
		if err != nil {
			return false, err
		}

		nbrB, err = normalizeNumberValue(b)
		if err != nil {
			return false, err
		}

		return compareNumber(nbrA, operator, nbrB)
	case string:
		var strA string
		var strB string

		strA, err = normalizeStringValue(a)
		if err != nil {
			return false, err
		}

		strB, err = normalizeStringValue(b)
		if err != nil {
			return false, err
		}

		return compareString(strA, operator, strB)
	default:
		return false, errors.New("Unsupported Type")
	}
}

// compares multiple Criteria (runs Eval on each Criteria)
func compareMultipleCriteria(rules []Criteria, input interface{}) (bool, error) {
	var compare bool
	for _, criteria := range rules {
		value, err := jsonpointer.Get(input, criteria.Key)
		if err != nil {
			return false, err
		}

		compare, err = Eval(value, criteria.Operator, criteria.Compare)
		if err != nil {
			return compare, err
		}

		if compare == false {
			return false, nil
		}
	}

	return true, nil
}

// checks if an item is in the array
func inArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// logical compare two strings based on Operator provided
func compareString(str string, operator Operator, compare string) (bool, error) {
	switch operator {
	case EQUAL:
		return str == compare, nil
	case NOT_EQUAL:
		return str != compare, nil
	case NOT_IN:
		return !inArray(str, strings.Split(compare, ",")), nil
	case IN:
		return inArray(str, strings.Split(compare, ",")), nil
	default:
		return false, errors.New("Unsupported operator")
	}
}

// logical compare two numbers based on Operator provided
func compareNumber(nbr float64, operator Operator, compare float64) (bool, error) {
	switch operator {
	case EQUAL:
		return nbr == compare, nil
	case NOT_EQUAL:
		return nbr != compare, nil
	case GREATER_THAN:
		return nbr > compare, nil
	case GREATER_THAN_EQUAL:
		return nbr >= compare, nil
	case LESS_THAN:
		return nbr < compare, nil
	case LESS_THAN_EQUAL:
		return nbr <= compare, nil
	default:
		return false, errors.New("Unsupported operator")
	}
}

// normalizes an interface into a string
func normalizeStringValue(value interface{}) (string, error) {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int)), nil
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64), nil
	case string:
		return value.(string), nil
	default:
		return "", errors.New("Unsupported string type")
	}
}

// normalizes an interface into a number
func normalizeNumberValue(value interface{}) (float64, error) {
	switch value.(type) {
	case int:
		return float64(value.(int)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return value.(float64), nil
	case string:
		return strconv.ParseFloat(value.(string), 64)
	default:
		return 0, errors.New("Unsupported number type")
	}
}
