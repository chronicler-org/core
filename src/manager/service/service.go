package managerService

import (
	"time"

	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	managerRepository "github.com/chronicler-org/core/src/manager/repository"
	serviceErrors "github.com/chronicler-org/core/src/utils/errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ManagerService struct {
	repository *managerRepository.ManagerRepository
	validate   *validator.Validate
}

func InitManagerService(r *managerRepository.ManagerRepository, v *validator.Validate) *ManagerService {
	return &ManagerService{
		repository: r,
		validate:   v,
	}
}

func (service *ManagerService) FindByID(id string) (managerModel.Manager, error) {
	manager, err := service.repository.FindByID(id)

	return manager, err
}

func (service *ManagerService) Create(dto managerDTO.CreateManagerDTO) (managerModel.Manager, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return managerModel.Manager{}, err
	}

	model := managerModel.Manager{
		ID:        uuid.New(),
		Name:      dto.Name,
		CPF:       dto.CPF,
		Email:     dto.Email,
		Password:  string(newPassword),
		BirthDate: dto.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.repository.Create(model)

	return model, err
}

func (service *ManagerService) Update(id string, dto managerDTO.UpdateManagerDTO) (managerModel.Manager, error) {
	updatedManager, err := service.repository.FindByID(id)
	if err != nil {
		return updatedManager, err
	}
	if updatedManager.ID == uuid.Nil {
		return updatedManager, serviceErrors.NewError("Gerente não encontrado")
	}

	if dto.CPF != "" {
		if dto.ValidateCPF() {
			updatedManager.CPF = dto.CPF
		}
		return updatedManager, serviceErrors.NewError("novo CPF é inválido")
	}
	if dto.Name != "" {
		updatedManager.Name = dto.Name
	}
	if dto.Email != "" {
		if dto.ValidateEmail() {
			updatedManager.Email = dto.Email
		} else {
			return updatedManager, serviceErrors.NewError("novo email é Email invalido")
		}
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

func (service *ManagerService) FindAll() ([]managerModel.Manager, error) {
	return service.repository.FindAll()
}

func (service *ManagerService) Delete(id string) error {
	return service.repository.Delete(id)
}
