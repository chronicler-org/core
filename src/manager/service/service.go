package managerService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerExceptionMessage "github.com/chronicler-org/core/src/manager/messages"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	managerRepository "github.com/chronicler-org/core/src/manager/repository"
)

type ManagerService struct {
	managerRepository *managerRepository.ManagerRepository
}

func InitManagerService(r *managerRepository.ManagerRepository) *ManagerService {
	return &ManagerService{
		managerRepository: r,
	}
}

func (service *ManagerService) FindByID(id string) (managerModel.Manager, error) {
	result, err := service.managerRepository.FindByID(id)
	manager, _ := result.(*managerModel.Manager)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *manager, appException.NotFoundException(managerExceptionMessage.MANAGER_NOT_FOUND)
	}
	return *manager, nil
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

	err = service.managerRepository.Create(model)

	return model, err
}

func (service *ManagerService) Update(id string, dto managerDTO.UpdateManagerDTO) (managerModel.Manager, error) {
	managerExists, err := service.FindByID(id)
	if err != nil {
		return managerModel.Manager{}, err
	}

	appUtil.UpdateModelFromDTO(&managerExists, dto)
	if dto.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
		if err == nil {
			managerExists.Password = string(newPassword)
		}
	}

	managerExists.UpdatedAt = time.Now()
	err = service.managerRepository.Update(managerExists)
	return managerExists, err
}

func (service *ManagerService) FindAll(dto appDto.PaginationDTO) (int64, []managerModel.Manager, error) {
	var managers []managerModel.Manager
	totalCount, err := service.managerRepository.FindAll(dto.GetLimit(), dto.GetPage(), &managers)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, managers, nil
}

func (service *ManagerService) Delete(id string) (managerModel.Manager, error) {
	managerExists, err := service.FindByID(id)
	if err != nil {
		return managerModel.Manager{}, err
	}

	err = service.managerRepository.Delete(id)
	if err != nil {
		return managerModel.Manager{}, err
	}
	return managerExists, nil
}
