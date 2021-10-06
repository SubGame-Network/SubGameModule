package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	subgame "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *MDRepository) SubGameGetLastBlockLog() (*domain.SubGameBlockLog, error) {
	var m subgame.SubgameBlockLog
	err := repo.db.Last(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	output := subgame.SubGameBlockLogModelToDomain(&m)
	return output, nil
}

func (repo *MDRepository) SubGameInsertUpdateBlockLog(input *domain.SubGameBlockLog) {
	m := subgame.SubGameBlockLogDomainToModel(input)
	repo.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Table: "subgame_block_log", Name: "num"},
		},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"done":        m.Done,
			"error_count": m.ErrorCount,
			"updated_at":  time.Now().Format(time.RFC3339),
		}),
	}).Create(&m)
}

func (repo *MDRepository) SubGameDeleteOldBlock(maxBlockNum uint64) error {
	return repo.db.
		Where("num < ?", maxBlockNum).
		Where("done = ?", 1).
		Delete(&subgame.SubgameBlockLog{}).
		Error
}

func (repo *MDRepository) SubGameGetBlockGapLog() ([]*domain.SubGameBlockLog, error) {
	var m []*subgame.SubgameBlockLog
	err := repo.db.
		Where("done = ?", 0).
		Order("num ASC").
		Find(&m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if m == nil {
		return nil, nil
	}

	output := []*domain.SubGameBlockLog{}
	for _, val := range m {
		output = append(output, subgame.SubGameBlockLogModelToDomain(val))
	}
	return output, nil
}
