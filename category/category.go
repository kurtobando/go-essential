package category

type Category struct {
	Title       string
	Description string
}

func (p *Category) Get() string {
	return "Category: " + p.Title + " " + p.Description
}

func (p *Category) Set(t string, d string) {
	p.Title = t
	p.Description = d
}

func New(title string, desc string) *Category {
	return &Category{
		Title:       title,
		Description: desc,
	}
}
