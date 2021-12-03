package controller

import (
	"awesomeProject1/models"
	"awesomeProject1/repositories"
	"awesomeProject1/server"
	"awesomeProject1/utils"
	"awesomeProject1/view"
	"net/http"
	"strings"
)

type ProfileViewData struct {
	RegisterError      string
	ResetPasswordError string
	DeleteUserError    string
	Users              []*models.User
}

func Profile(w http.ResponseWriter, r *http.Request) {
	data := ProfileViewData{RegisterError: "", ResetPasswordError: "", DeleteUserError: ""}
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
				addUser(r, &data)
			} else if fromForm == "deleteUser" {
				deleteUser(r, &data)
			}
		}
	}
	read, err := repositories.GetUserRepository().Read()
	if err != nil {
		data.DeleteUserError = "Internal error"
		panic(err)
	}
	for e := read.Front(); e != nil; e = e.Next() {
		user := e.Value.(*models.User)
		data.Users = append(data.Users, user)
	}
	showPage(&data, w)
}

func deleteUser(r *http.Request, p *ProfileViewData) {
	value := r.FormValue("uuid")
	repository := repositories.GetUserRepository()
	user, err := repository.FindById(value)
	if err != nil || user == nil {
		p.DeleteUserError = "Impossible to delete this user "
		return
	}
	errorDelete := repository.Delete(user)
	if errorDelete != nil {
		p.DeleteUserError = "Impossible to delete this user "
	} else {
		p.DeleteUserError = "User have been deleted"
	}
}

func resetPassword(r *http.Request, data *ProfileViewData, user *models.User) {
	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")
	checkNewPassword := r.FormValue("againNewPassword")
	if len(oldPassword) > 0 && len(newPassword) > 0 && len(checkNewPassword) > 0 {
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
	} else {
		data.ResetPasswordError = "Empty input is forbidden"
	}

}

func addUser(r *http.Request, p *ProfileViewData) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) > 0 && len(password) > 0 {
		newUser := models.NewUser(username, password)
		repository := repositories.GetUserRepository()
		_, err := repository.Create(newUser)

		if err != nil {
			if strings.Contains(err.Error(), " Duplicate entry") {
				p.RegisterError = "Username already exist"
			} else {
				p.RegisterError = "Internal error "
				panic(err)
			}

		} else {
			p.RegisterError = "User have been created"
		}
	} else {
		p.RegisterError = "username or password is empty !"
	}
}

func showPage(data *ProfileViewData, w http.ResponseWriter) {
	view.LoadView(w, "user/ProfileTemplate", data)
}
