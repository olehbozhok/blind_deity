package components

type Inhabitant struct {
	reproduction Reproduction
	moves        int
	height       int
	weight       int
	limbs        Limbs
}

type Reproduction struct {
	change    int
	frequency int
}

type Limbs struct {
	count  int
	length int
}
