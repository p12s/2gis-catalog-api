package common

// Company - Фирма представляет собой карточку организации в справочнике и должна содержать в себе следующую информацию:
type Company struct {
	Id       int    `json:"id"`
	Name string `json:"name"`
	Phones []Phone `json:"phones"`
	Building Building `json:"building"`
	Rubric []Rubric `json:"rubric"`
}

type Phone struct {
	Id       int    `json:"-"`
	Number string `json:"number"`
}

// Building - здание содержит в себе как минимум информацию о конкретном здании, а именно
type Building struct {
	Address Address `json:"address"`
	Point string `json:"point"`
	//Longitude float32 `json:"longitude"`
	//Latitude float32 `json:"latitude"`
}

// Address - адрес
type Address struct {
	Id int `json:"-"`
	City City `json:"city"`
	Street Street `json:"street"`
	House int `json:"house"`
}
type City struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
type Street struct {
	Id int `json:"id"`
	CityId int `json:"city_id"`
	Name string `json:"name"`
}

// Rubric - Рубрика
type Rubric struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ParentRubricId int `json:"parent_rubric_id"`
}