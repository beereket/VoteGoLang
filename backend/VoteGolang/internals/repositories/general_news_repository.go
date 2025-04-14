package repositories

import (
	"VoteGolang/internals/data"
	"gorm.io/gorm"
)

type GeneralNewsRepository interface {
	Create(news *data.GeneralNews) error
	GetAll() ([]data.GeneralNews, error)
	GetByID(id uint) (*data.GeneralNews, error)
	Update(news *data.GeneralNews) error
	Delete(id uint) error
}

type generalNewsRepository struct {
	db *gorm.DB
}

func NewGeneralNewsRepository(db *gorm.DB) GeneralNewsRepository {
	return &generalNewsRepository{db: db}
}

func (r *generalNewsRepository) Create(news *data.GeneralNews) error {
	return r.db.Create(news).Error
}

func (r *generalNewsRepository) GetAll() ([]data.GeneralNews, error) {
	var newsList []data.GeneralNews
	err := r.db.Find(&newsList).Error
	return newsList, err
}

func (r *generalNewsRepository) GetByID(id uint) (*data.GeneralNews, error) {
	var news data.GeneralNews
	err := r.db.First(&news, id).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *generalNewsRepository) Update(news *data.GeneralNews) error {
	return r.db.Save(news).Error
}

func (r *generalNewsRepository) Delete(id uint) error {
	return r.db.Delete(&data.GeneralNews{}, id).Error
}
