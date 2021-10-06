package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	model "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
)

func (repo *MDRepository) InsertProgram(input domain.Program) error {
	m := model.ProgramDomainToModel(&input)
	err := repo.db.Create(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *MDRepository) GetAllProgram() ([]*domain.Program, error) {
	var m []*model.Program

	err := repo.db.
		Find(&m).Order("id DESC").Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := []*domain.Program{}
	for _, val := range m {
		output = append(output, model.ProgramModelToDomain(val))
	}
	return output, nil
}
