package abstract_factory

import "fmt"

type iSportFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

func GetSportsFactory(brand string) (iSportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
