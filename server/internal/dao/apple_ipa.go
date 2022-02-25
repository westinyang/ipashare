package dao

import (
	"supersign/internal/model"

	"gorm.io/gorm"
)

func newAppleIPA(db *gorm.DB) *appleIPA {
	return &appleIPA{db}
}

type appleIPA struct {
	db *gorm.DB
}

var _ model.AppleIPAStore = (*appleIPA)(nil)

func (a *appleIPA) Create(appleIPA *model.AppleIPA) error {
	return a.db.Create(appleIPA).Error
}

func (a *appleIPA) Del(uuid string) error {
	return a.db.Where("uuid = ?", uuid).Delete(&model.AppleIPA{}).Error
}

func (a *appleIPA) Query(uuid string) (*model.AppleIPA, error) {
	var appleIPA model.AppleIPA
	err := a.db.Where("uuid = ?", uuid).Take(&appleIPA).Error
	if err != nil {
		return nil, err
	}
	return &appleIPA, nil
}

func (a *appleIPA) UpdateMobileConfigLink(uuid, mobileConfigLink string) error {
	return a.db.Model(&model.AppleIPA{}).
		Where("uuid = ?", uuid).
		Update("mobile_config_link", mobileConfigLink).Error
}

func (a *appleIPA) AddCount(uuid string, num int) error {
	return a.db.Model(&model.AppleIPA{}).
		Where("uuid = ?", uuid).
		UpdateColumn("count", gorm.Expr("count + ?", num)).Error
}

func (a *appleIPA) List(page, pageSize *int) ([]model.AppleIPA, error) {
	var appleIPAs []model.AppleIPA
	err := a.db.Scopes(paginate(page, pageSize)).Find(&appleIPAs).Error
	if err != nil {
		return nil, err
	}
	return appleIPAs, nil
}
