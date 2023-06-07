-- Создание таблицы "Мероприятия"
CREATE TABLE Events
(
    id             BIGSERIAL PRIMARY KEY,
    title          VARCHAR(100),
    category       VARCHAR(50),
    description    VARCHAR(255),
    start_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_datetime   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    price          REAL,
    address        VARCHAR(100),
    organizer      VARCHAR(100),
    contacts       VARCHAR(100)
);

-- Создание таблицы "Пользователи"
CREATE TABLE Users
(
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(100),
    email           VARCHAR(100),
    password        VARCHAR(50),
    age             INT,
    address         VARCHAR(255),
    viewing_history VARCHAR(255)
);

-- Создание таблицы "Билеты"
CREATE TABLE Tickets
(
    id            BIGSERIAL PRIMARY KEY,
    price         DECIMAL(10, 2),
    purchase_date DATE,
    user_id       INT,
    event_id      INT,
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (event_id) REFERENCES Events (id)
);

-- Создание таблицы "Избранное"
CREATE TABLE Favorites
(
    id       BIGSERIAL PRIMARY KEY,
    event_id INT,
    user_id  INT,
    FOREIGN KEY (event_id) REFERENCES Events (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

-- Создание таблицы "Комментарии"
CREATE TABLE Comments
(
    id           BIGSERIAL PRIMARY KEY,
    comment_text VARCHAR(255),
    comment_date TIMESTAMP,
    event_id     INT,
    user_id      INT,
    FOREIGN KEY (event_id) REFERENCES Events (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);
