-- Создание таблицы "Мероприятия"
CREATE TABLE IF NOT EXISTS Events
(
    id             BIGSERIAL PRIMARY KEY,
    title          VARCHAR(100) NOT NULL,
    category       VARCHAR(50)  NOT NULL,
    description    VARCHAR(255) NOT NULL,
    start_datetime TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_datetime   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price          REAL         NOT NULL,
    address        VARCHAR(100) NOT NULL,
    organizer      VARCHAR(100) NOT NULL,
    contacts       VARCHAR(100) NOT NULL
);

-- Создание таблицы "Пользователи"
CREATE TABLE IF NOT EXISTS Users
(
    id              BIGSERIAL PRIMARY KEY,
    name            VARCHAR(100) NOT NULL DEFAULT 'user',
    email           VARCHAR(100) NOT NULL,
    password        VARCHAR(400)  NOT NULL,
    age             INT          NOT NULL DEFAULT 0,
    address         VARCHAR(255) NOT NULL DEFAULT '',
    viewing_history VARCHAR(255) NOT NULL DEFAULT ''
);

-- Создание таблицы "Билеты"
CREATE TABLE IF NOT EXISTS Tickets
(
    id            BIGSERIAL PRIMARY KEY,
    price         DECIMAL(10, 2) NOT NULL,
    purchase_date DATE           NOT NULL DEFAULT CURRENT_DATE,
    user_id       INT            NOT NULL,
    event_id      INT            NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (event_id) REFERENCES Events (id)
);

-- Создание таблицы "Избранное"
CREATE TABLE IF NOT EXISTS Favorites
(
    id       BIGSERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    user_id  INT NOT NULL,
    FOREIGN KEY (event_id) REFERENCES Events (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

-- Создание таблицы "Комментарии"
CREATE TABLE IF NOT EXISTS Comments
(
    id           BIGSERIAL PRIMARY KEY,
    comment_text VARCHAR(255) NOT NULL,
    comment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    event_id     INT          NOT NULL,
    user_id      INT          NOT NULL,
    FOREIGN KEY (event_id) REFERENCES Events (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);
