package square
//package main

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sideNumType int
const pi float64 = math.Pi
func CalcSquare(sideLen float64, sidesNum sideNumType) float64 {
	switch sidesNum{
	case 0:
		return pi * math.Pow(sideLen,2)
	case 3:
		return math.Sqrt(3)* math.Pow(sideLen,2)/4
	case 4:
		return  math.Pow(sideLen,2)/4
	default:
		return 0
	}
}
func main (){

}

