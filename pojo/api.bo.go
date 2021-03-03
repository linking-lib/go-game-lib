package pojo

type ApiBO struct {
	Module string
	Name   string
}

func (a ApiBO) GetFullName() string {
	return a.Module + "." + a.Name
}
