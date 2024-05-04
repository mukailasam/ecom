package handlers

import (
	"github.com/ftsog/ecom/logger"
	"github.com/ftsog/ecom/models"
	"github.com/ftsog/ecom/session"
)

type Handler struct {
	Db        *models.Model
	Logger    *logger.Logger
	RDsession *session.Session
}
