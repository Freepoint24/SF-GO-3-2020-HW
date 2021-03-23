package electronic

type applePhone struct {
	phone
}

func NewApplePhone(model string) applePhone {
	return applePhone{phone{brand: "Apple", model: model, phoneType: phoneTypeSmartphone}}
}

func (p applePhone) OS() string {
	return "iOS"
}
