package number

import (
	"fmt"
	"strconv"
)

var roundMap map[uint8]string = map[uint8]string{0: "%.0f", 1: "%.1f", 2: "%.2f", 3: "%.3f", 4: "%.4f", 5: "%.5f", 6: "%.6f", 7: "%.7f", 8: "%.8f", 9: "%.9f", 10: "%.10f", 11: "%.11f", 12: "%.12f", 13: "%.13f", 14: "%.14f", 15: "%.15f", 16: "%.16f"}

// digit should <= 16
func Round(num float64, digit uint8) float64 {
	if digit > 16 {
		return num
	}
	floatStr := fmt.Sprintf(roundMap[digit], num)
	result, _ := strconv.ParseFloat(floatStr, 64)
	return result
}
