package home

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/mugen64/turtlor/configs"
	"github.com/mugen64/turtlor/pkg/logger"
)

func InitView(cfg *configs.Config, logger *logger.Logger, r chi.Router, svc *HomeService) {
	r.Handle("/", templ.Handler(Index(cfg)))
}
