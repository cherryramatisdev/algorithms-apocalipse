package two_crystal_ball

import "math"

func TwoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && j < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}

	return -1
}
