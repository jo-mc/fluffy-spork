package main

import (
	"fmt"
	"math"
)

type RollsT struct {
	n              uint64
	M1, M2, M3, M4 float64
}

func (st *RollsT) pushs(x float64) {
	n1 := st.n
	st.n++
	delta := x - st.M1
	deltaN := delta / float64(st.n)
	deltaN2 := deltaN * deltaN
	term1 := delta * deltaN * float64(n1)
	st.M1 += deltaN
	st.M4 += term1*deltaN2*(float64(st.n*st.n)-float64(3*st.n+3)) + 6.0*deltaN2*st.M2 - 4.0*deltaN*st.M3
	st.M3 += term1*deltaN*(float64(st.n-2)) - 3*deltaN*st.M2
	st.M2 += term1
}

func main() {
	fmt.Printf("hello, world\n")
	var sdats = RollsT{}
	sdats.pushs(2.0)
	sdats.pushs(4.0)
	sdats.pushs(5.0)
	sdats.pushs(4.0)
	sdats.pushs(6.0)
	fmt.Println(sdats)
	fmt.Println("mean ", sdats.M1)
	fmt.Println("variance ", sdats.M2/float64(sdats.n-1))
	fmt.Println("standard deviation ", math.Sqrt(sdats.M2/float64(sdats.n-1)))
	fmt.Println("skewness ", math.Sqrt(float64(sdats.n))*sdats.M3/math.Pow(sdats.M2, 1.5))
	fmt.Println("kurtosis ", float64(sdats.n)*sdats.M4/(sdats.M2*sdats.M2)-3.0)

}
