package automobile

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type auto struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

func (a auto) Brand() string {
	return a.brand
}

func (a auto) Model() string {
	return a.model
}

func (a auto) Dimensions() Dimensions {
	return a.dimensions
}

func (a auto) MaxSpeed() int {
	return a.maxSpeed
}

func (a auto) EnginePower() int {
	return a.enginePower
}

func newEuAuto(brand, model string, dim Dimensions, maxSpeed, enginePower int) Auto {
	return auto{
		brand,
		model,
		NewCMDimensions(dim.Length().Get(CM), dim.Width().Get(CM), dim.Height().Get(CM)),
		maxSpeed,
		enginePower,
	}
}

func newUSAuto(brand, model string, dim Dimensions, maxSpeed, enginePower int) Auto {
	return auto{
		brand,
		model,
		NewInchDimensions(dim.Length().Get(Inch), dim.Width().Get(Inch), dim.Height().Get(Inch)),
		maxSpeed,
		enginePower,
	}
}

type BMWAuto struct {
	auto
}

func NewBMWAuto(model string, dim Dimensions, maxSpeed, enginePower int) Auto {
	return newEuAuto("BMW", model, dim, maxSpeed, enginePower)
}

type MercedesAuto struct {
	auto
}

func NewMercedesAuto(model string, dim Dimensions, maxSpeed, enginePower int) Auto {
	return newEuAuto("Mercedes", model, dim, maxSpeed, enginePower)
}

type DodgeAuto struct {
	auto
}

func NewDodgeAuto(model string, dim Dimensions, maxSpeed, enginePower int) Auto {
	return newUSAuto("Dodge", model, dim, maxSpeed, enginePower)
}
