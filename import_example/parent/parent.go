package parent

type Father struct {
	Name string
}

func (f *Father) Data(name string) string {
	f.Name = name
	return f.Name
}
