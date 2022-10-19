func getChar(num int) string {
	switch num {
	case 1:
		return "I"
	case 5:
		return "V"
	case 10:
		return "X"
	case 50:
		return "L"
	case 100:
		return "C"
	case 500:
		return "D"
	case 1000:
		return "M"
	case 4:
		return "IV"
	case 9:
		return "IX"
	case 40:
		return "XL"
	case 90:
		return "XC"
	case 400:
		return "CD"
	case 900:
		return "CM"
	}
	return ""
}
func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	var value string
	for num > 0 {
		var caseNums = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
		for i := len(caseNums) - 1; i >= 0; i-- {
			if num >= caseNums[i] {
				value += getChar(caseNums[i])
				num -= caseNums[i]
				break
			}
		}
	}

	return value
}