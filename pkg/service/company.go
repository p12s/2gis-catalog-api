package service

import (
	"errors"
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type CompanyService struct {
	repo              repository.Company
	repoBuilding      repository.Building
	repoPhone         repository.Phone
	repoRubric        repository.Rubric
	repoCompanyRubric repository.CompanyRubric
}

func NewCompanyService(
	repo repository.Company,
	repoBuilding repository.Building,
	repoPhone repository.Phone,
	repoRubric repository.Rubric,
	repoCompanyRubric repository.CompanyRubric,
) *CompanyService {
	return &CompanyService{
		repo:              repo,
		repoBuilding:      repoBuilding,
		repoPhone:         repoPhone,
		repoRubric:        repoRubric,
		repoCompanyRubric: repoCompanyRubric,
	}
}

// Create - создание компании
// Считаем, что все города и улицы уже занесены в БД, здесь приходят их id, которые валидируем
/* Пришедшее содержимое запроса в объекте common.Company:
{
    "name": "Rambler", // создаем компанию (post)
    "phones": [
        {
            "company_id": 21, // игнорируем
            "number": "32-32-32" // записываем номера в конце - когда имеем id компании
        },
        {
            "company_id": null,
            "number": "33-33-33"
        }
    ],
    "building": {
        "id": null, // игнорируем
        "city_id": 1, // если по совокупности полей city_id  street_id street_id нет записи - создаем, иначе берем id готового
        "street_id": 2,
        "house": 12,
        "point": "(1.234567890, -90.87654321)" // записываем, если пришло (не проверяем)
    },
    "rubric": [
        {
            "id": null, // id всегда есть и он существует (не проверяем)
            "name": "Новая рубрика",
            "parent_rubric_id": null
        }
    ]
}
*/
func (c *CompanyService) Create(requestCompany common.CompanyCreateRequest) (int, error) {
	err := c.repo.OpenTransaction()
	if err != nil {
		return 0, err
	}

	// проверка компании на существование
	isCompanyExists, err := c.repo.Exists(requestCompany.Name)
	if err != nil {
		rollbackErr := c.repo.RollbackTransaction()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	if isCompanyExists {
		return 0, errors.New("company with name " + requestCompany.Name + " already exists")
	}

	// проверка здания на существование и его добавление по-условию (для простоты считаем, что id города и улиц существуют)
	existsBuilding, err := c.repoBuilding.GetByCityStreetHouse(requestCompany.Building.CityId, requestCompany.Building.StreetId, requestCompany.Building.House)
	if err != nil {
		rollbackErr := c.repo.RollbackTransaction()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}
	if existsBuilding.Id == 0 {
		existsBuilding.Id, err = c.repoBuilding.Create(
			requestCompany.Building.CityId,
			requestCompany.Building.StreetId,
			requestCompany.Building.House,
			requestCompany.Building.Point)
		rollbackErr := c.repo.RollbackTransaction()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		if err != nil {
			return 0, err
		}
	}
	requestCompany.Building.Id = existsBuilding.Id

	// добавление компании
	companyId, err := c.repo.Create(requestCompany)
	if err != nil {
		rollbackErr := c.repo.RollbackTransaction()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	// проверка рубрик на существование и их добавление по-условию
	for _, rubric := range requestCompany.Rubric {
		if rubric.Id == 0 && rubric.Name != "" {
			// id не пришел - ищем в БД по названию
			foundRubric, err := c.repoRubric.FindByName(rubric.Name)
			if err != nil {
				rollbackErr := c.repo.RollbackTransaction()
				if rollbackErr != nil {
					return 0, rollbackErr
				}
				return 0, err
			}
			if foundRubric.Id != 0 {
				rubric.Id = foundRubric.Id
				rubric.ParentRubricId = foundRubric.ParentRubricId
				rubric.Name = foundRubric.Name
			}
			createdRubricId, err := c.repoRubric.Create(rubric.Name, rubric.ParentRubricId)
			if err != nil {
				rollbackErr := c.repo.RollbackTransaction()
				if rollbackErr != nil {
					return 0, rollbackErr
				}
				return 0, err
			}
			rubric.Id = createdRubricId
		}

		err = c.repoCompanyRubric.Create(companyId, rubric.Id)
		if err != nil {
			rollbackErr := c.repo.RollbackTransaction()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err
		}
	}

	// добавление номеров телефонов
	for _, phone := range requestCompany.Phones {
		phone.CompanyId = companyId
		err = c.repoPhone.Create(phone.CompanyId, phone.Number)
		if err != nil {
			rollbackErr := c.repo.RollbackTransaction()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err
		}
	}

	return companyId, err
}

func (c *CompanyService) GetById(companyId int) (common.Company, error) {
	return c.repo.GetById(companyId)
}

func (c *CompanyService) GetAllInBuilding(buildingId int) ([]common.Company, error) {
	return c.repo.GetAllInBuilding(buildingId)
}

func (c *CompanyService) GetAllByRubricId(rubricId int) ([]common.Company, error) {
	return c.repo.GetAllByRubricId(rubricId)
}

func (c *CompanyService) Delete(companyId int) error {
	return c.repo.Delete(companyId)
}
