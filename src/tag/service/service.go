package tagService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagExceptionMessage "github.com/chronicler-org/core/src/tag/messages"
	tagModel "github.com/chronicler-org/core/src/tag/model"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
)

type TagService struct {
	tagRepository *tagRepository.TagRepository
}

func InitTagService(r *tagRepository.TagRepository) *TagService {
	return &TagService{
		tagRepository: r,
	}
}

func (service *TagService) FindByID(id string) (tagModel.Tag, error) {
	result, err := service.tagRepository.FindByID(id)
	tag, _ := result.(*tagModel.Tag)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *tag, appException.NotFoundException(tagExceptionMessage.TAG_NOT_FOUND)
	}
	return *tag, nil
}

func (service *TagService) Create(dto tagDTO.CreateTagDTO) (tagModel.Tag, error) {
	model := tagModel.Tag{
		ID:        uuid.New(),
		Name:      dto.Name,
		Color:     dto.Color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := service.tagRepository.Create(model)

	return model, err
}

func (service *TagService) Update(id string, dto tagDTO.UpdateTagDTO) (tagModel.Tag, error) {
	tagExists, err := service.FindByID(id)
	if err != nil {
		return tagModel.Tag{}, err
	}

	appUtil.UpdateModelFromDTO(&tagExists, dto)

	tagExists.UpdatedAt = time.Now()
	err = service.tagRepository.Update(tagExists)
	return tagExists, err
}

func (service *TagService) FindAll(dto appDto.PaginationDTO) (int64, []tagModel.Tag, error) {
	var tags []tagModel.Tag
	totalCount, err := service.tagRepository.FindAll(dto.GetLimit(), dto.GetPage(), &tags)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, tags, nil
}

func (service *TagService) Delete(id string) (tagModel.Tag, error) {
	tagExists, err := service.FindByID(id)
	if err != nil {
		return tagModel.Tag{}, err
	}

	err = service.tagRepository.Delete(id)
	if err != nil {
		return tagModel.Tag{}, err
	}
	return tagExists, nil
}
