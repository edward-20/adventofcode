package main
import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var y float64 = 1
	for z, zbefore := 1.0, x; math.Abs(z - zbefore) > 0.25; zbefore, z = z, z - (z*z- x) / (2*z) {
		y = z
	}
	return y
}
func main() {
	fmt.Println(sqrt(64))
}
