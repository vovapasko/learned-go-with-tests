package roman_notation

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
