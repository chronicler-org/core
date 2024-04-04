package managerService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerExceptionMessage "github.com/chronicler-org/core/src/manager/messages"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	managerRepository "github.com/chronicler-org/core/src/manager/repository"
	teamModel "github.com/chronicler-org/core/src/team/model"
	teamService "github.com/chronicler-org/core/src/team/service"
)

type ManagerService struct {
	managerRepository *managerRepository.ManagerRepository
	teamService       *teamService.TeamService
}

func InitManagerService(managerRepository *managerRepository.ManagerRepository, teamService *teamService.TeamService) *ManagerService {
	return &ManagerService{
		managerRepository: managerRepository,
		teamService:       teamService,
	}
}

func (service *ManagerService) FindByID(id string) (managerModel.Manager, error) {
	result, err := service.managerRepository.FindOneByField("ID", id, "Team")
	manager, _ := result.(*managerModel.Manager)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *manager, appException.NotFoundException(managerExceptionMessage.MANAGER_NOT_FOUND)
	}
	return *manager, nil
}

func (service *ManagerService) FindManagerByEmail(email string) (managerModel.Manager, error) {
	result, err := service.managerRepository.FindOneByField("Email", email, "Team")
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

	team := teamModel.Team{}
	if dto.TeamId != "" {
		team, err = service.teamService.FindByID(dto.TeamId)
		if err != nil {
			return managerModel.Manager{}, err
		}
		model.TeamID = team.ID
	}

	err = service.managerRepository.Create(model)
	model.Team = team
	return model, err
}

func (service *ManagerService) Update(id string, dto managerDTO.UpdateManagerDTO) (managerModel.Manager, error) {
	managerExists, err := service.FindByID(id)
	if err != nil {
		return managerModel.Manager{}, err
	}

	appUtil.UpdateModelFromDTO(&managerExists, &dto)
	if dto.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
		if err == nil {
			managerExists.Password = string(newPassword)
		}
	}
	managerExists.UpdatedAt = time.Now()

	if dto.TeamId != "" {
		team, err := service.teamService.FindByID(dto.TeamId)
		if err != nil {
			return managerModel.Manager{}, err
		}
		err = service.managerRepository.Update(managerExists)
		managerExists.Team = team
		return managerExists, err
	} else {
		err = service.managerRepository.Update(managerExists)
		return managerExists, err
	}
}

func (service *ManagerService) FindAll(queryCustomerDTO managerDTO.QueryManagerDTO) (int64, []managerModel.Manager, error) {
	var managers []managerModel.Manager
	totalCount, err := service.managerRepository.FindAll(queryCustomerDTO, &managers, "Team")
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

	err = service.managerRepository.Delete("ID", id)
	if err != nil {
		return managerModel.Manager{}, err
	}
	return managerExists, nil
}
