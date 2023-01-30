package reflection

import (
	"reflect"
	"testing"
)

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
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			var got []string
			x := testCase.Input

			Walk(x, func(input string) {
				got = append(got, input)
			})

			if len(got) != 1 {
				t.Errorf("Expected got = %d", len(got))
			}
			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("Expected %s but got %s", testCase.ExpectedCalls, got)
			}
		})
	}

}
