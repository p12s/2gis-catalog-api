INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Транспорт', NULL);

INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Автомобили', (SELECT id FROM rubric WHERE name = 'Транспорт')),
(DEFAULT, 'Мотоциклы и мототехника', (SELECT id FROM rubric WHERE name = 'Транспорт')),
(DEFAULT, 'Грузовики и спецтехника', (SELECT id FROM rubric WHERE name = 'Транспорт')),
(DEFAULT, 'Водный транспорт', (SELECT id FROM rubric WHERE name = 'Транспорт')),
(DEFAULT, 'Запчасти и аксессуары', (SELECT id FROM rubric WHERE name = 'Транспорт'));

INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Недвижимость', NULL);
INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Квартиры', (SELECT id FROM rubric WHERE name = 'Недвижимость')),
(DEFAULT, 'Комнаты', (SELECT id FROM rubric WHERE name = 'Недвижимость')),
(DEFAULT, 'Земельные участки', (SELECT id FROM rubric WHERE name = 'Недвижимость')),
(DEFAULT, 'Гаражи и машиноместа', (SELECT id FROM rubric WHERE name = 'Недвижимость')),
(DEFAULT, 'Коммерческая недвижимость', (SELECT id FROM rubric WHERE name = 'Недвижимость')),
(DEFAULT, 'Недвижимость за рубежом', (SELECT id FROM rubric WHERE name = 'Недвижимость'));

INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Работа', NULL);
INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Вакансии', (SELECT id FROM rubric WHERE name = 'Работа')),
(DEFAULT, 'Резюме', (SELECT id FROM rubric WHERE name = 'Работа')),
(DEFAULT, 'Услуги', (SELECT id FROM rubric WHERE name = 'Работа'));

INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Личные вещи', NULL);
INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Одежда, обувь, аксессуары', (SELECT id FROM rubric WHERE name = 'Личные вещи')),
(DEFAULT, 'Детская одежда и обувь', (SELECT id FROM rubric WHERE name = 'Личные вещи')),
(DEFAULT, 'Товары для детей и игрушки', (SELECT id FROM rubric WHERE name = 'Личные вещи')),
(DEFAULT, 'Часы и украшения', (SELECT id FROM rubric WHERE name = 'Личные вещи')),
(DEFAULT, 'Красота и здоровье', (SELECT id FROM rubric WHERE name = 'Личные вещи'));

INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Хобби и отдых', NULL);
INSERT INTO rubric (id, name, parent_rubric_id) VALUES
(DEFAULT, 'Билеты и путешествия', (SELECT id FROM rubric WHERE name = 'Хобби и отдых')),
(DEFAULT, 'Велосипеды', (SELECT id FROM rubric WHERE name = 'Хобби и отдых')),
(DEFAULT, 'Книги и журналы', (SELECT id FROM rubric WHERE name = 'Хобби и отдых')),
(DEFAULT, 'Коллекционирование', (SELECT id FROM rubric WHERE name = 'Хобби и отдых')),
(DEFAULT, 'Музыкальные инструменты', (SELECT id FROM rubric WHERE name = 'Хобби и отдых')),
(DEFAULT, 'Охота и рыбалка', (SELECT id FROM rubric WHERE name = 'Хобби и отдых')),
(DEFAULT, 'Спорт и отдых', (SELECT id FROM rubric WHERE name = 'Хобби и отдых'));

----------------------------------------

INSERT INTO city (id, name) VALUES
(DEFAULT, 'Иваново'),
(DEFAULT, 'Керчь'),
(DEFAULT, 'Севастополь'),
(DEFAULT, 'Киев'),
(DEFAULT, 'Магадан'),
(DEFAULT, 'Калуга');

INSERT INTO street (id, city_id, name) VALUES
(DEFAULT, (SELECT id FROM city WHERE name = 'Иваново'), 'Ленина'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Иваново'), 'Вениамина Тарасова'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Иваново'), 'Искандера Беллинского'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Иваново'), 'Веры Брежневой'),

(DEFAULT, (SELECT id FROM city WHERE name = 'Керчь'), 'Ленина'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Керчь'), 'Вениамина Тарасова'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Керчь'), 'Искандера Беллинского'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Керчь'), 'Веры Брежневой'),

(DEFAULT, (SELECT id FROM city WHERE name = 'Севастополь'), 'Ленина'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Севастополь'), 'Вениамина Тарасова'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Севастополь'), 'Искандера Беллинского'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Севастополь'), 'Веры Брежневой'),

(DEFAULT, (SELECT id FROM city WHERE name = 'Киев'), 'Ленина'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Киев'), 'Вениамина Тарасова'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Киев'), 'Искандера Беллинского'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Киев'), 'Веры Брежневой'),

(DEFAULT, (SELECT id FROM city WHERE name = 'Магадан'), 'Ленина'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Магадан'), 'Вениамина Тарасова'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Магадан'), 'Искандера Беллинского'),
(DEFAULT, (SELECT id FROM city WHERE name = 'Магадан'), 'Веры Брежневой');

----------------------------------------
INSERT INTO building (id, city_id, street_id, house, point) VALUES
(DEFAULT, (SELECT id FROM city WHERE name = 'Иваново'), (SELECT id FROM street WHERE name = 'Ленина' AND city_id = (SELECT id FROM city WHERE name = 'Иваново')), 13, '-12.0987654321, 173.0987654321');

INSERT INTO company (id, name, building_id) VALUES
(DEFAULT, 'iSpring', currval('building_id_seq'));

INSERT INTO phone (id, company_id, number) VALUES
(DEFAULT, (SELECT id FROM company WHERE name = 'iSpring'), '11-11-11'),
(DEFAULT, (SELECT id FROM company WHERE name = 'iSpring'), '22-22-22'),
(DEFAULT, (SELECT id FROM company WHERE name = 'iSpring'), '33-33-33');

INSERT INTO company_rubric (company_id, rubric_id) VALUES
((SELECT id FROM company WHERE name = 'iSpring'), (SELECT id FROM rubric WHERE name = 'Запчасти и аксессуары')),
((SELECT id FROM company WHERE name = 'iSpring'), (SELECT id FROM rubric WHERE name = 'Резюме')),
((SELECT id FROM company WHERE name = 'iSpring'), (SELECT id FROM rubric WHERE name = 'Велосипеды'));

----------------------------------------