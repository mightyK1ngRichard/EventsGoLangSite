-- Заполнение таблицы "Мероприятия" (Events)
INSERT INTO Events (title, category, description, start_datetime, end_datetime, price, address, organizer, contacts)
VALUES
    ('Концерт вечных хитов', 'Музыка', 'Концерт легендарных хитов разных эпох', '2023-06-10 19:00:00', '2023-06-10 22:00:00', 50.00, 'Концертный зал "Академия"', 'Музыкальное агентство "Звезда"', 'info@agency.com'),
    ('Выставка современного искусства', 'Искусство', 'Выставка современных произведений искусства', '2023-06-15 10:00:00', '2023-06-15 18:00:00', 20.00, 'Галерея "АртПлаза"', 'Ассоциация современного искусства', 'info@gallery.com'),
    ('Фестиваль гастрономии', 'Еда', 'Фестиваль с участием лучших шеф-поваров и ресторанов города', '2023-06-20 16:00:00', '2023-06-20 22:00:00', 30.00, 'Центральный парк', 'Гастрономическое сообщество', 'info@foodfest.com');

INSERT INTO Events (title, category, description, start_datetime, end_datetime, price, address, organizer, contacts)
VALUES
    ('Концерт рок-группы', 'Музыка', 'Крупный концерт рок-группы с участием известных музыкантов', '2023-07-05 20:00:00', '2023-07-05 23:00:00', 60.00, 'Стадион "Рок-Арена"', 'Концертное агентство "Роковая Волна"', 'info@rockwave.com'),
    ('Театральное представление', 'Театр', 'Премьерный спектакль в театре', '2023-07-10 19:30:00', '2023-07-10 21:30:00', 40.00, 'Театр им. Пушкина', 'Театральная компания "АртПро"', 'info@artprotheatre.com'),
    ('Выставка фотографии', 'Искусство', 'Выставка фотографий природы и путешествий', '2023-07-15 11:00:00', '2023-07-15 17:00:00', 15.00, 'Галерея "ФотоАрт"', 'Фотоагентство "Пейзажи"', 'info@photoartagency.com'),
    ('Спортивный турнир', 'Спорт', 'Международный спортивный турнир по футболу', '2023-07-20 14:00:00', '2023-07-20 18:00:00', 25.00, 'Стадион "СпортАрена"', 'Спортивная федерация "Футбол+"', 'info@footballplus.com'),
    ('Конференция по IT-технологиям', 'Бизнес', 'Конференция с участием ведущих экспертов IT-индустрии', '2023-07-25 09:00:00', '2023-07-25 17:00:00', 50.00, 'Конгресс-центр "TechHub"', 'IT-компания "ТехноГлобус"', 'info@techglobus.com'),
    ('Фестиваль кино', 'Кино', 'Фестиваль с показом лучших фильмов и встречами с режиссерами', '2023-07-30 18:30:00', '2023-07-30 23:00:00', 35.00, 'Кинотеатр "СинемаПлаза"', 'Киностудия "Империя"', 'info@empirecinema.com'),
    ('Мастер-класс по живописи', 'Искусство', 'Мастер-класс с известным художником', '2023-08-05 15:00:00', '2023-08-05 17:00:00', 20.00, 'Художественная студия "Творчество"', 'Художественная ассоциация "Креатив"', 'info@creativeart.org'),
    ('Выставка автомобилей', 'Авто', 'Выставка новых моделей автомобилей', '2023-08-10 10:00:00', '2023-08-10 18:00:00', 10.00, 'Выставочный центр "АвтоШоу"', 'Автосалон "МоторЛайн"', 'info@motorline.com'),
    ('Фестиваль моды', 'Мода', 'Показ мод с участием известных дизайнеров', '2023-08-15 17:30:00', '2023-08-15 22:00:00', 30.00, 'Модный театр "Glamour"', 'Модельное агентство "FashionStyle"', 'info@fashionstyleagency.com'),
    ('Кулинарный мастер-класс', 'Еда', 'Мастер-класс по приготовлению итальянской кухни', '2023-08-20 16:30:00', '2023-08-20 19:00:00', 25.00, 'Кулинарная школа "Delizioso"', 'Шеф-повар "Итальянский вкус"', 'info@italiantaste.com'),
    ('Концерт симфонического оркестра', 'Музыка', 'Выступление классического симфонического оркестра', '2023-08-25 20:30:00', '2023-08-25 22:30:00', 45.00, 'Концертный зал "Симфония"', 'Культурная ассоциация "Классика"', 'info@classicalarts.org'),
    ('Тематическая вечеринка', 'Развлечения', 'Вечеринка в стиле 80-х', '2023-08-30 22:00:00', '2023-08-31 02:00:00', 15.00, 'Ночной клуб "Retro"', 'Организационный клуб "PartyTime"', 'info@partytimeclub.com'),
    ('Спектакль детского театра', 'Театр', 'Яркий спектакль для детей', '2023-09-05 11:00:00', '2023-09-05 12:30:00', 10.00, 'Детский театр "Сказочный мир"', 'Театральная студия "Детские грезы"', 'info@childrenstheatre.com'),
    ('Концерт джазовой музыки', 'Музыка', 'Концерт джазовой группы с выступлениями солистов', '2023-09-10 19:30:00', '2023-09-10 22:30:00', 35.00, 'Джазовый клуб "JazzMood"', 'Музыкальное агентство "JazzGroove"', 'info@jazzgroove.com'),
    ('Выставка ювелирных изделий', 'Искусство', 'Выставка уникальных ювелирных украшений', '2023-09-15 12:00:00', '2023-09-15 20:00:00', 25.00, 'Выставочный центр "JewelExpo"', 'Ювелирный дом "Precious Gems"', 'info@preciousgems.com'),
    ('Фестиваль научной фантастики', 'Развлечения', 'Фестиваль с презентациями фильмов и встречами с авторами', '2023-09-20 16:00:00', '2023-09-20 23:00:00', 20.00, 'Конференц-центр "Sci-Fi"', 'Ассоциация фантастики "CosmoWorld"', 'info@cosmoworld.com'),
    ('Спортивный марафон', 'Спорт', 'Международный марафон с участием профессиональных бегунов', '2023-09-25 08:00:00', '2023-09-25 14:00:00', 30.00, 'Городской парк', 'Спортивная федерация "Беговые трассы"', 'info@runtracks.com'),
    ('Выставка художественных ремесел', 'Искусство', 'Выставка уникальных изделий художественных ремесел', '2023-09-30 11:30:00', '2023-09-30 18:30:00', 15.00, 'Художественный центр "CraftArt"', 'Ассоциация ремесленников "Artisans"', 'info@artisans.org'),
    ('Концерт хоровой музыки', 'Музыка', 'Выступление профессионального хора с классической музыкой', '2023-10-05 19:00:00', '2023-10-05 21:00:00', 40.00, 'Церковь "Святая Гармония"', 'Музыкальное объединение "Harmony"', 'info@harmonymusic.org'),
    ('Танцевальный марафон', 'Развлечения', 'Танцевальный марафон с участием профессиональных танцоров', '2023-10-10 15:00:00', '2023-10-10 23:00:00', 25.00, 'Танцевальный клуб "DanceFever"', 'Танцевальная академия "Rhythm"', 'info@rhythmdance.com'),
    ('Кинофестиваль независимого кино', 'Кино', 'Фестиваль с премьерами независимых фильмов', '2023-10-15 17:30:00', '2023-10-15 23:00:00', 30.00, 'Кинотеатр "IndieCinema"', 'Независимое киностудия "CineFreedom"', 'info@cinefreedom.com'),
    ('Мастер-класс по фотографии', 'Искусство', 'Мастер-класс по фотографии с известным фотографом', '2023-10-20 14:00:00', '2023-10-20 16:00:00', 20.00, 'Фотостудия "PhotoCraft"', 'Фотоагентство "ArtisticLens"', 'info@artisticlens.com'),
    ('Мюзикл "Золушка"', 'Театр', 'Сказочный мюзикл о Золушке и её волшебной истории', '2023-10-25 18:30:00', '2023-10-25 21:30:00', 35.00, 'Театр оперы и балета', 'Театральная компания "Broadway"', 'info@broadwaytheatre.com'),
    ('Фестиваль народных традиций', 'Культура', 'Фестиваль с демонстрацией народных ремесел и традиций', '2023-10-30 11:00:00', '2023-10-30 19:00:00', 15.00, 'Народный парк "Folklore"', 'Ассоциация народных художеств "Traditions"', 'info@traditions.org'),
    ('Концерт поп-музыки', 'Музыка', 'Концерт популярной музыки с участием поп-звезд', '2023-11-05 20:00:00', '2023-11-05 23:00:00', 30.00, 'Концертный комплекс "PopStage"', 'Музыкальное агентство "PopStar"', 'info@popstaragency.com'),
    ('Выставка скульптуры', 'Искусство', 'Выставка скульптурных произведений современных художников', '2023-11-10 10:00:00', '2023-11-10 17:00:00', 20.00, 'Галерея современного искусства', 'Ассоциация скульпторов "ArtSculpt"', 'info@artsculptors.com'),
    ('Фестиваль электронной музыки', 'Музыка', 'Фестиваль с выступлениями лучших диджеев и музыкальных коллективов', '2023-11-15 18:00:00', '2023-11-15 23:59:00', 40.00, 'Открытая площадка "ElectroZone"', 'Музыкальный лейбл "ElectroBeats"', 'info@electrobeats.com'),
    ('Кулинарный фестиваль', 'Еда', 'Фестиваль с дегустацией блюд разных кухонь мира', '2023-11-20 16:00:00', '2023-11-20 21:00:00', 25.00, 'Парк кулинарии "FoodPark"', 'Кулинарная ассоциация "TasteWorld"', 'info@tasteworld.com'),
    ('Выставка фотографии', 'Искусство', 'Выставка фотографических работ профессиональных фотографов', '2023-11-25 12:00:00', '2023-11-25 19:00:00', 15.00, 'Фотогалерея "PhotoArt"', 'Фотографическое агентство "ArtLens"', 'info@artlensagency.com'),
    ('Мюзикл "Маленькая русская душа"', 'Театр', 'Мюзикл о любви, дружбе и приключениях в стиле народной культуры', '2023-11-30 19:30:00', '2023-11-30 22:30:00', 35.00, 'Национальный театр', 'Театральная студия "RussianSoul"', 'info@russiansoultheatre.com'),
    ('Фестиваль комедийного кино', 'Кино', 'Фестиваль с премьерами комедийных фильмов', '2023-12-05 17:30:00', '2023-12-05 23:00:00', 30.00, 'Кинотеатр "ComedyCinema"', 'Киностудия "LaughFactory"', 'info@laughfactory.com'),
    ('Мастер-класс по актерскому мастерству', 'Искусство', 'Мастер-класс по актерскому мастерству с известным актером', '2023-12-10 14:00:00', '2023-12-10 17:00:00', 20.00, 'Театральная студия "ActingSkills"', 'Актерское агентство "StageCraft"', 'info@stagecraft.com'),
    ('Опера "Кармен"', 'Театр', 'Известная опера о любви, страсти и предательстве', '2023-12-15 18:00:00', '2023-12-15 21:00:00', 40.00, 'Государственный оперный театр', 'Театральный коллектив "OperaDrama"', 'info@operadrama.com'),
    ('Выставка научных достижений', 'Наука', 'Выставка с презентацией новых научных открытий и технологий', '2023-12-20 10:00:00', '2023-12-20 18:00:00', 25.00, 'Научный центр "SciTech"', 'Научное общество "Innovation"', 'info@innovation.org'),
    ('Концерт рок-музыки', 'Музыка', 'Концерт легендарной рок-группы с исполнением хитов', '2023-12-25 20:30:00', '2023-12-25 23:30:00', 35.00, 'Стадион "RockArena"', 'Музыкальное агентство "RockStar"', 'info@rockstaragency.com'),
    ('Фестиваль народного танца', 'Танцы', 'Фестиваль с выступлениями танцевальных коллективов разных народностей', '2023-12-30 15:00:00', '2023-12-30 22:00:00', 30.00, 'Танцевальный комплекс "FolkDance"', 'Ассоциация народного танца "DanceWorld"', 'info@danceworld.org');

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

insert into comments (comment_text, comment_date, event_id, user_id) values ('Слабо', '2023-06-09', '1', '2');

insert into comments (comment_text, comment_date, event_id, user_id) values ('Мда, ждал лучше...', '2023-07-09', '1', '3');