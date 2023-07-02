package touris

//pembuatan entitas seusai dengan table database
type Touris struct {
	ID         int
	Nama       string
	Rating     float32
	Lokasi   string
	ImgURL     string
	Budget 	   int
}
