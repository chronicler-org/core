package tagService

import (
	"time"

	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagModel "github.com/chronicler-org/core/src/tag/model"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
	"github.com/google/uuid"
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
		return uuid.Nil, nil
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

func (service *TagService) FindAll() ([]tagModel.Tag, error) {
	return service.repository.FindAll()
}

func (service *TagService) Delete(id string) error {
	return service.repository.Delete(id)
}
