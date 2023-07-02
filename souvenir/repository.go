package souvenir

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Souvenir, error)
	FindById(ID int) (Souvenir, error)
	FindByRating(rating float32) ([]Souvenir, error)
	FindByLokasi(lokasi string) ([]Souvenir, error)
	FindByAll(rating float32, lokasi string) ([]Souvenir, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Souvenir, error) {
	var souvenir []Souvenir

	err := r.db.Find(&souvenir).Error
	if err != nil {
		return souvenir, err
	}
	return souvenir, nil
}

func (r *repository) FindById(ID int) (Souvenir, error) {
	var souvenir Souvenir

	err := r.db.Where("id = ?", ID).Find(&souvenir).Error

	if err != nil {
		return souvenir, err
	}
	return souvenir, nil
}

func (r *repository) FindByAll(rating float32, lokasi string) ([]Souvenir, error){
	var all []Souvenir

	err := r.db.Where("rating = ? and lokasi = ?", rating, lokasi).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil

}

func (r *repository) FindByRating(rating float32) ([]Souvenir, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var rate []Souvenir

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("rating = ?", rating).Find(&rate).Error
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (r *repository) FindByLokasi(lokasi string) ([]Souvenir, error) {
	var lokation []Souvenir

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("location like ?", "%"+lokasi+"%").Find(&lokation).Error
	if err != nil {
		return lokation, err
	}
	return lokation, nil
}
