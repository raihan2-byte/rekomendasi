package cafe

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Cafe, error)
	FindById(ID int) (Cafe, error)
	FindByRating(rating float32) ([]Cafe, error)
	FindByLokasi(lokasi string) ([]Cafe, error)
	FindByAll(rating float32, lokasi string) ([]Cafe, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Cafe, error) {
	var cafe []Cafe

	err := r.db.Find(&cafe).Error
	if err != nil {
		return cafe, err
	}
	return cafe, nil
}

func (r *repository) FindById(ID int) (Cafe, error) {
	var cafe Cafe

	err := r.db.Where("id = ?", ID).Find(&cafe).Error

	if err != nil {
		return cafe, err
	}
	return cafe, nil
}

func (r *repository) FindByAll(rating float32, lokasi string) ([]Cafe, error){
	var all []Cafe

	err := r.db.Where("rating = ? and lokasi like ?", rating, lokasi).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil

}

func (r *repository) FindByRating(rating float32) ([]Cafe, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var rate []Cafe

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("rating = ?", rating).Find(&rate).Error
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (r *repository) FindByLokasi(lokasi string) ([]Cafe, error) {
	var lokation []Cafe

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("lokasi = ?", lokasi).Find(&lokation).Error
	if err != nil {
		return lokation, err
	}
	return lokation, nil
}
