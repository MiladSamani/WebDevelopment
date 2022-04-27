package views

import "html/template"

type View struct {
	Template *template.Template
	Layout   string
}

// NewView files is a slice of filenames
func NewView(layout string, files ...string) *View {
	files = append(files, "views/layouts/Footer.gohtml", "views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
