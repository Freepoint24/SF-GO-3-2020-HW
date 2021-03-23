package automobile

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDodge(t *testing.T) {
	cmDims := NewCMDimensions(485, 184.3, 149.7)
	wantDims := NewInchDimensions(485/InchCMRatio, 184.3/InchCMRatio, 149.7/InchCMRatio)
	wantType := reflect.TypeOf(wantDims)

	car := NewDodgeAuto("Avenger", cmDims, 200, 164)
	factDims := car.Dimensions()
	factType := reflect.TypeOf(factDims)

	fmt.Printf("%+v\n", car)

	// check type
	if wantType != factType {
		t.Fatalf("Dodge Avenger Dimensions() type => %s; want => %s", factType, wantType)
	}

	// check values
	if wantDims != factDims {
		t.Fatalf("Dodge Avenger Dimensions() => %+v; want => %+v", factDims, wantDims)
	}

}

func TestBMW(t *testing.T) {
	inchDims := NewInchDimensions(492.2/InchCMRatio, 200/InchCMRatio, 174.5/InchCMRatio)
	wantDims := NewCMDimensions(492.2, 200, 174.5)
	wantType := reflect.TypeOf(wantDims)

	car := NewBMWAuto("X5", inchDims, 230, 265)
	factDims := car.Dimensions()
	factType := reflect.TypeOf(factDims)

	fmt.Printf("%+v\n", car)

	// check type
	if wantType != factType {
		t.Fatalf("BMW X5 Dimensions() type => %s; want => %s", factType, wantType)
	}

	// check values
	if wantDims != factDims {
		t.Fatalf("BMW X5 Dimensions() => %+v; want => %+v", factDims, wantDims)
	}

}

func TestMercedes(t *testing.T) {
	inchDims := NewInchDimensions(504/InchCMRatio, 185/InchCMRatio, 174.5/InchCMRatio)
	wantDims := NewCMDimensions(504, 185, 174.5)
	wantType := reflect.TypeOf(wantDims)

	car := NewMercedesAuto("W220", inchDims, 250, 242)
	factDims := car.Dimensions()
	factType := reflect.TypeOf(factDims)

	fmt.Printf("%+v\n", car)

	// check type
	if wantType != factType {
		t.Fatalf("Mercedes W220 Dimensions() type => %s; want => %s", factType, wantType)
	}

	// check values
	if wantDims != factDims {
		t.Fatalf("Mercedes W220 Dimensions() => %+v; want => %+v", factDims, wantDims)
	}

}
