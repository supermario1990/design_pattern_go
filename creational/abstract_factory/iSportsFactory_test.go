package abstract_factory_test

import (
	"design_pattern/creational/abstract_factory"
	"testing"
)

func TestAdidas(t *testing.T) {
	adidasFactory, _ := abstract_factory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	logo := adidasShoe.GetLogo()
	if logo == "adidas" {
		t.Logf("adidas shoe logo = %v, want %v", logo, "adidas")
	} else {
		t.Errorf("adidas shoe logo = %v, want %v", logo, "adidas")
	}

	size := adidasShoe.GetSize()
	if size == 14 {
		t.Logf("adidas shoe size = %v, want %v", size, 14)
	} else {
		t.Errorf("adidas shoe size = %v, want %v", size, 14)
	}
}

func TestNike(t *testing.T) {
	nikeFactory, _ := abstract_factory.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	logo := nikeShoe.GetLogo()
	if logo == "nike" {
		t.Logf("nike shoe logo = %v, want %v", logo, "nike")
	} else {
		t.Errorf("nike shoe logo = %v, want %v", logo, "nike")
	}

	size := nikeShoe.GetSize()
	if size == 14 {
		t.Logf("nike shoe size = %v, want %v", size, 14)
	} else {
		t.Errorf("nike shoe size = %v, want %v", size, 14)
	}
}
