package meal

type Meal struct {
	Name string
}

func New(s string) *Meal {
	return &Meal{Name: s}
}
