package common

// Company - Фирма представляет собой карточку организации в справочнике и должна содержать в себе следующую информацию:
type Company struct {
	Id       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name" binding:"required"`
	Phones   []Phone  `json:"phones"`
	Building Building `json:"building"`
	Rubric   []Rubric `json:"rubric"`
}

// Phone - телефон
type Phone struct {
	Id        int    `json:"-" db:"id"`
	CompanyId int    `json:"-" db:"company_id"`
	Number    string `json:"number" db:"number"`
}

// Building - здание содержит в себе как минимум информацию о конкретном здании, а именно
type Building struct {
	Id       int    `json:"-" db:"id"`
	CityId   int    `json:"city_id" db:"city_id" binding:"required"`
	City     string `json:"city" db:"-"` // название города (если есть), используем при создании записей в БД, если такого еще нет
	StreetId int    `json:"street_id" db:"street_id" binding:"required"`
	Street   string `json:"street" db:"-"` // название улицы (если есть), используем при создании записей в БД, если такой улици еще нет
	House    int    `json:"house" db:"house" binding:"required"`
	Point    string `json:"point" db:"point"`
}

// City - город
type City struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

// Street - улица
type Street struct {
	Id     int    `json:"id"`
	CityId int    `json:"city_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

// Rubric - Рубрика
type Rubric struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	ParentRubricId int    `json:"parent_rubric_id" db:"parent_rubric_id"`
}

// CompanyCreateRequest - пришедший запрос на создание компании
type CompanyCreateRequest struct {
	Name   string `json:"name"`
	Phones []struct {
		CompanyId int    `json:"-"`
		Number    string `json:"number"` // номер телефона
	} `json:"phones"`
	Building struct {
		Id       int    `json:"-"`
		CityId   int    `json:"city_id"`   // id города должен существовать в БД
		StreetId int    `json:"street_id"` // id улицы с привязкой к городу должен существовать в БД
		House    int    `json:"house"`
		Point    string `json:"point"` // в формате "(1.00234567, -90.00876211)"
	} `json:"building"`
	Rubric []struct {
		Id             int    `json:"id"` // id рубрики должен существовать в БД
		Name           string `json:"-"`
		ParentRubricId int    `json:"-"`
	} `json:"rubric"`
}

// CompanyResponse - объект компании, который отдаем как результат запроса
type CompanyResponse struct {
	Name   string  `json:"name"`
	Phones []Phone `json:"phones"`
}

// BuildingCreateRequest - пришедший запрос на создание здания
type BuildingCreateRequest struct {
	CityId   int    `json:"city_id" db:"city_id"`
	City     string `json:"city" db:"-"` // название города (если есть), используем при создании записей в БД, если такого еще нет
	StreetId int    `json:"street_id" db:"street_id"`
	Street   string `json:"street" db:"-"` // название улицы (если есть), используем при создании записей в БД, если такой улици еще нет
	House    int    `json:"house" db:"house" binding:"required"`
	Point    string `json:"point" db:"point"`
}
