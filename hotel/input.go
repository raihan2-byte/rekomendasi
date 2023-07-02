package hotel

// type GetHotelID struct {
// 	ID int `uri:"id" binding:"required"`
// }

type InputHotel struct {
	Lokasi string `json:"lokasi"`
	Rating float32 `json:"rating"`
	Budget int `json:"budget"`
}