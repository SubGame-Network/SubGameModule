package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	model "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *MDRepository) InsertContact(input domain.Contact) error {
	m := model.ContactDomainToModel(&input)
	err := repo.db.Create(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *MDRepository) UpdateContact(module *domain.Contact) error {
	var m model.Contact

	m_module := model.ContactDomainToModel(module)
	err := repo.db.Model(&m).
		Where("id = ?", module.ID).
		Updates(m_module).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MDRepository) GetAllContact() ([]*domain.Contact, error) {
	var m []*model.Contact

	err := repo.db.
		Find(&m).Order("id DESC").Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := []*domain.Contact{}
	for _, val := range m {
		output = append(output, model.ContactModelToDomain(val))
	}
	return output, nil
}

func (repo *MDRepository) GetContactById(id uint64) (*domain.Contact, error) {
	var m *model.Contact

	err := repo.db.
		Where("id = ?", id).
		First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.ContactModelToDomain(m), nil
}
