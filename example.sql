CREATE database gis;
CREATE TABLE rubric (
                        id BIGSERIAL PRIMARY KEY,
                        parent_id INT REFERENCES rubric(id),
                        name VARCHAR(30) NULL NULL
);
CREATE INDEX rubric_name ON rubric(name);

INSERT INTO rubric (name) VALUES ('Еда'), ('Автомобили'), ('Спорт');
INSERT INTO rubric (parent_id, name) VALUES (1, 'Полуфабрикаты оптом'), (1, 'Мясная продукция');
INSERT INTO rubric (parent_id, name) VALUES (2, 'Грузовые'), (2, 'Легковые');
INSERT INTO rubric (parent_id, name) VALUES (7, 'Запчасти для подвески'), (7, 'Шины/Диски');

CREATE TABLE rubric_closure (
                                id BIGSERIAL PRIMARY KEY,
                                parent_id INT REFERENCES rubric(id),
                                child_id INT REFERENCES rubric(id)
);
CREATE INDEX rubric_closure_parent_id ON rubric_closure(parent_id);
INSERT INTO rubric_closure (parent_id, child_id) VALUES (1,1), (1,4), (1,5), (2,2), (2,6), (2,7), (3,3), (4,4), (5,5), (6,6), (7,7), (7,8), (7,9), (8,8), (9,9);

CREATE TABLE building (
                          id BIGSERIAL PRIMARY KEY,
                          country VARCHAR (2) NOT NULL,
                          city VARCHAR (30) NOT NULL,
                          street VARCHAR (30) NOT NULL,
                          building_number VARCHAR (5) NOT NULL,
                          coordinates POINT NOT NULL,
                          postcode VARCHAR (6) NOT NULL
);
CREATE INDEX building_country ON building(country);
CREATE INDEX building_city ON building(city);
CREATE INDEX building_street ON building(street);
CREATE INDEX building_building_number ON building(building_number);
CREATE INDEX building_coordinates ON building USING gist(coordinates);
INSERT INTO building (country, city, street, building_number, coordinates, postcode) VALUES
('RU', 'Санкт-Петербург', 'Гороховая', '13', POINT(59.93379788302153, 30.313648028544524), '190000'),
('RU', 'Москва', 'Красная площадь', '1', POINT(55.753625937403406, 37.619999975260306), '109012'),
('RU', 'Ковров', 'Брюсова', '19', POINT(56.35468657988904, 41.30583580257589), '601901'),
('RU', 'Рыбинск', 'Яблочкина', '11', POINT(58.03329079498624, 38.84243761185397), '152919')
;

CREATE TABLE company (
                         id BIGSERIAL PRIMARY KEY,
                         phones jsonb,
                         name VARCHAR (100) NOT NULL,
                         building_id INT REFERENCES building(id)
);
CREATE INDEX company_phones ON company USING gin(phones);
CREATE INDEX company_name ON company(name);
CREATE INDEX company_building_id ON company(building_id);
INSERT INTO company (phones, name, building_id) VALUES
('["79346759364"]', 'ООО Норипт', 1),
('["79547560132", "74235679984"]', 'ОАО Пивоваренный завод имени Степана Разина', 1),
('["79997348321"]', 'ООО Просвящение', 2),
('["79532136749", "74574736534"]', 'ООО Oyster', 3),
('["79997845321"]', 'ООО Илим', 3),
('["73413467432"]', 'ОАО Трансаэро', 4)
;

CREATE TABLE company_rubric (
                                company_id INT REFERENCES company(id),
                                rubric_id INT REFERENCES rubric(id)
);
CREATE INDEX company_rubric_company_id ON company_rubric(company_id);
CREATE INDEX company_rubric_rubric_id ON company_rubric(rubric_id);