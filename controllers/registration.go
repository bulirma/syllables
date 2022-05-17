package controllers

import (
	//"fmt"
	"net/http"

	"github.com/bulirma/syllables/services"
	"github.com/bulirma/syllables/services/validation"
	"github.com/bulirma/syllables/views"
)

type RegistrationController struct {
	tokenMgr services.TokenManager
}

func NewRegistrationController(tokenMgr services.TokenManager) RegistrationController {
	regCtr := RegistrationController {
		tokenMgr: tokenMgr,
	}
	return regCtr
}

func (regCtr *RegistrationController) RegistrationFormGet(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	tokens, present := params["token"]
	if !present || len(tokens) != 1 || len(tokens[0]) == 0 {
		tokenPrompt := views.TokenPromptView {}
		tokenPrompt.Render(w)
		return
	}
	if !regCtr.tokenMgr.IsTokenValid(tokens[0]) {
		tokenPrompt := views.TokenPromptView {
			ErrorMessage: "Token is invalid.",
		}
		tokenPrompt.Render(w)
		return
	}
	regView := views.RegistrationFormView {
		Token: tokens[0],
	}
	regView.Render(w)
}

func (regCtr *RegistrationController) RegistrationFormPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := r.FormValue("token")
	if !regCtr.tokenMgr.IsTokenValid(token) {
		tokenPrompt := views.TokenPromptView {
			ErrorMessage: "Registration time exceeded. Token is now invalid.",
		}
		tokenPrompt.Render(w)
		return
	}
	username := r.FormValue("username")
	var (
		validationResult uint
		errorMsg string
	)
	validationResult = validation.ValidateUsername(username)
	if validationResult > 0 {
		switch validationResult {
		case validation.UsernameTooShort:
			errorMsg = "Username should be at least 1 character long."
		case validation.UsernameTooLong:
			errorMsg = "Username is too long, 24 characters should be sufficient."
		case validation.UsernameInvalidCharacter:
			errorMsg = "Username should contain at least one from these characters: a-z, A-Z, 0-9 or hyphen, underscore and dot."
		}
		regView := views.RegistrationFormView {
			Token: token,
			ErrorMessage: errorMsg,
		}
		regView.Render(w)
		return
	}
	password := r.FormValue("password")
	validationResult = validation.ValidatePassword(password)
	if validationResult > 0 {
		switch validationResult {
		case validation.PasswordTooShort:
			errorMsg = "Username should be at least 1 character long."
		case validation.PasswordTooLong:
			errorMsg = "Username is too long, 24 characters should be sufficient."
		case validation.PasswordInvalidCharacter:
			errorMsg = "Username should contain at least one from these characters: a-z, A-Z, 0-9 or hyphen, underscore and dot."
		}
		regView := views.RegistrationFormView {
			Token: token,
			Username: username,
			ErrorMessage: errorMsg,
		}
		regView.Render(w)
		return
	}
	passwordConf := r.FormValue("password_confirm")
	if password != passwordConf {
		regView := views.RegistrationFormView {
			Token: token,
			Username: username,
			ErrorMessage: "Passwords don't match.",
		}
		regView.Render(w)
		return
	}
	ok := services.ProsodyRegister(username, password)
	if !ok {
		regView := views.RegistrationFormView {
			Token: token,
			ErrorMessage: "Username is already taken.",
		}
		regView.Render(w)
		return
	}
	http.Redirect(w, r, "/registration-complete", http.StatusSeeOther)
}
