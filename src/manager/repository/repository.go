package managerRepository

import (
	"os"

	managerModel "github.com/chronicler-org/core/src/manager/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ManagerRepository struct {
	db *gorm.DB
}

func initDB() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&managerModel.Manager{})
	return db
}

func InitManagerRepository() *ManagerRepository {
	return &ManagerRepository{
		db: initDB(),
	}
}

func (repository *ManagerRepository) Create(manager managerModel.Manager) error {
	return repository.db.Model(&managerModel.Manager{}).Create(manager).Error
}

func (repository *ManagerRepository) FindByID(id string) (managerModel.Manager, error) {
	var manager managerModel.Manager
	err := repository.db.Find(&manager, "id = ?", id).Error
	return manager, err
}

func (repository *ManagerRepository) Update(updatedManager managerModel.Manager) error {
	return repository.db.Save(updatedManager).Error
}

func (repository *ManagerRepository) FindAll() ([]managerModel.Manager, error) {
	var managers []managerModel.Manager
	err := repository.db.Find(&managers).Error
	return managers, err
}

func (repository *ManagerRepository) Delete(id string) error {
	err := repository.db.Delete(&managerModel.Manager{}, "id = ?", id).Error
	return err
}
