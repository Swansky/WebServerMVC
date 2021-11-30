package controller

import (
	"awesomeProject1/models"
	"awesomeProject1/repositories"
	"awesomeProject1/server"
	"awesomeProject1/utils"
	"awesomeProject1/view"
	"net/http"
)

type ProfileViewData struct {
	RegisterError      string
	ResetPasswordError string
}

func Profile(w http.ResponseWriter, r *http.Request) {
	data := ProfileViewData{RegisterError: "", ResetPasswordError: ""}
	if r.Method == http.MethodPost {
		instance := server.GetInstance()
		user, err := instance.GetActiveUser(r)
		errParsing := r.ParseForm()

		if err != nil || errParsing != nil || user == nil {
			data.RegisterError = "Internal error"
			data.ResetPasswordError = "Internal error"
		} else {
			fromForm := r.FormValue("from")
			if fromForm == "resetPassword" {
				resetPassword(r, &data, user)
			} else if fromForm == "addUser" {
				addUser(w, r, &data, user)
			}
		}
	}
	showPage(&data, w)
}

func resetPassword(r *http.Request, data *ProfileViewData, user *models.User) {
	if r.Method == http.MethodPost {
		{
			oldPassword := r.FormValue("oldPassword")
			newPassword := r.FormValue("newPassword")
			checkNewPassword := r.FormValue("againNewPassword")

			if newPassword == checkNewPassword {
				if utils.ComparePassword(oldPassword, user) {
					password := utils.HashPassword(newPassword)
					user.Password = password
					repository := repositories.GetUserRepository()
					err := repository.Update(user)
					if err != nil {
						data.ResetPasswordError = "Impossible to save your password please contact an admin"
					} else {
						data.ResetPasswordError = "Your password have been change"
					}
				} else {
					data.ResetPasswordError = "Wrong old password"
				}
			} else {
				data.ResetPasswordError = "New password and again new password is not the same"
			}
		}
	}
}

func addUser(w http.ResponseWriter, r *http.Request, p *ProfileViewData, user *models.User) {

}

func showPage(data *ProfileViewData, w http.ResponseWriter) {
	view.LoadView(w, "user/ProfileTemplate", data)
}
