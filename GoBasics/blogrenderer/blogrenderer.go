package blogrenderer

import (
	"embed"
	blogposts "github.com/dawiddzhafarov/GoProjects/GoBasics/reading_files"
	"html/template"
	"io"
)

type PostRenderer struct {
	templ *template.Template
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

//NewPostRenderer parses templates once for all the posts, instead of repeating it
// giving improvement in benchmarks
func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", post)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
