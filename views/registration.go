package views

import (
	"html/template"
	"io"
)

type TokenPromptView struct {
	ErrorMessage string
}

func (view *TokenPromptView) Render(w io.Writer) {
	t := template.Must(template.ParseFiles("templates/token-prompt.html"))
	t.Execute(w, view)
}

type RegistrationFormView struct {
	Token string
	ErrorMessage string
	Username string
}

func (view *RegistrationFormView) Render(w io.Writer) {
	t := template.Must(template.ParseFiles("templates/registration-form.html"))
	t.Execute(w, view)
}
