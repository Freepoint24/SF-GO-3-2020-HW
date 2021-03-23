package electronic

type androidPhone struct {
	phone
}

func NewAndroidPhone(brand, model string) androidPhone {
	return androidPhone{phone{brand: brand, model: model, phoneType: phoneTypeSmartphone}}
}

func (p androidPhone) OS() string {
	return "Android"
}
