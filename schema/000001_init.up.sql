CREATE TABLE IF NOT EXISTS rubric
(
    id                  serial          not null unique,
    name                varchar(255)    not null unique,
    parent_rubric_id    int             null
);
ALTER TABLE rubric ADD CONSTRAINT rubric_fk FOREIGN KEY (parent_rubric_id) REFERENCES rubric (id);

CREATE TABLE IF NOT EXISTS city
(
    id          serial          not null unique,
    name        varchar(255)    not null unique
);
CREATE TABLE IF NOT EXISTS street
(
    id          serial                      not null unique,
    city_id     int references city (id)    on delete cascade not null,
    name        varchar(255)                not null
);
ALTER TABLE street ADD CONSTRAINT city_id_name UNIQUE (city_id, name);

CREATE TABLE IF NOT EXISTS building
(
    id              serial                      not null unique,
    city_id         int references city (id)    on delete cascade not null,
    street_id       int references street (id)  on delete cascade not null,
    house           int                         null,
    point           point                       default '0.00, 0.00' not null
);
ALTER TABLE building ADD CONSTRAINT city_id_street_id_house UNIQUE (city_id, street_id, house);

CREATE TABLE IF NOT EXISTS company
(
    id          serial                          not null unique,
    name        varchar(255)                    not null,
    building_id int references building (id)    on delete cascade not null
);
ALTER TABLE company ADD CONSTRAINT name_building_id UNIQUE (name, building_id);

CREATE TABLE IF NOT EXISTS company_rubric
(
    company_id  int references company (id)     on delete cascade not null,
    rubric_id   int references rubric (id)      on delete cascade not null
);
ALTER TABLE company_rubric ADD CONSTRAINT company_id_rubric_id UNIQUE (company_id, rubric_id);

CREATE TABLE IF NOT EXISTS phone
(
    id          serial          not null unique,
    company_id  int references company (id)     on delete cascade not null,
    number      varchar(255)    not null
);
ALTER TABLE phone ADD CONSTRAINT company_id_number UNIQUE (company_id, number);