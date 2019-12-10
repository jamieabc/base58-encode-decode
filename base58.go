package base58

import (
	"math"
	"strings"
)

const (
	mapping = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func Encode(str string) string {
	sum := float64(0)

	// ascii value * 2^(index * 8)
	for i, r := range str {
		sum += float64(int(r)) * math.Pow(float64(2), float64((len(str)-1-i)*8))
	}

	// modulo
	reversed := make([]int, 0)
	count := 0
	for sum >= 1 {
		reversed = append(reversed, int(int64(sum)%int64(58)))
		sum /= 58
		count++
	}

	var sb strings.Builder
	for i := count - 1; i >= 0; i-- {
		sb.WriteString(string(mapping[reversed[i]]))
	}
	return sb.String()
}

func Decode(str string) string {
	length := len(str)
	sum := float64(0)
	var j int
	for i, r := range str {
		for j = 0; j < len(mapping); j++ {
			if mapping[j] == uint8(r) {
				break
			}
		}
		sum += float64(j) * math.Pow(float64(58), float64(length-1-i))
	}

	digitLength := 0
	for sum/math.Pow(float64(2), float64(8*digitLength)) >= 1 {
		digitLength++
	}

	ordered := make([]int, digitLength)
	for i := 0; i < digitLength; i++ {
		divider := math.Pow(float64(2), float64(8*(digitLength-1-i)))
		modulo := int(int64(sum) % int64(divider))
		ordered[i] = int(int64(sum) / int64(divider))
		sum = float64(modulo)
	}

	var sb strings.Builder
	for _, r := range ordered {
		sb.WriteRune(rune(r))
	}
	return sb.String()
}
