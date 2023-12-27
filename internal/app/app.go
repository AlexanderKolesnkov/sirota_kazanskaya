package app

import (
	"github.com/AlexanderKolesnkov/sirota_kazanskaya/internal/transport/rest"
	"log"
	"os"
)

type Service struct {
	Rest *rest.Rest
	Port string
}

func New() *Service {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	return &Service{
		Rest: rest.New(),
		Port: port,
	}
}

func (s *Service) Init() {
	s.Rest.Routes()

	log.Printf("Listening on port %s", s.Port)
	s.Rest.Router.Run(":" + s.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
