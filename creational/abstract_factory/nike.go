package abstract_factory

type Nike struct {
}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) MakeShort() IShort {
	return &NikeShort{
		Short{
			logo: "nike",
			size: 14,
		},
	}
}
