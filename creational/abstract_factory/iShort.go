package abstract_factory

type IShort interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Short struct {
	logo string
	size int
}

func (s *Short) SetLogo(logo string) {
	s.logo = logo
}

func (s *Short) GetLogo() string {
	return s.logo
}

func (s *Short) SetSize(size int) {
	s.size = size
}

func (s *Short) GetSize() int {
	return s.size
}
