package entities

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{Name: name}
}
