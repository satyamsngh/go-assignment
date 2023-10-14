package temperature

var c int

func Temperature(f int) int {
	c = (f - 32) * 5 / 9

	return c
}
