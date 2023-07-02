package food

// type GetHotelID struct {
// 	ID int `uri:"id" binding:"required"`
// }

type InputFood struct {
	Lokasi string `json:"lokasi"`
	Rating float32 `json:"rating"`
}