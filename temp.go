package temperature

import "fmt"

func Converter(x float64, y string) (float64, error) {

	if y == "C" || y == "c" {
		return tempToF(x), nil
	} else if y == "F" || y == "f" {
		return tempToC(x), nil
	}
	return 0, fmt.Errorf("unrecognized symbol: %s, want: F, f, C or c", y)

}

func tempToC(a float64) float64 {
	return (a - 32) * 5 / 9

}

func tempToF(b float64) float64 {
	return (b * 1.8) - 32
}
