package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/store"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	homeURL = "/home"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) Start() error {
	if err := a.configLogger(); err != nil {
		return err
	}
	if err := a.configStore(); err != nil {
		return err
	}
	a.configRouter()
	a.logger.Info("start api server")
	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (a *APIServer) configLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)
	return nil
}

func (a *APIServer) configRouter() {
	// TODO: удалить test
	a.router.HandleFunc("/test", a.test())
	a.router.HandleFunc("/events", a.events())
	a.router.HandleFunc(homeURL, a.home())
}

func (a *APIServer) configStore() error {
	st := store.NewStore(a.config.Store, a.logger)
	if err := st.Open(); err != nil {
		return err
	}
	a.store = st
	return nil
}

func (a *APIServer) home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// TODO: удалить test
func (a *APIServer) test() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newTest, err := a.store.Test().Create(&model.Test{
			Name: "Dima",
			Info: "Boss",
		})
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			a.logger.Fatalln(err)

		} else {
			marshal, err2 := json.Marshal(newTest)
			if err2 != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				a.logger.Error(err)
			}
			a.logger.Infof("create new note: %v", newTest)
			fmt.Fprintf(w, "%s", marshal)
		}
	}
}

func (a *APIServer) events() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := a.store.Event().List()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		} else {
			rootDir, err := filepath.Abs(".")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			templatePath := filepath.Join(rootDir, "templates", "events.html")
			tmpl, err := template.ParseFiles(templatePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, list)
		}
	}
}
