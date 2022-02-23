package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type FigureSides int

const (
	SidesTriangle FigureSides = 0
	SidesSquare               = 3
	SidesCircle               = 4
)

func calcTriangleSquare(sideLen float64) float64 {
	var square float64 = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4

	return square
}

func calcSquareSquare(sideLen float64) float64 {
	var square float64 = math.Pow(sideLen, 2)

	return square
}

func calcCircleSquare(sideLen float64) float64 {
	var square float64 = math.Pi * math.Pow(sideLen, 2)

	return square
}

func CalcSquare(sideLen float64, sidesNum FigureSides) float64 {
	switch sidesNum {
	case SidesTriangle:
		return calcTriangleSquare(sideLen)
	case SidesSquare:
		return calcSquareSquare(sideLen)
	case SidesCircle:
		return calcCircleSquare(sideLen)
	default:
		return 0
	}
}
