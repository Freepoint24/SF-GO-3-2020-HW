package main

import (
	"fmt"
	"sf-golang-8-8-1/electronic"
)

func main() {
	applePhone := electronic.NewApplePhone("iPhone 12 Pro Max")
	androidPhone := electronic.NewAndroidPhone("Samsung", "Galaxy 100 S+")
	radioPhone := electronic.NewRadioPhone("Sanyo", "CLT-85", 16)

	printCharacteristics(applePhone)
	printCharacteristics(androidPhone)
	printCharacteristics(radioPhone)
}

func printCharacteristics(phone electronic.Phone) {
	result := fmt.Sprintf("Brand: %s, Model: %s, Type: %s", phone.Brand(), phone.Model(), phone.Type())

	switch phone.(type) {
	case electronic.Smartphone:
		result += fmt.Sprintf(", OS: %s", phone.(electronic.Smartphone).OS())
	case electronic.StationPhone:
		result += fmt.Sprintf(", Buttons: %d", phone.(electronic.StationPhone).ButtonsCount())
	default:
		fmt.Println("Unknown phone type")
		return
	}

	fmt.Println(result)
}
