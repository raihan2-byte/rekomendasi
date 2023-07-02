package hotel

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Hotel, error)
	FindById(ID int) (Hotel, error)
	FindByRating(rating float32) ([]Hotel, error)
	FindByLokasi(lokasi string) ([]Hotel, error)
	FindByBudget(budget int) ([]Hotel, error)
	FindByBudgetAndLokasi(lokasi string, budget int) ([]Hotel, error)
	FindByRatingAndBudget(rating float32, budget int) ([]Hotel, error)
	FindByAll(rating float32, lokasi string, budget int) ([]Hotel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Hotel, error) {
	var hotel []Hotel

	err := r.db.Find(&hotel).Error
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (r *repository) FindById(ID int) (Hotel, error) {
	var hotel Hotel

	err := r.db.Where("id = ?", ID).Find(&hotel).Error

	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (r *repository) FindByAll(rating float32, lokasi string, budget int) ([]Hotel, error){
	var all []Hotel

	err := r.db.Where("rating = ? and lokasi like ? and budget = ?", rating, "%"+lokasi+"%", budget).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil

}

func (r *repository) FindByRatingAndBudget(rating float32, budget int) ([]Hotel, error){
	var all []Hotel

	err := r.db.Where("rating = ? and budget = ?", rating, budget).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil
}


func (r *repository) FindByBudgetAndLokasi(lokasi string, budget int) ([]Hotel, error){
	var all []Hotel

	err := r.db.Where("lokasi like ? and budget = ?", "%"+lokasi+"%", budget).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil
}

func (r *repository) FindByBudget(budget int) ([]Hotel, error){
	var newBudget []Hotel

	err := r.db.Where("budget = ?", budget).Find(&newBudget).Error
	if err != nil {
		return newBudget, err
	}
	return newBudget, nil
}


func (r *repository) FindByRating(rating float32) ([]Hotel, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var rate []Hotel

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("rating = ?", rating).Find(&rate).Error
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (r *repository) FindByLokasi(lokasi string) ([]Hotel, error) {
	var lokation []Hotel

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("lokasi like ?", "%"+lokasi+"%").Find(&lokation).Error
	if err != nil {
		return lokation, err
	}
	return lokation, nil
}
