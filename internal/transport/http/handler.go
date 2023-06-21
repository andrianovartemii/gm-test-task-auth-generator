package http

import (
	"gm-test-task-auth-generator/internal/services"
	"net/http"
	"strings"
)

type Handler struct {
	authService *services.AuthService
}

func NewHttpHandler(authService *services.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (handler *Handler) Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/generate", handler.GenerateToken)
	mux.HandleFunc("/validate", handler.ValidateToken)
	return mux
}

func (handler *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("pong"))
}

func (handler *Handler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("invalid method, use GET"))
		return
	}
	login := r.URL.Query().Get("login")
	token, err := handler.authService.GenerateToken(login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(token))
}

func (handler *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("invalid method, use GET"))
		return
	}
	splitAuthorizationHeader := strings.Split(r.Header.Get("Authorization"), " ")
	if len(splitAuthorizationHeader) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid Bearer token format"))
		return
	}
	token := strings.TrimSpace(splitAuthorizationHeader[1])
	err := handler.authService.ValidateToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
}
