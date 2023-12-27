package app

import (
	"github.com/AlexanderKolesnkov/sirota_kazanskaya/internal/config"
	"github.com/AlexanderKolesnkov/sirota_kazanskaya/internal/transport/rest"
	"log"
)

type Service struct {
	Cfg  *config.Config
	Rest *rest.Rest
}

func New() *Service {
	return &Service{
		Cfg:  config.New(),
		Rest: rest.New(),
	}
}

func (s *Service) Init() {
	s.Rest.Routes()

	log.Printf("Listening on port %s", s.Cfg.Port)
	s.Rest.Router.Run(":" + s.Cfg.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
