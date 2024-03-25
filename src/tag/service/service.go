package tagService

import (
	"time"

	"github.com/google/uuid"

	appDto "github.com/chronicler-org/core/src/app/dto"
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagModel "github.com/chronicler-org/core/src/tag/model"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
	serviceErrors "github.com/chronicler-org/core/src/utils/errors"
)

type TagService struct {
	repository *tagRepository.TagRepository
}

func InitTagService(r *tagRepository.TagRepository) *TagService {
	return &TagService{
		repository: r,
	}
}

func (service *TagService) FindByID(id string) (tagModel.Tag, error) {
	return service.repository.FindByID(id)
}

func (service *TagService) Create(dto tagDTO.CreateTagDTO) (uuid.UUID, error) {
	if !dto.ValidateHexColor() {
		return uuid.Nil, serviceErrors.NewError("erro ao validar o codigo hex para a cor da tag")
	}
	model := tagModel.Tag{
		ID:        uuid.New(),
		Title:     dto.Title,
		Color:     dto.Color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := service.repository.Create(model)

	return model.ID, err
}

func (service *TagService) Update(id string, dto tagDTO.UpdateTagDTO) (tagModel.Tag, error) {
	updatedTag, err := service.repository.FindByID(id)
	// implementar validaçao da tag
	if err != nil {
		return updatedTag, err
	}
	if updatedTag.ID == uuid.Nil {
		return updatedTag, err
	}

	if dto.Title != "" {
		updatedTag.Title = dto.Title
	}
	if dto.Color != "" {
		updatedTag.Color = dto.Color
	}

	updatedTag.UpdatedAt = time.Now()

	err = service.repository.Update(updatedTag)

	return updatedTag, err
}

func (service *TagService) FindAll(dto appDto.PaginationDTO) (int64, []tagModel.Tag, error) {
	var tags []tagModel.Tag
	totalCount, err := service.repository.FindAll(dto.GetLimit(), dto.GetPage(), &tags)
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

	err = service.repository.Delete(id)
	if err != nil {
		return tagModel.Tag{}, err
	}
	return tagExists, nil
}
