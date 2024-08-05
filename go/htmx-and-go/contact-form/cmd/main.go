package main

import (
	"html/template"
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GeneralTemplate struct {
	tpl *template.Template
}

func (t *GeneralTemplate) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.tpl.ExecuteTemplate(w, name, data)
}

func newGeneralTemplate() *GeneralTemplate {
	return &GeneralTemplate{
		tpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Page struct {
	Form FormData
	List ListData
}

func newPage() Page {
	return Page{
		Form: newFormData(),
		List: newListData(),
	}
}

var contactIdCounter int = 1

type Contact struct {
	Id    int
	Name  string
	Email string
}

func newContact(name, email string) Contact {
	contactIdCounter++
	return Contact{
		Id:    contactIdCounter,
		Name:  name,
		Email: email,
	}
}

type ListData struct {
	Contacts []Contact
}

func newListData() ListData {
	return ListData{
		Contacts: []Contact{
			// {Name: "John Doe", Email: "john@doe.com"},
			// {Name: "John Wood", Email: "john@wood.com"},
		},
	}
}

func (d *ListData) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func main() {
	e := echo.New()
	e.Static("/images", "images")
	e.Static("/css", "css")
	e.Renderer = newGeneralTemplate()
	e.Use(middleware.Logger())
	page := newPage()

	e.GET("/", func(ctx echo.Context) error {
		page.Form = newFormData()
		return ctx.Render(200, "index", page)
	})

	e.POST("/contacts", func(ctx echo.Context) error {
		params, err := ctx.FormParams()
		if err != nil {
			panic(err)
		}

		name := params.Get("name")
		email := params.Get("email")

		if page.List.hasEmail(email) {
			form := &page.Form
			form.Values["name"] = name
			form.Values["email"] = email
			form.Errors["email"] = "Email already exists"
			return ctx.Render(422, "contact-form", form)
		}

		contact := newContact(name, email)
		page.List.Contacts = append(page.List.Contacts, contact)
		ctx.Render(200, "contact-form", newFormData())
		return ctx.Render(200, "oob-contact-list", page.List)
	})

	e.DELETE("/contacts/:id", func(ctx echo.Context) error {
		time.Sleep(3 * time.Second) // simulate work

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)

		if err != nil {
			return ctx.String(400, "Id must be an integer")
		}

		isDeleted := false
		for idx, contact := range page.List.Contacts {
			if contact.Id == id {
				page.List.Contacts = append(page.List.Contacts[:idx], page.List.Contacts[idx+1:]...)
				isDeleted = true
				break
			}
		}

		if !isDeleted {
			return ctx.String(400, "Contact not found")
		}

		return ctx.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
