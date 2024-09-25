package apiserver

import (
	"net/http"

	"github.com/rs/cors"
)

func Start(config *Config) error {
	s := newServer()
	// используется rs/cors, только так cors разрешает post запрос с фронтенда
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{config.FrontendUrl}, //адресса, имеющие доступ к серверу
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true, //для cookie (вроде)
	}).Handler(s.router)

	return http.ListenAndServe(config.BindAddr, corsHandler)
}
