package controllers

import (
	"GoWeb/app"
	"GoWeb/models"
	"GoWeb/security"
	"GoWeb/templating"
	"net/http"
)

// GetController is a wrapper struct for the App struct
type GetController struct {
	App *app.App
}

func (getController *GetController) ShowHome(w http.ResponseWriter, r *http.Request) {
	type dataStruct struct {
		Test string
	}

	user, err := models.GetCurrentUser(getController.App, r)
	if err != nil {
		return
	}

	data := dataStruct{
		Test: "Selamat datang " + user.Username,
	}

	templating.RenderTemplate(getController.App, w, "templates/pages/home.html", data)
}

func (getController *GetController) ShowRegister(w http.ResponseWriter, r *http.Request) {
	type dataStruct struct {
		CsrfToken string
	}

	// Create csrf token
	CsrfToken, err := security.GenerateCsrfToken(w, r)
	if err != nil {
		return
	}

	data := dataStruct{
		CsrfToken: CsrfToken,
	}

	templating.RenderTemplate(getController.App, w, "templates/pages/register.html", data)
}

func (getController *GetController) ShowLogin(w http.ResponseWriter, r *http.Request) {
	type dataStruct struct {
		CsrfToken string
	}

	// Create csrf token
	CsrfToken, err := security.GenerateCsrfToken(w, r)
	if err != nil {
		return
	}

	data := dataStruct{
		CsrfToken: CsrfToken,
	}

	templating.RenderTemplate(getController.App, w, "templates/pages/login.html", data)
}

func (getController *GetController) Logout(w http.ResponseWriter, r *http.Request) {
	models.LogoutUser(getController.App, w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
