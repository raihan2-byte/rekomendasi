package food

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Food, error)
	FindById(ID int) (Food, error)
	FindByRating(rating float32) ([]Food, error)
	FindByLokasi(lokasi string) ([]Food, error)
	FindByAll(rating float32, lokasi string) ([]Food, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Food, error) {
	var food []Food

	err := r.db.Find(&food).Error
	if err != nil {
		return food, err
	}
	return food, nil
}

func (r *repository) FindById(ID int) (Food, error) {
	var food Food

	err := r.db.Where("id = ?", ID).Find(&food).Error

	if err != nil {
		return food, err
	}
	return food, nil
}

func (r *repository) FindByAll(rating float32, lokasi string) ([]Food, error){
	var all []Food

	err := r.db.Where("rating = ? and lokasi like ?", rating, lokasi).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil

}

func (r *repository) FindByRating(rating float32) ([]Food, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var rate []Food

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("rating = ?", rating).Find(&rate).Error
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (r *repository) FindByLokasi(lokasi string) ([]Food, error) {
	var lokation []Food

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("lokasi = ?", lokasi).Find(&lokation).Error
	if err != nil {
		return lokation, err
	}
	return lokation, nil
}
