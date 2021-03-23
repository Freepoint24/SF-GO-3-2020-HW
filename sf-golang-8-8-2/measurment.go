package automobile

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

const InchCMRatio float64 = 2.54

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch {
		case u.T == Inch && t == CM: // in => cm
			value *= InchCMRatio
		case u.T == CM && t == Inch: // cm => in
			value /= InchCMRatio
		default:
			panic("unknown measurement")
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type InchDimensions struct {
	length float64
	width  float64
	height float64
}

func NewInchDimensions(length, width, height float64) InchDimensions {
	return InchDimensions{length, width, height}
}

func (d InchDimensions) Length() Unit {
	return Unit{d.length, Inch}
}

func (d InchDimensions) Width() Unit {
	return Unit{d.width, Inch}
}

func (d InchDimensions) Height() Unit {
	return Unit{d.height, Inch}
}

type CMDimensions struct {
	length float64
	width  float64
	height float64
}

func NewCMDimensions(length, width, height float64) CMDimensions {
	return CMDimensions{length, width, height}
}

func (d CMDimensions) Length() Unit {
	return Unit{d.length, CM}
}

func (d CMDimensions) Width() Unit {
	return Unit{d.width, CM}
}

func (d CMDimensions) Height() Unit {
	return Unit{d.height, CM}
}
