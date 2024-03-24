package managerRepository

import (
	managerModel "github.com/chronicler-org/core/src/manager/model"
	"gorm.io/gorm"
)

type ManagerRepository struct {
	db *gorm.DB
}

func InitManagerRepository(db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{
		db: db,
	}
}

func (repository *ManagerRepository) Create(manager managerModel.Manager) error {
	return repository.db.Model(&managerModel.Manager{}).Create(manager).Error
}

func (repository *ManagerRepository) FindByID(id string) (managerModel.Manager, error) {
	var manager managerModel.Manager
	err := repository.db.Where("id = ?", id).First(&manager).Error
	return manager, err
}

func (repository *ManagerRepository) Update(updatedManager managerModel.Manager) error {
	return repository.db.Save(updatedManager).Error
}

func (repository *ManagerRepository) FindAll(limit, page int) ([]managerModel.Manager, error) {
	var managers []managerModel.Manager
	offset := (page - 1) * limit
	err := repository.db.Limit(limit).Offset(offset).Find(&managers).Error
	return managers, err
}

func (repository *ManagerRepository) Count() (int64, error) {
	var count int64
	err := repository.db.Model(&managerModel.Manager{}).Count(&count).Error
	return count, err
}

func (repository *ManagerRepository) Delete(id string) error {
	return repository.db.Delete(&managerModel.Manager{}, "id = ?", id).Error
}
