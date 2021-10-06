package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	model "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *MDRepository) InsertModule(input domain.Module) error {
	m := model.ModuleDomainToModel(&input)
	err := repo.db.Create(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *MDRepository) UpdateModule(module *domain.Module) error {
	var m model.Module

	m_module := model.ModuleDomainToModel(module)
	err := repo.db.Model(&m).
		Where("id = ?", module.ID).
		Updates(m_module).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MDRepository) GetAllModule() ([]*domain.Module, error) {
	var m []*model.Module

	err := repo.db.
		Find(&m).Order("id DESC").Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := []*domain.Module{}
	for _, val := range m {
		output = append(output, model.ModuleModelToDomain(val))
	}
	return output, nil
}

func (repo *MDRepository) GetModuleById(id uint64) (*domain.Module, error) {
	var m *model.Module

	err := repo.db.
		Where("id = ?", id).
		First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.ModuleModelToDomain(m), nil
}
