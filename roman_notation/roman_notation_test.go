package roman_notation

import "testing"

func TestRomanNumerals(t *testing.T) {
	testCases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			got := ConvertToRoman(testCase.Arabic)
			want := testCase.Want

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
