package home

import (
	"github.com/mugen64/turtlor/configs"
	"github.com/mugen64/turtlor/pkg/logger"
)

type HomeService struct {
	cfg *configs.Config
	log *logger.Logger
}

func NewHomeService(
	cfg *configs.Config,
	logger *logger.Logger,
) *HomeService {
	return &HomeService{
		cfg: cfg,
		log: logger,
	}
}
