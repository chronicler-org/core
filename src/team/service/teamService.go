package teamService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	teamDTO "github.com/chronicler-org/core/src/team/dto"
	teamExceptionMessage "github.com/chronicler-org/core/src/team/messages"
	teamModel "github.com/chronicler-org/core/src/team/model"
	teamRepository "github.com/chronicler-org/core/src/team/repository"
)

type TeamService struct {
	teamRepository *teamRepository.TeamRepository
}

func InitTeamService(r *teamRepository.TeamRepository) *TeamService {
	return &TeamService{
		teamRepository: r,
	}
}

func (service *TeamService) FindByID(id string) (teamModel.Team, error) {
	result, err := service.teamRepository.FindByID(id)
	team, _ := result.(*teamModel.Team)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *team, appException.NotFoundException(teamExceptionMessage.TEAM_NOT_FOUND)
	}
	return *team, nil
}

func (service *TeamService) Create(dto teamDTO.CreateTeamDTO) (teamModel.Team, error) {
	model := teamModel.Team{
		ID:        uuid.New(),
		Name:      dto.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := service.teamRepository.Create(model)

	return model, err
}

func (service *TeamService) Update(id string, dto teamDTO.UpdateTeamDTO) (teamModel.Team, error) {
	teamExists, err := service.FindByID(id)
	if err != nil {
		return teamModel.Team{}, err
	}

	appUtil.UpdateModelFromDTO(&teamExists, dto)

	teamExists.UpdatedAt = time.Now()
	err = service.teamRepository.Update(teamExists)
	return teamExists, err
}

func (service *TeamService) FindAll(dto appDto.PaginationDTO) (int64, []teamModel.Team, error) {
	var teams []teamModel.Team
	totalCount, err := service.teamRepository.FindAll(dto, &teams)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, teams, nil
}

func (service *TeamService) Delete(id string) (teamModel.Team, error) {
	teamExists, err := service.FindByID(id)
	if err != nil {
		return teamModel.Team{}, err
	}

	err = service.teamRepository.Delete(id)
	if err != nil {
		return teamModel.Team{}, err
	}
	return teamExists, nil
}
