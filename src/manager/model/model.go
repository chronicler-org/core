package managerModel

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager struct {
  ID        uuid.UUID `gorm:"primarykey"`
  CPF       string    `json:"cpf"`
  Name      string    `json:"name"`
  Email     string    `json:"email"`
  Password  string    `json:"-"` 
  BirthDate time.Time `json:"birth_date"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type CreateManagerDTO struct {
  CPF       string    `validate:"required,max=11" json:"cpf"`
  Name      string    `validate:"required,max=50" json:"name"`
  Email     string    `validate:"required,max=50" json:"email"` 
  Password  string    `validate:"required,min=8" json:"Password"`
  BirthDate time.Time `validate:"required" json:"birth_date"`
}

type UpdateManagerDTO struct {
  CPF       string    `json:"cpf,omitempty"`
  Name      string    `json:"name,omitempty"`
  Email     string    `json:"email,omitempty"` 
  Password  string    `json:"Password,omitempty"`
  BirthDate time.Time `json:"birth_date,omitempty"`
}

type Handler struct {
  db *gorm.DB
}

func initDB () *gorm.DB {
  dbURL := os.Getenv("DATABASE_URL")
  db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
  }
  
  db.AutoMigrate(&Manager{})
  return db
}

func InitHandler () *Handler {
  return &Handler{
    db: initDB(),
  }
}

func (handler *Handler) Create (managerDTO *CreateManagerDTO) (uuid.UUID, error) {
  var newID uuid.UUID = uuid.New()
  encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(managerDTO.Password), 10) 
  if err != nil {
    return newID, err
  }
  manager := &Manager{
    ID:         newID,
    CPF:        managerDTO.CPF,
    Name:       managerDTO.Name,
    Email:      managerDTO.Email,
    Password:   string(encryptedPassword),
    BirthDate:  managerDTO.BirthDate,
    CreatedAt:  time.Now(),
    UpdatedAt:  time.Now(),
  } 
  return newID, handler.db.Model(&Manager{}).Create(manager).Error
}

func (handler *Handler) FindByID (id string) (Manager, error) {
  var manager Manager
  err := handler.db.Find(&manager, "id = ?", id).Error
  return manager, err
}

func (handler *Handler) Update (manager *Manager, updateData *CreateManagerDTO) error {
  if updateData.CPF != "" {
    manager.CPF = updateData.CPF
  }
  if updateData.Name != "" {
    manager.Name = updateData.Name
  }
  if updateData.Email != "" {
    manager.Email = updateData.Email
  }
  if updateData.Password != "" {
    newPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.CPF), 10)
    if err != nil {
      return err
    }
    manager.Password = string(newPassword)
  }
  if !updateData.BirthDate.IsZero() {
    manager.BirthDate = updateData.BirthDate
  }
  manager.UpdatedAt = time.Now()
  err := handler.db.Save(manager).Error
  return err
}

func (handler *Handler) FindAll () ([]Manager, error) {
  var managers []Manager
  err := handler.db.Find(&managers).Error
  return managers, err
}

func (handler *Handler) Delete (id string) error {
  err := handler.db.Delete(&Manager{}, "id = ?", id).Error
  return err
}
