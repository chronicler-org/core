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
	attendantRepository           *attendantRepository.AttendantRepository
	attendantEvaluationRepository *attendantRepository.AttendantEvaluationRepository
	teamService                   *teamService.TeamService
}

func InitAttendantService(
	attendantRepository *attendantRepository.AttendantRepository,
	attendantEvaluationRepository *attendantRepository.AttendantEvaluationRepository,
	teamService *teamService.TeamService,
) *AttendantService {
	return &AttendantService{
		attendantRepository:           attendantRepository,
		attendantEvaluationRepository: attendantEvaluationRepository,
		teamService:                   teamService,
	}
}

func (service *AttendantService) FindAttendantByID(id string) (attendantModel.Attendant, error) {
	result, err := service.attendantRepository.FindOneByField("ID", id, "Team")
	attendant, _ := result.(*attendantModel.Attendant)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *attendant, appException.NotFoundException(attendantExceptionMessage.ATTENDANT_NOT_FOUND)
	}

	return *attendant, err
}

func (service *AttendantService) FindAttendantByEmail(email string) (attendantModel.Attendant, error) {
	result, err := service.attendantRepository.FindOneByField("Email", email, "Team")
	attendant, _ := result.(*attendantModel.Attendant)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *attendant, appException.NotFoundException(attendantExceptionMessage.ATTENDANT_NOT_FOUND)
	}
	return *attendant, nil
}

func (service *AttendantService) CreateAttendant(dto attendantDTO.CreateAttendantDTO) (attendantModel.Attendant, error) {
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
		TeamID:    team.ID,
		BirthDate: dto.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.attendantRepository.Create(model)
	model.Team = team
	return model, err
}

func (service *AttendantService) UpdateAttendant(id string, dto attendantDTO.UpdateAttendantDTO) (attendantModel.Attendant, error) {
	attendantExists, err := service.FindAttendantByID(id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}

	appUtil.UpdateModelFromDTO(&attendantExists, &dto)
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
		attendantExists.TeamID = team.ID
	}

	attendantExists.UpdatedAt = time.Now()
	err = service.attendantRepository.Update(attendantExists)
	return attendantExists, err
}

func (service *AttendantService) FindAllAttendants(dto appDto.PaginationDTO) (int64, []attendantModel.Attendant, error) {
	var attendants []attendantModel.Attendant
	totalCount, err := service.attendantRepository.FindAll(dto, &attendants, "Team")
	if err != nil {
		return 0, nil, err
	}
	return totalCount, attendants, nil
}

func (service *AttendantService) DeleteAttendant(id string) (attendantModel.Attendant, error) {
	attendantExists, err := service.FindAttendantByID(id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}

	err = service.attendantRepository.Delete("ID", id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}
	return attendantExists, nil
}

func (service *AttendantService) FindAttendantEvaluationByID(id string) (attendantModel.AttendantEvaluation, error) {
	result, err := service.attendantEvaluationRepository.FindOneByField("ID", id, "Avaluated", "Avaluator")
	attendantEvaluation, _ := result.(*attendantModel.AttendantEvaluation)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *attendantEvaluation, appException.NotFoundException(attendantExceptionMessage.ATTENDANT_EVALUATION_NOT_FOUND)
	}
	return *attendantEvaluation, nil
}

func (service *AttendantService) CreateAttendantEvaluation(
	dto attendantDTO.CreateAttendantEvaluationDTO,
	avaluator attendantModel.Attendant,
) (attendantModel.AttendantEvaluation, error) {

	AvaluatedExists, err := service.FindAttendantByID(dto.AvaluatedID)
	if err != nil {
		return attendantModel.AttendantEvaluation{}, err
	}

	if AvaluatedExists.ID.String() == avaluator.ID.String() {
		return attendantModel.AttendantEvaluation{}, appException.ConflictException(attendantExceptionMessage.ATTENDANT_EVALUATION_SELF_EVALUATION)
	}

	model := attendantModel.AttendantEvaluation{
		Score:       dto.Score,
		Description: dto.Description,
		AvaluatedID: AvaluatedExists.ID,
		AvaluatorID: avaluator.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = service.attendantEvaluationRepository.Create(model)
	if err != nil {
		return attendantModel.AttendantEvaluation{}, err
	}

	model.Avaluated = AvaluatedExists
	model.Avaluator = avaluator
	return model, err
}

func (service *AttendantService) UpdateAttendantEvaluation(
	id string,
	dto attendantDTO.UpdateAttendantEvaluationDTO,
) (attendantModel.AttendantEvaluation, error) {
	attendantEvaluationExists, err := service.FindAttendantEvaluationByID(id)
	if err != nil {
		return attendantModel.AttendantEvaluation{}, err
	}

	appUtil.UpdateModelFromDTO(&attendantEvaluationExists, &dto)

	attendantEvaluationExists.UpdatedAt = time.Now()
	err = service.attendantEvaluationRepository.Update(attendantEvaluationExists)
	if err != nil {
		return attendantModel.AttendantEvaluation{}, err
	}

	return attendantEvaluationExists, err
}


func (service *AttendantService) FindAllAttedantEvaluations(
	dto attendantDTO.QueryAttendantEvaluationDTO,
) (int64, []attendantModel.AttendantEvaluation, error) {

	var attendantEvaluations []attendantModel.AttendantEvaluation
	totalCount, err := service.attendantEvaluationRepository.FindAll(dto, &attendantEvaluations, "Avaluated", "Avaluator")
	if err != nil {
		return 0, nil, err
	}
	return totalCount, attendantEvaluations, nil
}

func (service *AttendantService) DeleteAttedantEvaluation(evaluationId string) (attendantModel.AttendantEvaluation, error) {
	attendantEvaluationExists, err := service.FindAttendantEvaluationByID(evaluationId)
	if err != nil {
		return attendantModel.AttendantEvaluation{}, err
	}

	err = service.attendantEvaluationRepository.Delete("ID", evaluationId)
	if err != nil {
		return attendantModel.AttendantEvaluation{}, err
	}
	return attendantEvaluationExists, nil
}
