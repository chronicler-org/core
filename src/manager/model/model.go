package managerModel

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager struct {
  ID uuid.UUID `json:"id" gorm:"primaryKey"`
  CPF string `json:"cpf"`
  Name string `json:"name"`
  Email string `json:"email"`
  Password string `json:"-"`
  BirthDate time.Time `json:"birth_date"`
}

type handler struct {
  db *gorm.DB
}

func initDB(dbURL string) *gorm.DB {
  db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
  }
  
  db.AutoMigrate(&Manager{})
  return db
}
