func romanToInt(s string) int {
	const (
		I = 1
		V = 5
		X = 10
		L = 50
		C = 100
		D = 500
		M = 1000
	)
	sLength := len(s)
	if sLength < 1 || sLength > 15 {
		return -1
	}

	var value int

	for i := 0; i < sLength; i++ {
		firstChar := string(s[i])
		if i+1 < sLength {
			secondChar := string(s[i+1])
			if firstChar == "I" {
				switch secondChar {
				case "V":
					value += 4
					i++
					continue
				case "X":
					value += 9
					i++
					continue
				}
			}
			if firstChar == "X" {
				switch secondChar {
				case "L":
					value += 40
					i++
					continue
				case "C":
					value += 90
					i++
					continue
				}
			}
			if firstChar == "C" {
				switch secondChar {
				case "D":
					value += 400
					i++
					continue
				case "M":
					value += 900
					i++
					continue
				}
			}
		}
		switch firstChar {
		case "I":
			value += I
		case "V":
			value += V
		case "X":
			value += X
		case "L":
			value += L
		case "C":
			value += C
		case "D":
			value += D
		case "M":
			value += M
		default:
			return -1
		}
	}

	if value < 1 || value > 3999 {
		return -1
	}
	return value
}