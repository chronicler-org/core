package appUtil

func CalculatePercentVariation(currentValue, lastValue float64) float64 {
	var percentVariation float64

	if lastValue != 0 {
		percentVariation = (currentValue - lastValue) / lastValue * 100
	} else {
		if currentValue == 0 {
			percentVariation = 0
		} else {
			percentVariation = 100
		}
	}

	return percentVariation
}
