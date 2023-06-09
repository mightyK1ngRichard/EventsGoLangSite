package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/store"
	"github.com/mightyK1ngRichard/EventsGoLangSite/templates"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path/filepath"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start Запуск сервера.
func (a *Server) Start() error {
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

// Наш логгер.
func (a *Server) configLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)
	return nil
}

// configRouter. Наши маршруты.
func (a *Server) configRouter() {
	a.router.HandleFunc(templates.EventsURL, a.events())
	a.router.HandleFunc(templates.TicketsURL, a.tickets())
	a.router.HandleFunc(templates.NewEventURL, a.newEvent())
	a.router.HandleFunc(templates.EventURL, a.event())
	a.router.HandleFunc(templates.SignUpURL, a.signUp())
	a.router.HandleFunc(templates.SignInURL, a.signIn())
	a.router.HandleFunc(templates.HomeURL, a.home())
	a.router.NotFoundHandler = a.notFoundPage()
}

// Подключаем хранилище для работы с БД.
func (a *Server) configStore() error {
	st := store.NewStore(a.config.Store, a.logger)
	if err := st.Open(); err != nil {
		return err
	}
	a.store = st
	return nil
}

// Домашнаяя страница.
func (a *Server) home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// Все мероприятия.
func (a *Server) events() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			list, err := a.store.Event().List()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			} else {
				tmpl, err := getHTMLPage("events")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				if err := tmpl.Execute(w, map[string]interface{}{
					"base_html": template.HTML(templates.GetBaseHTML()),
					"list":      list,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}

		case "POST":
			titleFromForm := r.FormValue("searching_text")
			events, err := a.store.Event().EventByTitle(titleFromForm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
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
			tmpl.Execute(w, map[string]interface{}{
				"base_html": template.HTML(templates.GetBaseHTML()),
				"list":      events,
			})

		default:
			http.NotFound(w, r)
		}

	}
}

// Создать мероприятие.
func (a *Server) newEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tmpl, err := getHTMLPage("create_event")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, map[string]interface{}{
				"base_html": template.HTML(templates.GetBaseHTML()),
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		case "POST":
			// Ридерект на старницу всех событий
			defer http.Redirect(w, r, "/events", http.StatusFound)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
				return
			}
			title := r.Form.Get("title")
			category := r.Form.Get("category")
			description := r.Form.Get("description")
			dateStart := r.Form.Get("date_start")
			dateEnd := r.Form.Get("date_end")
			price := r.Form.Get("price")
			address := r.Form.Get("address")
			organization := r.Form.Get("organization")
			contacts := r.Form.Get("contacts")
			if err := a.store.Event().CreateEvent(title, category, description, dateStart, dateEnd, price, address,
				organization, contacts); err != nil {
				a.logger.Error(err)
				return
			}

		default:
			http.NotFound(w, r)
		}
	}
}

// Билеты.
func (a *Server) tickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickets, err := a.store.Ticket().Tickets()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		} else {
			tmpl, err := getHTMLPage("tickets")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, map[string]interface{}{
				"base_html": template.HTML(templates.GetBaseHTML()),
				"tickets":   tickets,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// Открыть мероприятие.
func (a *Server) event() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		event, comments, err := a.store.Event().EventByID(id)
		if event == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		} else if comments == nil {
			comments = make([]*model.Comment, 0)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		} else {
			tmpl, err := getHTMLPage("event")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, map[string]interface{}{
				"event":     event,
				"comments":  comments,
				"base_html": template.HTML(templates.GetBaseHTML()),
				//"cssPath":  cssPath,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// Авторизация.
func (a *Server) signUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tmpl, err := getHTMLPage("register")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, map[string]interface{}{
				"base_html": template.HTML(templates.GetBaseHTML()),
			})

		case "POST":
			defer http.Redirect(w, r, "/events", http.StatusFound)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
				return
			}
			email := r.Form.Get("email")
			password := r.Form.Get("password")

			u := &model.User{
				Email:    email,
				Password: password,
			}
			user, err := a.store.User().Create(u)
			if err != nil {
				a.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			// Скроем пароль.
			user.Password = ""
			//a.respond(w, r, http.StatusOK, user)
		}
	}
}

// Вход.
func (a *Server) signIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tmpl, err := getHTMLPage("sign_in")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		case "POST":
			defer http.Redirect(w, r, "/events", http.StatusFound)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
				return
			}
			email := r.Form.Get("email")
			password := r.Form.Get("password")

			u := &model.User{
				Email:    email,
				Password: password,
			}

			user, err := a.store.User().Create(u)
			if err != nil {
				a.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			// Скроем пароль.
			user.Password = ""
			//a.respond(w, r, http.StatusOK, user)
		}
	}
}

// Вывод ошибки в json.
func (a *Server) error(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	a.respond(w, r, statusCode, map[string]string{"error": err.Error()})
}

// Вывод данных в json.
func (a *Server) respond(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}

// Страница не найдена.
func (a *Server) notFoundPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		tmpl, err := getHTMLPage("not_found_page")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Возваращает html страницу для рендера.
func getHTMLPage(nameHTMLFile string) (*template.Template, error) {
	rootDir, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	templatePath := filepath.Join(rootDir, "templates", nameHTMLFile+".html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
