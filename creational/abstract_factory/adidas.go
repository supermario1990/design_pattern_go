package abstract_factory

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeShort() IShort {
	return &AdidasThirt{
		Short{
			logo: "adidas",
			size: 14,
		},
	}
}
