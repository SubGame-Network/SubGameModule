package gorm

import (
	"github.com/SubGame-Network/SubGameModuleService/domain"
	model "github.com/SubGame-Network/SubGameModuleService/internal/md/model/gorm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *MDRepository) InsertUser(input domain.User) error {
	m := model.UserDomainToModel(&input)
	err := repo.db.Create(&m).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo *MDRepository) UpdateUser(user *domain.User) error {
	var m model.User

	m_user := model.UserDomainToModel(user)
	err := repo.db.Model(&m).
		Where("id = ?", user.ID).
		Updates(m_user).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *MDRepository) GetAllUser() ([]*domain.User, error) {
	var m []*model.User

	err := repo.db.
		Find(&m).Order("id DESC").Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	output := []*domain.User{}
	for _, val := range m {
		output = append(output, model.UserModelToDomain(val))
	}
	return output, nil
}
func (repo *MDRepository) GetAllUserByPage(row, page int, account string, email string, address string) ([]*domain.User, int64, error) {
	// page = 0，代表第一頁
	offset := page * row

	var m []*model.User

	sql := repo.db.Table("user")

	if account != "" {
		sql = sql.Where("account LIKE ?", account+"%")
	}

	if email != "" {
		sql = sql.Where("email LIKE ?", email+"%")
	}

	if address != "" {
		sql = sql.Where("address LIKE ?", address+"%")
	}

	var count int64
	sql.Count(&count)

	if row > 0 {
		sql = sql.Offset(offset).Limit(row)
	}

	err := sql.Order("id DESC").Find(&m).Error
	if err != nil {
		return nil, count, errors.WithStack(err)
	}

	var output []*domain.User
	for _, v := range m {
		output = append(output, model.UserModelToDomain(v))
	}
	return output, count, nil
}

func (repo *MDRepository) GetUserByAccountOrEmailOrAddress(account string, email string, address string) (*domain.User, error) {
	var m *model.User

	err := repo.db.
		Where("email = ?", email).
		Or("account = ?", account).
		Or("address = ?", address).
		First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.UserModelToDomain(m), nil
}

func (repo *MDRepository) GetUserByAccount(account string) (*domain.User, error) {
	var m *model.User

	err := repo.db.
		Where("account = ?", account).
		First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.UserModelToDomain(m), nil
}

func (repo *MDRepository) GetUserById(id uint64) (*domain.User, error) {
	var m *model.User

	err := repo.db.
		Where("id = ?", id).
		First(&m).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.UserModelToDomain(m), nil
}

func (repo *MDRepository) GetUserByAddress(address string) (*domain.User, error) {
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

func (repo *MDRepository) GetUserByEmail(email string) (*domain.User, error) {
	var m model.User
	err := repo.db.Where("email = ?", email).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	output := model.UserModelToDomain(&m)
	return output, nil
}

func (repo *MDRepository) UpdateUserNonce(id uint64, nonce string) error {
	var m model.User
	err := repo.db.Model(&m).
		Where("id = ?", id).
		Update("nonce", nonce).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
