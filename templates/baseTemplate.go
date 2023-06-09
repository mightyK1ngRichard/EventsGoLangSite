package templates

func GetBaseHTML() string {
	return `
	<body>
		<header class="p-3 text-bg-dark">
			<div class="container">
				<div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
		
					<ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
						<li><a href="#" class="nav-link px-2 text-secondary">Home</a></li>
						<li><a href="/events" class="nav-link px-2 text-white">Мероприятия</a></li>
						<li><a href="/tickets" class="nav-link px-2 text-white">Билеты</a></li>
						<li><a href="/create-event" class="nav-link px-2 text-white">Создать</a></li>
					</ul>
		
					<form class="col-12 col-lg-auto mb-3 mb-lg-0 me-lg-3" role="search" method="POST" action="/events">
						<input name="searching_text" type="search" class="form-control form-control-dark text-bg-dark" placeholder="Поиск заголовка" aria-label="Search">
					</form>
						<div class="text-end">
						<button type="button" class="btn btn-outline-light me-2">Войти</button>
						<button type="button" class="btn btn-warning">Регистрация</button>
					</div>
				</div>
			</div>
		</header>
	</body>
	`
}
