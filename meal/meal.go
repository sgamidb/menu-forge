package meal

type Meal struct {
	Name    string
	hasMeat bool
}

func (m *Meal) IsVegan() bool {
	return m.hasMeat == false
}

func New(name string, hasMeat bool) *Meal {
	return &Meal{Name: name, hasMeat: hasMeat}
}
