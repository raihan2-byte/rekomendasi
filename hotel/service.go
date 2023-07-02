package hotel

type Service interface {
	GetHotels(ID int) ([]Hotel, error)
	GetByInput(input InputHotel) ([]Hotel, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetHotels(ID int) ([]Hotel, error) {
	photo, err := s.repository.FindAll()
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *service) GetByInput(input InputHotel) ([]Hotel, error) {

	if input.Rating != 0 && input.Lokasi != "" && input.Budget != 0 {
		newhotel, err := s.repository.FindByAll(input.Rating, input.Lokasi, input.Budget)
		if err != nil {
			return newhotel, err
		}
		return newhotel, nil
	} else if input.Lokasi != "" && input.Budget != 0 {
		newLokasi, err := s.repository.FindByBudgetAndLokasi(input.Lokasi, input.Budget)
		if err != nil {
			return newLokasi, err
		}
		return newLokasi, nil
	} else if input.Rating != 0 && input.Budget != 0 {
		newRating, err := s.repository.FindByRatingAndBudget(input.Rating, input.Budget)
		if err != nil {
			return newRating, err
		}
		return newRating, nil
	} else if input.Lokasi != "" {
		newBudget, err := s.repository.FindByLokasi(input.Lokasi)
		if err != nil {
			return newBudget, err
		}
		return newBudget, nil
	} else if input.Rating != 0 {
		newBudget, err := s.repository.FindByRating(input.Rating)
		if err != nil {
			return newBudget, err
		}
		return newBudget, nil
	} else {
		newBudget, err := s.repository.FindByBudget(input.Budget)
		if err != nil {
			return newBudget, err
		}
		return newBudget, nil
	}
	// return newHotel, nil
	// if err != nil {
	// 	return newHotel, err
}

// if err != nil {
// 	return newHotel, err
// }

// newLokasi, err := s.repository.FindByLokasi(input.Lokasi)
// if err != nil {
// 	return newLokasi, err
// }
// return newLokasi, nil

// return newLokasi, nil
// newRate, err := s.repository.FindAll()
// if err != nil {
// 	return newRate, err
// }
// return newHotel, nil
