package render

import (
	"github.com/edenofjinx/go-bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	request, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(request.Context(), "flash", "123")

	result := AddDefaultData(&td, request)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	cache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = cache
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
	err = RenderTemplate(&ww, r, "asdf.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}
}

func getSession() (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := request.Context()
	ctx, _ = session.Load(ctx, request.Header.Get("X-Session"))
	request = request.WithContext(ctx)
	return request, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}