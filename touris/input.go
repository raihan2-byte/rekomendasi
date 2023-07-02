package touris

//menginisiasi input dari user
type InputGetTouris struct {
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
	Budget int `json:"budget"`
}