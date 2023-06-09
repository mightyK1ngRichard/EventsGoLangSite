package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/model"
	"github.com/mightyK1ngRichard/EventsGoLangSite/internal/app/store"
	"github.com/mightyK1ngRichard/EventsGoLangSite/templates"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	sessionName        = "mightyK1ngRichard"
	ctxKeyUser  ctxKey = iota
)

// Ключи контекста.
type ctxKey int8

type Server struct {
	config       *Config
	logger       *logrus.Logger
	router       *mux.Router
	store        *store.Store
	sessionStore sessions.Store
}

func New(config *Config, session sessions.Store) *Server {
	return &Server{
		config:       config,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
		sessionStore: session,
	}
}

// Start Запуск сервера.
func (a *Server) Start() error {
	defer a.store.Close()
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
	a.router.HandleFunc(templates.EventsURL, a.authenticateUser(a.events()).ServeHTTP)
	a.router.HandleFunc(templates.TicketsURL, a.authenticateUser(a.tickets()).ServeHTTP)
	a.router.HandleFunc(templates.NewEventURL, a.authenticateUser(a.newEvent()).ServeHTTP)
	a.router.HandleFunc(templates.EventURL, a.authenticateUser(a.event()).ServeHTTP)
	a.router.HandleFunc(templates.SignUpURL, a.signUp())
	a.router.HandleFunc(templates.SignInURL, a.signIn())
	a.router.HandleFunc(templates.LogoutURL, a.logout())
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

// Аутинфиакация юзера.
func (a *Server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := a.sessionStore.Get(r, sessionName)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			next.ServeHTTP(w, r)
			return
		}
		user, err2 := a.store.User().FindById(id.(string))
		if err2 != nil {
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, user)))
	})
}

// Проверка авторизации.
//func (a *Server) isAuth(r *http.Request) *model.User {
//	session, err := a.sessionStore.Get(r, sessionName)
//	if err == nil {
//		userId, ok := session.Values["user_id"]
//		if ok {
//			res, ok := userId.(string)
//			if ok {
//				user, err := a.store.User().FindById(res)
//				if err == nil {
//					return user
//				}
//			}
//		}
//	}
//	return nil
//}

// Удаление сессии.
func (a *Server) deleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := a.sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	if err := a.sessionStore.Save(r, w, session); err != nil {
		return err
	}
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

				// Смотрим контекст на авторизацию.
				user, ok := r.Context().Value(ctxKeyUser).(*model.User)
				if ok {
					if err := tmpl.Execute(w, map[string]interface{}{
						"list":   list,
						"isAuth": true,
						"user":   user,
					}); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}

				} else {
					if err := tmpl.Execute(w, map[string]interface{}{
						"base_html": template.HTML(templates.GetBaseHTML()),
						"list":      list,
						"isAuth":    false,
						"user":      nil,
					}); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
				}
			}

		case "POST":
			titleFromForm := r.FormValue("searching_text")
			events, err := a.store.Event().EventByTitle(titleFromForm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl, err := getHTMLPage("events")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Смотрим контекст на авторизацию.
			user, ok := r.Context().Value(ctxKeyUser).(*model.User)
			if ok {
				if err := tmpl.Execute(w, map[string]interface{}{
					"list":   events,
					"isAuth": true,
					"user":   user,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			} else {
				if err := tmpl.Execute(w, map[string]interface{}{
					"base_html": template.HTML(templates.GetBaseHTML()),
					"list":      events,
					"isAuth":    false,
					"user":      nil,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}

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

			// Смотрим контекст на авторизацию.
			user, ok := r.Context().Value(ctxKeyUser).(*model.User)
			if ok {
				if err := tmpl.Execute(w, map[string]interface{}{
					"isAuth": true,
					"user":   user,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			} else {
				if err := tmpl.Execute(w, map[string]interface{}{
					"base_html": template.HTML(templates.GetBaseHTML()),
					"isAuth":    false,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
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

			// Смотрим контекст на авторизацию.
			user, ok := r.Context().Value(ctxKeyUser).(*model.User)
			if ok {
				if err := tmpl.Execute(w, map[string]interface{}{
					"isAuth":  true,
					"user":    user,
					"tickets": tickets,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			} else {
				if err := tmpl.Execute(w, map[string]interface{}{
					"base_html": template.HTML(templates.GetBaseHTML()),
					"isAuth":    false,
					"tickets":   tickets,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
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
			// Смотрим контекст на авторизацию.
			user, ok := r.Context().Value(ctxKeyUser).(*model.User)
			if ok {
				if err := tmpl.Execute(w, map[string]interface{}{
					"isAuth":   true,
					"user":     user,
					"event":    event,
					"comments": comments,
				}); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			} else {
				if err := tmpl.Execute(w, map[string]interface{}{
					"isAuth":    false,
					"user":      user,
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
			defer http.Redirect(w, r, templates.EventsURL, http.StatusFound)
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

			user, err := a.store.User().CheckUser(u)
			if err != nil {
				http.Redirect(w, r, templates.SignInURL, http.StatusFound)
				return
			}
			// Скроем пароль.
			user.Password = ""

			session, err := a.sessionStore.Get(r, sessionName)
			if err != nil {
				a.errorPage(w)
				return
			}
			session.Values["user_id"] = user.ID
			if err := a.sessionStore.Save(r, w, session); err != nil {
				a.errorPage(w)
				return
			}

			//a.respond(w, r, http.StatusOK, user)
		}
	}
}

// Разлогирование
func (a *Server) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if err := a.deleteSession(w, r); err != nil {
				a.errorPage(w)
				return
			}
			http.Redirect(w, r, templates.EventsURL, http.StatusFound)
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

func (a *Server) errorPage(w http.ResponseWriter) {
	htmlString := `
	<html><body>
	<h1>Что-то пошло не так:(</h1>
	<p>Перейдите на страницу <a href="%s">Events</a></p>
	</body></html>
	`
	fmt.Fprintf(w, htmlString, templates.EventsURL)
}
