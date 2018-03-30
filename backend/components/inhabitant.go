package components

type Inhabitant struct {
	height int
	weight int
	limbs  Limbs
}

type Limbs struct {
	count  int
	length int
}
