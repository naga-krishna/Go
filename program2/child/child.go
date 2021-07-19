package child

type Son struct {
	Name string
}

func (s *Son) Data(name string) string {
	s.Name = name
	return s.Name
}
