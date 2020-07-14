package iBuilder_test

import (
	"design_pattern/creational/iBuilder"
	"testing"
)

func TestNormalBuilder(t *testing.T) {
	normalBuilder := iBuilder.GetBuilder("normal")
	director := iBuilder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	if normalHouse.Floor == 2 {
		t.Logf("Normal House Num Floor: %d\n", normalHouse.Floor)
	} else {
		t.Errorf("Normal House Num Floor: %d\n", normalHouse.Floor)
	}

	if normalHouse.WindowType == "Wooden Window" {
		t.Logf("Normal House Window Type: %s\n", normalHouse.WindowType)
	} else {
		t.Errorf("Normal House Window Type: %s\n", normalHouse.WindowType)
	}

	if normalHouse.DoorType == "Wooden Door" {
		t.Logf("Normal House Door Type: %s\n", normalHouse.DoorType)
	} else {
		t.Errorf("Normal House Door Type: %s\n", normalHouse.DoorType)
	}
}

func TestIglooBuilder(t *testing.T) {
	iglooBuilder := iBuilder.GetBuilder("igloo")
	director := iBuilder.NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse()

	if iglooHouse.Floor == 1 {
		t.Logf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
	} else {
		t.Errorf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
	}

	if iglooHouse.WindowType == "Snow Window" {
		t.Logf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	} else {
		t.Errorf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	}

	if iglooHouse.DoorType == "Snow Door" {
		t.Logf("Igloo House Door Type: %s\n", iglooHouse.DoorType)
	} else {
		t.Errorf("Igloo House Door Type: %s\n", iglooHouse.DoorType)
	}
}
