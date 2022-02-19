package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"


type sidesCount int

const(
	SidesCircle sidesCount = iota
	_
	_
	SidesTriangle sidesCount = iota
	SidesSquare sidesCount = iota
)

func CalcSquare(sideLen float64, sidesNum sidesCount) float64 {
	var sq float64 = 0
	switch sidesNum {
	case SidesCircle: sq = math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle: sq = (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	case SidesSquare: sq = math.Pow(sideLen, 2)
	default: return 0
	}
	return sq
}

	 
// comment