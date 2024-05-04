package settings

import (
	"flag"
	"log"
	"os"

	"github.com/ftsog/ecom/config"
	"github.com/ftsog/ecom/database"
	"github.com/ftsog/ecom/handlers"
	"github.com/ftsog/ecom/logger"
	"github.com/ftsog/ecom/mailer"
	"github.com/ftsog/ecom/models"
	"github.com/ftsog/ecom/routers"
	"github.com/ftsog/ecom/session"
	"github.com/go-chi/chi/v5"
)

var (
	environment string
)

type appConfig struct {
	cfg *config.Config
}

func init() {
	env := flag.String("env", "", "ecom -env container")
	flag.Parse()

	if len(os.Args[:]) <= 1 {
		flag.Usage()
		os.Exit(1)
	} else {
		environment = *env
	}
}

func (c *appConfig) NewModel() *models.Model {
	dbConn, err := database.DatabaseConnection(c.cfg.Postgresql.Host, c.cfg.Postgresql.Port, c.cfg.Postgresql.User,
		c.cfg.Postgresql.DbName, c.cfg.Postgresql.Password, c.cfg.Postgresql.SSLmode)
	if err != nil {
		log.Panic(err)
	}

	m := &models.Model{
		DB: dbConn,
	}

	return m
}

func (c *appConfig) NewSession() *session.Session {
	rdStore, err := database.RediStore(10, "tcp", c.cfg.Redis.Host, c.cfg.Redis.Port, c.cfg.Redis.Password, []byte(c.cfg.Redis.SecretKey))
	if err != nil {
		log.Panic(err)
	}

	s := &session.Session{
		RediStore: rdStore,
	}

	return s
}

func NewHandler() *handlers.Handler {
	cfg := appConfig{
		cfg: config.LoadConfig(environment),
	}

	//set mail config
	mailer.MailConfig = cfg.cfg

	m := cfg.NewModel()
	l := logger.NewLogger()
	s := cfg.NewSession()

	h := &handlers.Handler{
		Db:        m,
		Logger:    l,
		RDsession: s,
	}

	return h
}

func NewRouter(r *chi.Mux) *routers.Router {
	handler := NewHandler()
	router := &routers.Router{
		Route:   r,
		Handler: handler,
	}

	return router
}
