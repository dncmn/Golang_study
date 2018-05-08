package ccc

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Test_Hello(T *testing.T) {

	var n float64 = 0.65397215
	res := fmt.Sprintf("%0.4f", n)
	fmt.Println(res)

	fmt.Println(strconv.ParseFloat(res, 64))

}
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
