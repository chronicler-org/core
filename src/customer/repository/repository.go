package customerRepository

import (
	"os"

	customerModel "github.com/chronicler-org/core/src/customer/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func initDB() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&customerModel.Customer{})
	return db
}

func InitCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: initDB(),
	}
}

func (repository *CustomerRepository) Create(customer customerModel.Customer) error {
	return repository.db.Model(&customerModel.Customer{}).Create(customer).Error
}

func (repository *CustomerRepository) FindByID(id string) (customerModel.Customer, error) {
	var customer customerModel.Customer
	err := repository.db.Find(&customer, "id = ?", id).Error
	return customer, err
}

func (repository *CustomerRepository) Update(updatedCustomer customerModel.Customer) error {
	return repository.db.Save(updatedCustomer).Error
}

func (repository *CustomerRepository) FindAll() ([]customerModel.Customer, error) {
	var customers []customerModel.Customer
	err := repository.db.Find(&customers).Error
	return customers, err
}

func (repository *CustomerRepository) Delete(id string) error {
	err := repository.db.Delete(&customerModel.Customer{}, "id = ?", id).Error
	return err
}
