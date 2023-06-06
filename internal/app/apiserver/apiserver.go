package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/store"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	homeURL   = "/home"
	eventsURL = "/events"
	eventURL  = "/event/{id}"
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
	a.router.HandleFunc(eventsURL, a.events())
	a.router.HandleFunc(eventURL, a.event())
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

func (a *APIServer) event() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		event, comments, err := a.store.Event().EventByID(id)
		if event == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		} else if comments == nil {
			fmt.Fprintf(w, "Странно это всё, но обрабатывать лень")
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		} else {
			rootDir, err := filepath.Abs(".")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			templatePath := filepath.Join(rootDir, "templates", "event.html")
			//cssPath := filepath.Join(rootDir, "static", "styles.css")
			tmpl, err := template.ParseFiles(templatePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tmpl.Execute(w, map[string]interface{}{
				"event":    event,
				"comments": comments,
				//"cssPath":  cssPath,
			})
		}
	}
}
