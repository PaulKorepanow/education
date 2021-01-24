package server

import (
	store "bookLibrary/internal/store"
	"bookLibrary/internal/store/sqlstore"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func NewServer(db store.Store) *Server {
	s := &Server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  db,
	}

	s.configureRouter()
	return s
}

func (s *Server) Start(config *Config) error {
	if err := s.configureLogger(config); err != nil {
		return err
	}
	s.logger.Info("Starting api server")

	return http.ListenAndServe(config.BindAddr, s.router)
}

func ConnectToDB(databaseURL string) (store.Store, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      false,
			LogLevel:      logger.Silent,
		},
	)

	newLogger.LogMode(logger.Silent)

	db, err := gorm.Open(
		postgres.Open(databaseURL),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		return nil, err
	}
	sqlDB := sqlstore.NewStore(db)
	return sqlDB, nil
}

func (s Server) ServeHttp(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureLogger(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/api/user/new", s.handleUsersCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/user/login", s.Authenticate()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/user/{id:[0-9]+}/password", s.UpdatePassword()).Methods(http.MethodPut)
	s.router.HandleFunc("/api/user/{id:[0-9]+}/book", s.AddBook()).Methods(http.MethodPost)

}
