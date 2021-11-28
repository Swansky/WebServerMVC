package server

import (
	"awesomeProject1/models"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port    int
	userLog map[string]*models.User
}

var instance *Server

func NewServer(port int) *Server {
	server := new(Server)
	server.port = port
	server.userLog = make(map[string]*models.User)
	instance = server
	return server
}

func GetInstance() *Server {
	if instance == nil {
		panic("Server is nil !")
		return nil
	}
	return instance
}

func (s *Server) AddUserLog(uuid string, user *models.User) {
	s.userLog[uuid] = user
}

func (s *Server) RemoveUserLog(uuid string) {
	delete(s.userLog, uuid)
}

func (s Server) AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) {

	http.HandleFunc(route, s.addRouteChecker(route, handler))
}

func (s Server) AddAuthRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	s.AddRoute(route, s.auth(handler))
}

func (s Server) addRouteChecker(route string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != route {
			http.RedirectHandler("/notFound", 302).ServeHTTP(writer, request)
			return
		}
		handler(writer, request)
	}
}

func (s *Server) auth(handler func(response http.ResponseWriter, request *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("token")
		if cookie != nil {
			if len(cookie.Value) > 0 {
				if s.userIsLog(cookie.Value) {
					handler(w, r)
					return
				}
			}
		}
		http.RedirectHandler("/login", 302).ServeHTTP(w, r)
	})
}

func (s *Server) userIsLog(token string) bool {
	for index, _ := range s.userLog {
		if index == token {
			return true
		}
	}
	return false
}

func (s *Server) UserIsAuthenticate(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}
	if cookie != nil {
		if s.userIsLog(cookie.Value) {
			return true
		}
	}
	return false
}

func (s *Server) getUserLog() map[string]*models.User {
	return s.userLog
}

func (s Server) Start() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/public", fileServer)
	stringPort := fmt.Sprintf(":%d", s.port)
	fmt.Printf(fmt.Sprintf("Starting server at port %s \n", stringPort))
	if err := http.ListenAndServe(stringPort, nil); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) LogoutUser(request *http.Request) {
	cookie, err := request.Cookie("token")
	if err != nil {
		panic(err)
		return
	}
	if cookie != nil {
		s.RemoveUserLog(cookie.Value)
	}
}
