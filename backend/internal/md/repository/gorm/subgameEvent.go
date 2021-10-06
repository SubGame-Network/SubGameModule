package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	model "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *MDRepository) EventGetStakeRecordByTxhash(txHash string) (*domain.StakeRecord, error) {
	var m model.StakeRecord
	err := repo.db.Where("tx_hash = ?", txHash).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := model.StakeRecordModelToDomain(&m)
	return output, nil
}

func (repo *MDRepository) EventUpdateStakeRecord(txHash string, input *domain.StakeRecord) error {
	m := model.StakeRecordDomainToModel(input)
	err := repo.db.Where("tx_hash = ?", txHash).Updates(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
