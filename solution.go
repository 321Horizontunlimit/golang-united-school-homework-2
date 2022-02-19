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

func squareCircle(radius float64) float64{
	return math.Pi * math.Pow(radius, 2)
}

func squareTriangle(side float64) float64{
	return (math.Sqrt(3) / 4) * math.Pow(side, 2)
}

func squareBox(side float64) float64{
	return math.Pow(side, 2)
}


func CalcSquare(sideLen float64, sidesNum sidesCount) float64 {
	switch sidesNum {
	case SidesCircle: return squareCircle(sideLen)
	case SidesTriangle: return squareTriangle(sideLen)
	case SidesSquare: return squareBox(sideLen)
	default: return 0
	}
}

	 
// comment