package gorm

import (
	"time"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	model "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func (repo *MDRepository) StakeGetUserByAddress(address string) (*domain.User, error) {
	var m model.User
	err := repo.db.Where("address = ?", address).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	output := model.UserModelToDomain(&m)
	return output, nil
}

func (repo *MDRepository) StakeInsertRecord(input domain.StakeRecord) error {

	// 新增提領記錄
	m := model.StakeRecordDomainToModel(&input)
	err := repo.db.Create(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MDRepository) StakeRecordByUserIdAndTime(userId uint64, from, to time.Time) ([]*domain.StakeRecord, error) {
	var m []*model.StakeRecord

	sql := repo.db.Where("user_id = ?", userId)
	if !from.IsZero() {
		sql = sql.Where("created_at >= ?", from.Format(time.RFC3339))
	}
	if !to.IsZero() {
		sql = sql.Where("created_at <= ?", to.Format(time.RFC3339))
	}

	err := sql.Order("id DESC").Find(&m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var output []*domain.StakeRecord
	for _, v := range m {
		output = append(output, model.StakeRecordModelToDomain(v))
	}
	return output, nil
}

func (repo *MDRepository) StakeAllRecords(from, to time.Time) ([]*domain.StakeRecord, int64, error) {
	var m []*model.StakeRecord

	sql := repo.db.Table("withdraw_record")
	if !from.IsZero() {
		sql = sql.Where("created_at >= ?", from.Format(time.RFC3339))
	}
	if !to.IsZero() {
		sql = sql.Where("created_at <= ?", to.Format(time.RFC3339))
	}

	var count int64
	sql.Count(&count)

	err := sql.Order("id DESC").Find(&m).Error
	if err != nil {
		return nil, count, errors.WithStack(err)
	}

	var output []*domain.StakeRecord
	for _, v := range m {
		output = append(output, model.StakeRecordModelToDomain(v))
	}
	return output, count, nil
}

func (repo *MDRepository) StakeRecordByID(id uint64) (*domain.StakeRecord, error) {
	var m *model.StakeRecord
	err := repo.db.Where("id = ?", id).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	output := model.StakeRecordModelToDomain(m)
	return output, nil
}

func (repo *MDRepository) StakeUpdateRecordByID(id uint64, input domain.StakeRecord) error {
	m := model.StakeRecordDomainToModel(&input)
	err := repo.db.Where("id = ?", id).Updates(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *MDRepository) StakeGetTimeout() ([]*domain.StakeRecord, error) {
	var m []*model.StakeRecord
	err := repo.db.
		Where("tx_status = ?", domain.Pending).
		Where("notify_time IS NULL").
		Where("DATE_ADD(created_at, INTERVAL 10 MINUTE) <= ?", time.Now().Format(time.RFC3339)).
		Find(&m).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := []*domain.StakeRecord{}
	for _, val := range m {
		output = append(output, model.StakeRecordModelToDomain(val))
	}
	return output, nil
}

func (repo *MDRepository) StakeUpdateNotifyTime(id uint64) error {
	var m model.StakeRecord
	err := repo.db.Model(&m).
		Where("id = ?", id).
		Update("notify_time", time.Now().Format(time.RFC3339)).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *MDRepository) StakeAllRecordsByPage(userId uint64, row, page int, NftHash string, status *int, periodOfUseMonth int) ([]*domain.StakeRecord, int64, error) {
	// page = 0，代表第一頁
	offset := page * row
	var m []*model.StakeRecord

	sql := repo.db.Table("stake_record").Where("user_id = ?", userId)
	if NftHash != "" {
		sql = sql.Where("nft_hash Like ?", NftHash+"%")
	}

	if status != nil {
		sql = sql.Where("tx_status = ?", *status)
	}

	if periodOfUseMonth != 0 {
		sql = sql.Where("period_of_use_month = ?", periodOfUseMonth)
	}

	var count int64
	sql.Count(&count)

	if row > 0 {
		sql = sql.Offset(offset).Limit(row)
	}

	err := sql.Order("id DESC").Find(&m).Error
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	var output []*domain.StakeRecord
	for _, v := range m {
		output = append(output, model.StakeRecordModelToDomain(v))
	}
	return output, count, nil
}

func (repo *MDRepository) StakeTotalStakedByUserId(userId uint64) (decimal.Decimal, error) {
	var m model.StakeRecord

	amountStruct := struct {
		Amount decimal.Decimal
	}{}

	sql := repo.db.Model(&m).
		Select("SUM(stake_sgb) AS amount").
		Where("user_id = ?", userId).
		Where("end_time > now()")

	err := sql.First(&amountStruct).Error
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}
	return amountStruct.Amount, nil
}

func (repo *MDRepository) StakeTotalWithrawnByUserId(userId uint64) (decimal.Decimal, error) {
	var m model.StakeRecord

	amountStruct := struct {
		Amount decimal.Decimal
	}{}

	sql := repo.db.Model(&m).
		Select("SUM(stake_sgb) AS amount").
		Where("user_id = ?", userId).
		Where("end_time <= now()")

	err := sql.First(&amountStruct).Error
	if err != nil {
		return decimal.Zero, errors.WithStack(err)
	}
	return amountStruct.Amount, nil
}
