package components

type Ground struct {
	width     int
	height    int
	placeSize int
}

func NewGround(w, h, placeSize int) *Ground {
	ground := Ground{w, h, placeSize}

	return &ground
}
