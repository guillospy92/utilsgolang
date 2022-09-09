package render

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetView    *jet.Set
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (c *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(c.Renderer) {
	case "go":
		return c.GoPage(w, r, view, data)
	case "jet":
		return c.JetPage(w, r, view, variables, data)
	}

	return nil
}

func (c *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.html", c.RootPath, view))

	if err != nil {
		return nil
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)

	if err != nil {
		return err
	}

	return nil
}

// JetPage renders a template using jet engine
func (c *Render) JetPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {

	vars := make(jet.VarMap)

	if variables != nil {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	t, err := c.JetView.GetTemplate(fmt.Sprintf("%s.jet", view))

	if err != nil {
		return err
	}

	if err := t.Execute(w, vars, &td); err != nil {
		return err
	}

	return nil
}
