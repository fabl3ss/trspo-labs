package collatz

func Collatz(value int) []int {
	var values []int
	newValue := value
	values = append(values, newValue)
	for {
		if newValue != 1 {
			if newValue%2 == 0 {
				newValue = evenCalc(newValue)
			} else {
				newValue = oddCalc(newValue)
			}
			values = append(values, newValue)
		} else {
			break
		}
	}
	return values
}

func oddCalc(number int) int {
	return 3*number + 1
}

func evenCalc(number int) int {
	return number / 2
}
