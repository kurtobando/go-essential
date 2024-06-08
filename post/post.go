package post

type Post struct {
	Title       string
	Description string
}

func (p *Post) Get() string {
	return "Post: " + p.Title + " " + p.Description
}

func (p *Post) Set(t string, d string) {
	p.Title = t
	p.Description = d
}

func New(title string, desc string) *Post {
	return &Post{
		Title:       title,
		Description: desc,
	}
}
