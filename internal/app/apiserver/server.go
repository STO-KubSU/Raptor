package apiserver

//http -v --session=user http://localhost:8080/private/whoami

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Константные ключи/поля
const (
	sessionName        = "usersession"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

// Особые типы ошибок
var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("no authenticated")
)

// Тип для ключей контекста
type ctxKey int8

// Файл, определяющий сервер
type server struct {
	router *mux.Router
}

// Конструктор сервера (Хранилище(БД) -> сервер)
func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

// Ф-ция делегирования запросов в Роутер (используется для учучшения условий тестирования)
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Конфигурация роутера (запросов)
func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	//две строки ниже что-то делают ? вроде нет, а должны
	//s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	//s.router.Use(handlers.CORS(handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})))
}

// MIDW Присвоение запросам ID
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

// Служебный метод ошибки
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// Служебный метод ответа
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
