package controller

import (
	"awesomeProject1/models"
	"awesomeProject1/repositories"
	server2 "awesomeProject1/server"
	"awesomeProject1/view"
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/nu7hatch/gouuid"
	"net/http"
	"time"
)

type LoginViewData struct {
	Error string
}

func Login(w http.ResponseWriter, r *http.Request) {
	server := server2.GetInstance()
	if server.UserIsAuthenticate(r) {
		http.RedirectHandler("/", 302).ServeHTTP(w, r)
		return
	}

	var messageError = ""
	err := r.ParseForm()
	if err != nil {
		panic(err)
		return
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if len(username) > 0 && len(password) > 0 {

			read, err := repositories.GetUserRepository().Read()
			if err != nil {
				panic(err)
				return
			}

			for e := read.Front(); e != nil; e = e.Next() {
				value := e.Value
				user := value.(*models.User)
				passwordValue := checkPassword(password, user)
				if user.Username == username && passwordValue {
					v4, err := uuid.NewV4()
					if err != nil {
						panic(err)
						return
					}
					uuidUser := v4.String()

					server.AddUserLog(uuidUser, user)

					cookie := &http.Cookie{
						Name:   "token",
						Value:  uuidUser,
						MaxAge: 30000000,
					}
					http.SetCookie(w, cookie)
					http.RedirectHandler("/", 302).ServeHTTP(w, r)
					return
				} else {
					messageError = "Wrong Login or Password."
				}
			}
		} else {
			messageError = "Username or password is empty."
		}
	}

	data := LoginViewData{Error: messageError}
	view.LoadView(w, "LoginTemplate", data)
}

func Logout(w http.ResponseWriter, request *http.Request) {
	server := server2.GetInstance()
	server.LogoutUser(request)
	clearCookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}
	http.SetCookie(w, clearCookie)
	http.RedirectHandler("/login", 302).ServeHTTP(w, request)
}

func checkPassword(password string, user *models.User) bool {
	passwordHash := md5.Sum([]byte(password))
	passwordMatch := hex.EncodeToString(passwordHash[:]) == user.Password
	if passwordMatch {
		return true
	}
	return false
}
