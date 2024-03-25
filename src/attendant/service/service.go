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
)

type AttendantService struct {
	repository *attendantRepository.AttendantRepository
}

func InitAttendantService(r *attendantRepository.AttendantRepository) *AttendantService {
	return &AttendantService{
		repository: r,
	}
}

func (service *AttendantService) FindByID(id string) (attendantModel.Attendant, error) {
	result, err := service.repository.FindByID(id)
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

	model := attendantModel.Attendant{
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

	attendantExists.UpdatedAt = time.Now()
	err = service.repository.Update(attendantExists)
	return attendantExists, err
}

func (service *AttendantService) FindAll(dto appDto.PaginationDTO) (int64, []attendantModel.Attendant, error) {
	var attendants []attendantModel.Attendant
	totalCount, err := service.repository.FindAll(dto.GetLimit(), dto.GetPage(), &attendants)
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

	err = service.repository.Delete(id)
	if err != nil {
		return attendantModel.Attendant{}, err
	}
	return attendantExists, nil
}