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
	http.HandleFunc(route, handler)
}

func (s Server) AddAuthRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.Handle(route, s.auth(handler))
}

func (s *Server) auth(handler func(response http.ResponseWriter, request *http.Request)) http.Handler {
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

/*func (s Server) BasicAuth(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the username and password from the request
		// Authorization header. If no Authentication header is present
		// or the header value is invalid, then the 'ok' return value
		// will be false.
		username, password, ok := r.BasicAuth()
		if ok {
			// Calculate SHA-256 hashes for the provided and expected
			// usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte("swansky"))
			expectedPasswordHash := sha256.Sum256([]byte("test"))

			// Use the subtle.ConstantTimeCompare() function to check if
			// the provided username and password hashes equal the
			// expected username and password hashes. ConstantTimeCompare
			// will return 1 if the values are equal, or 0 otherwise.
			// Importantly, we should to do the work to evaluate both the
			// username and password before checking the return values to
			// avoid leaking information.
			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			// If the username and password are correct, then call
			// the next handler in the chain. Make sure to return
			// afterwards, so that none of the code below is run.
			if usernameMatch && passwordMatch {
				next(w, r)
				return
			}
		}

		// If the Authentication header is not present, is invalid, or the
		// username or password is wrong, then set a WWW-Authenticate
		// header to inform the client that we expect them to use basic
		// authentication and send a 401 Unauthorized response.
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
*/
