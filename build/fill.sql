-- Заполнение таблицы "Мероприятия" (Events)
INSERT INTO Events (title, category, description, start_datetime, end_datetime, price, address, organizer, contacts)
VALUES
    ('Концерт вечных хитов', 'Музыка', 'Концерт легендарных хитов разных эпох', '2023-06-10 19:00:00', '2023-06-10 22:00:00', 50.00, 'Концертный зал "Академия"', 'Музыкальное агентство "Звезда"', 'info@agency.com'),
    ('Выставка современного искусства', 'Искусство', 'Выставка современных произведений искусства', '2023-06-15 10:00:00', '2023-06-15 18:00:00', 20.00, 'Галерея "АртПлаза"', 'Ассоциация современного искусства', 'info@gallery.com'),
    ('Фестиваль гастрономии', 'Еда', 'Фестиваль с участием лучших шеф-поваров и ресторанов города', '2023-06-20 16:00:00', '2023-06-20 22:00:00', 30.00, 'Центральный парк', 'Гастрономическое сообщество', 'info@foodfest.com');

-- Заполнение таблицы "Пользователи" (Users)
INSERT INTO Users (name, email, password, age, address, viewing_history)
VALUES
    ('Иван Петров', 'ivan@example.com', 'password123', 25, 'ул. Цветная, 10', NULL),
    ('Анна Сидорова', 'anna@example.com', 'qwerty456', 30, 'ул. Солнечная, 5', NULL),
    ('Михаил Иванов', 'mikhail@example.com', 'secure789', 35, 'пр. Победы, 20', NULL);

-- Заполнение таблицы "Билеты" (Tickets)
INSERT INTO Tickets (price, purchase_date, user_id, event_id)
VALUES
    (50.00, '2023-06-09', 1, 1),
    (20.00, '2023-06-14', 2, 2),
    (30.00, '2023-06-19', 3, 3);

-- Заполнение таблицы "Избранное" (Favorites)
INSERT INTO Favorites (event_id, user_id)
VALUES
    (1, 2),
    (2, 3),
    (3, 1);

-- Заполнение таблицы "Комментарии" (Comments)
INSERT INTO Comments (comment_text, comment_date, event_id, user_id)
VALUES
    ('Отличное выступление, понравилось!', '2023-06-11 10:30:00', 1, 1),
    ('Художественные работы на выставке впечатляют.', '2023-06-16 15:45:00', 2, 2),
    ('Фестиваль гастрономии был великолепным!', '2023-06-21 20:15:00', 3, 3);
