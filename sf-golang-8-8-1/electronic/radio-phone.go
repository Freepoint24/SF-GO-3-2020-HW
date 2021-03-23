package electronic

type radioPhone struct {
	phone
	buttonsCount int
}

func NewRadioPhone(brand, model string, buttonsCount int) radioPhone {
	return radioPhone{
		phone:        phone{brand: brand, model: model, phoneType: phoneTypeStation},
		buttonsCount: buttonsCount,
	}
}

func (p radioPhone) ButtonsCount() int {
	return p.buttonsCount
}
