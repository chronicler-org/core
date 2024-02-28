package managerService

import (
	"time"

	"github.com/chronicler-org/core/src/manager/dto"
	"github.com/chronicler-org/core/src/manager/model"
	"github.com/chronicler-org/core/src/manager/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ManagerService struct {
  repository *managerRepository.ManagerRepository
}

func InitManagerService(r *managerRepository.ManagerRepository) *ManagerService{
  return &ManagerService{
    repository: r,
  }
}

func (service *ManagerService) FindByID (id string) (managerModel.Manager, error) {
  return service.repository.FindByID(id)
}

func (service *ManagerService) Create (dto managerDTO.CreateManagerDTO) (uuid.UUID, error) {
  newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
  if err != nil {
    return uuid.Nil, err
  }

  model := managerModel.Manager {
    ID: uuid.New(),
    Name: dto.Name,
    CPF: dto.CPF,
    Password: string(newPassword),
    BirthDate: dto.BirthDate,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  } 

  err = service.repository.Create(model)

  return model.ID, err
}

func (service *ManagerService) Update (id string, dto managerDTO.UpdateManagerDTO) (managerModel.Manager, error) {
  updatedManager, err := service.repository.FindByID(id) 
  if err != nil {
    return updatedManager, err
  }
  if updatedManager.ID == uuid.Nil {
    return updatedManager, err
  }

  if dto.CPF != "" {
    updatedManager.CPF = dto.CPF
  }  
  if dto.Name != "" {
    updatedManager.Name = dto.Name
  }
  if dto.Email != "" {
    updatedManager.Email = dto.Email
  }
  if dto.Password != "" {
    newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.CPF), 10)
    if err != nil {
      return managerModel.Manager{}, err
    }
    updatedManager.Password = string(newPassword)
  }
  if !dto.BirthDate.IsZero() {
    updatedManager.BirthDate = dto.BirthDate
  }
  updatedManager.UpdatedAt = time.Now()

  err = service.repository.Update(updatedManager)

  return updatedManager, err
}

func (service *ManagerService) FindAll () ([]managerModel.Manager, error) {
  return service.repository.FindAll()
}

func (service *ManagerService) Delete (id string) error {
  return service.repository.Delete(id)
}
