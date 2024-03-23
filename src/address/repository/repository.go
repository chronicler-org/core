package addressRepository

import (
	"os"

	addressModel "github.com/chronicler-org/core/src/address/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func initDB() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&addressModel.Address{})
	return db
}

func InitAddressRepository() *AddressRepository {
	return &AddressRepository{
		db: initDB(),
	}
}

func (repository *AddressRepository) Create(address addressModel.Address) error {
	return repository.db.Model(&addressModel.Address{}).Create(address).Error
}

func (repository *AddressRepository) FindByID(id string) (addressModel.Address, error) {
	var address addressModel.Address
	err := repository.db.Find(&address, "id = ?", id).Error
	return address, err
}

func (repository *AddressRepository) Update(updatedAddress addressModel.Address) error {
	return repository.db.Save(updatedAddress).Error
}

func (repository *AddressRepository) Delete(id string) error {
	err := repository.db.Delete(&addressModel.Address{}, "id = ?", id).Error
	return err
}