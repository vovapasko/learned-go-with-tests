package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	testCases := []struct {
		TestName      string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			TestName: "test struct with one field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			var got []string
			x := testCase.Input

			Walk(x, func(input string) {
				got = append(got, input)
			})

			if len(got) != len(testCase.ExpectedCalls) {
				t.Errorf("Expected got = %d", len(got))
			}
			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("Expected %s but got %s", testCase.ExpectedCalls, got)
			}
		})
	}

}
