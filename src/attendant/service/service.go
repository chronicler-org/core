package attendantService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantDTO "github.com/chronicler-org/core/src/attendant/dto"
	attendantExceptionMessage "github.com/chronicler-org/core/src/attendant/messages"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	attendantRepository "github.com/chronicler-org/core/src/attendant/repository"
	teamService "github.com/chronicler-org/core/src/team/service"
)

type AttendantService struct {
	attendantRepository *attendantRepository.AttendantRepository
	teamService         *teamService.TeamService
}

func InitAttendantService(
	attendantRepository *attendantRepository.AttendantRepository,
	teamService *teamService.TeamService,
) *AttendantService {
	return &AttendantService{
		attendantRepository: attendantRepository,
		teamService:         teamService,
	}
}

func (service *AttendantService) FindByID(id string) (attendantModel.Attendant, error) {
	result, err := service.attendantRepository.FindByID(id, "Team")
	manager, _ := result.(*attendantModel.Attendant)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *manager, appException.NotFoundException(attendantExceptionMessage.ATTENDANT_NOT_FOUND)
	}

	return *manager, err
}

func (service *AttendantService) Create(dto attendantDTO.CreateAttendantDTO) (attendantModel.Attendant, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return attendantModel.Attendant{}, err
	}

	team, err := service.teamService.FindByID(dto.TeamId)
	if err != nil {
		return attendantModel.Attendant{}, err
	}

	model := attendantModel.Attendant{
		ID:        uuid.New(),
		Name:      dto.Name,
		CPF:       dto.CPF,
		Email:     dto.Email,
		Password:  string(newPassword),
		Team:      team,
		BirthDate: dto.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.attendantRepository.Create(model)

	return model, err
}

func (service *AttendantService) Update(id string, dto attendantDTO.UpdateAttendantDTO) (attendantModel.Attendant, error) {
	attendantExists, err := service.FindByID(id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}

	appUtil.UpdateModelFromDTO(&attendantExists, dto)
	if dto.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
		if err == nil {
			attendantExists.Password = string(newPassword)
		}
	}
	if dto.TeamId != "" {
		team, err := service.teamService.FindByID(dto.TeamId)
		if err != nil {
			return attendantModel.Attendant{}, err
		}
		attendantExists.Team = team
	}

	attendantExists.UpdatedAt = time.Now()
	err = service.attendantRepository.Update(attendantExists)
	return attendantExists, err
}

func (service *AttendantService) FindAll(dto appDto.PaginationDTO) (int64, []attendantModel.Attendant, error) {
	var attendants []attendantModel.Attendant
	totalCount, err := service.attendantRepository.FindAll(dto.GetLimit(), dto.GetPage(), &attendants, "Team")
	if err != nil {
		return 0, nil, err
	}
	return totalCount, attendants, nil
}

func (service *AttendantService) Delete(id string) (attendantModel.Attendant, error) {
	attendantExists, err := service.FindByID(id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}

	err = service.attendantRepository.Delete(id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}
	return attendantExists, nil
}
