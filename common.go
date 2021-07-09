package common

// Company - Фирма представляет собой карточку организации в справочнике и должна содержать в себе следующую информацию:
type Company struct {
	Id       int    `json:"id"`
	Name string `json:"name"`
	Phones []string `json:"phones"`
	Building Building `json:"building"`
	Rubric []Rubric `json:"rubric"`
}

// Building - здание содержит в себе как минимум информацию о конкретном здании, а именно
type Building struct {
	Address Address `json:"address"`
	Coordinates Coordinates `json:"coordinates"`
}
// Address - адрес
type Address struct {
	City string `json:"city"`
	Street string `json:"street"`
	House int `json:"house"`
}
// Coordinates - координаты (широта, долгота)
type Coordinates struct {
	Longitude float32 `json:"longitude"`
	Latitude float32 `json:"latitude"`
}

// Rubric - Рубрика
type Rubric struct {
	Name string `json:"name"`
	ParentRubric *Rubric `json:"parent_rubric"`
}