package souvenir

type Service interface {
	GetSouvenir(ID int) ([]Souvenir, error)
	GetByInput(input InputSouvenir) ([]Souvenir, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetSouvenir(ID int) ([]Souvenir, error) {
	photo, err := s.repository.FindAll()
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *service) GetByInput(input InputSouvenir) ([]Souvenir, error) {
	if input.Rating != 0 && input.Lokasi != "" {
		newSouvenir, err := s.repository.FindByAll(input.Rating, input.Lokasi)
		if err != nil {
			return newSouvenir, err
		}
		return newSouvenir, nil
	} else if input.Lokasi != "" {
		newLokasi, err := s.repository.FindByLokasi(input.Lokasi)
		if err != nil {
			return newLokasi, err
		}
		return newLokasi, nil
	} else {
		newRating, err := s.repository.FindByRating(input.Rating)
		if err != nil {
			return newRating, err
		}
		return newRating, nil
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
