package touris

import "gorm.io/gorm"

type Repository interface {
	//mencari semua data, yang mana acuannya dari file touris.go 
	FindAll() ([]Touris, error)
	//mencari data dengan id yang ke brapa, karena maunya id ke brapa yauda parameternya ID 
	FindById(ID int) (Touris, error)
	//mencari rating ke brapa sesuai input, dg parmetnya rating yang berupa float yaitu ada koma
	FindByRating(rating float32) ([]Touris, error)
	//mencari lokasi nya sesuai input, dg parmetnya string karena bukan angka
	FindByLokasi(lokasi string) ([]Touris, error)
	//mencari lokasi dan rating jika user input keduanya.
	FindByAll(rating float32, lokasi string, budget int) ([]Touris, error)
	FindByBudgetAndLokasi(lokasi string, budget int) ([]Touris, error)
	FindByRatingAndBudget(rating float32, budget int) ([]Touris, error)
	FindByBudget(budget int) ([]Touris, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Touris, error) {
	var touris []Touris

	err := r.db.Find(&touris).Error
	if err != nil {
		return touris, err
	}
	return touris, nil
}

func (r *repository) FindById(ID int) (Touris, error) {
	var touris Touris

	err := r.db.Where("id = ?", ID).Find(&touris).Error

	if err != nil {
		return touris, err
	}
	return touris, nil
}

func (r *repository) FindByAll(rating float32, lokasi string, budget int) ([]Touris, error){
	var all []Touris

	err := r.db.Where("rating = ? and lokasi like ? and budget = ?", rating, "%"+lokasi+"%", budget).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil

}

func (r *repository) FindByRatingAndBudget(rating float32, budget int) ([]Touris, error){
	var all []Touris

	err := r.db.Where("rating = ? and budget = ?", rating, budget).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil
}


func (r *repository) FindByBudgetAndLokasi(lokasi string, budget int) ([]Touris, error){
	var all []Touris

	err := r.db.Where("lokasi like ? and budget = ?", "%"+lokasi+"%", budget).Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil
}

func (r *repository) FindByRating(rating float32) ([]Touris, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var rate []Touris

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("rating = ?", rating).Find(&rate).Error
	if err != nil {
		return rate, err
	}
	return rate, nil
}


func (r *repository) FindByLokasi(rating string) ([]Touris, error) {
	var lokasi []Touris

	err := r.db.Where("lokasi like ?", "%"+rating+"%").Find(&lokasi).Error
	if err != nil {
		return lokasi, err
	}
	return lokasi, nil
}

func(r *repository) FindByBudget(budget int) ([]Touris, error){
	var newBudget []Touris
	
	err := r.db.Where("budget = ?", budget).Find(&newBudget).Error
	if err != nil {
		return newBudget, err
	}
	return newBudget, nil
}

