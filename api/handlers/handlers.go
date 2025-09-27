package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/mugen64/turtlor/configs"
	"github.com/mugen64/turtlor/internal/app/ui/home"
	"github.com/mugen64/turtlor/pkg/logger"
)

type HttpHandler struct {
	config      *configs.Config
	log         *logger.Logger
	HomeService *home.HomeService
}

func NewHttpHandler(config *configs.Config, logger *logger.Logger, r chi.Router) *HttpHandler {
	h := HttpHandler{
		config:      config,
		log:         logger,
		HomeService: home.NewHomeService(config, logger),
	}
	h.init(config, logger, r)
	return &h
}

func (h *HttpHandler) init(config *configs.Config, logger *logger.Logger, r chi.Router) {
	home.InitView(config, logger, r, h.HomeService)
}
