package templates

import (
	"fmt"
)

const (
	HomeURL     = "/home"
	EventsURL   = "/events"
	NewEventURL = "/create-event"
	SignInURL   = "/signIn"
	SignUpURL   = "/signUp"
	TicketsURL  = "/tickets"
	EventURL    = "/event/{id}"
)

func GetBaseHTML() string {
	return fmt.Sprintf(
		`
	<body>
		<header class="p-3 text-bg-dark">
			<div class="container">
				<div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
		
					<ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
						<li><a href="#" class="nav-link px-2 text-secondary">Home</a></li>
						<li><a href="%s" class="nav-link px-2 text-white">Мероприятия</a></li>
						<li><a href="%s" class="nav-link px-2 text-white">Билеты</a></li>
						<li><a href="%s" class="nav-link px-2 text-white">Создать</a></li>
					</ul>
		
					<form class="col-12 col-lg-auto mb-3 mb-lg-0 me-lg-3" role="search" method="POST" action="%s">
						<input name="searching_text" type="search" class="form-control form-control-dark text-bg-dark" placeholder="Поиск заголовка" aria-label="Search">
					</form>
						<div class="text-end">
						<a type="button" href="%s" class="btn btn-outline-light me-2">Войти</a>
						<a type="button" href="%s" class="btn btn-warning">Регистрация</a>
					</div>
				</div>
			</div>
		</header>
	</body>
	`,
		EventsURL,
		TicketsURL,
		NewEventURL,
		EventsURL,
		SignInURL,
		SignUpURL,
	)
}
