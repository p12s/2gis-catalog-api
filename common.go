package common

// Company - Фирма представляет собой карточку организации в справочнике и должна содержать в себе следующую информацию:
type Company struct {
	Id       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name" binding:"required"`
	Phones   []Phone  `json:"phones"`
	Building Building `json:"building" binding:"required"`
	Rubric   []Rubric `json:"rubric" binding:"required"`
}

type Phone struct {
	Id        int    `json:"-" db:"id"`
	CompanyId int    `json:"company_id" db:"company_id"`
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

type City struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}
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

type CompanyCreateRequest struct {
	Name   string `json:"name"`
	Phones []struct {
		CompanyId int    `json:"-"`
		Number    string `json:"number"`
	} `json:"phones"`
	Building struct {
		Id       int    `json:"-"`
		CityId   int    `json:"city_id"`
		StreetId int    `json:"street_id"`
		House    int    `json:"house"`
		Point    string `json:"point"`
	} `json:"building"`
	Rubric []struct {
		Id             int    `json:"id"`
		Name           string `json:"-"`
		ParentRubricId int    `json:"-"`
	} `json:"rubric"`
}
